# ASCII-Art

A command-line tool written in Go that takes a string as input and outputs a graphic representation of that string using ASCII characters. It supports multiple banner styles, special characters, numbers, spaces, and newline sequences.

---

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
- [Banner Files](#banner-files)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Edge Cases](#edge-cases)
- [Running Tests](#running-tests)
- [Go Concepts Used](#go-concepts-used)
- [Allowed Packages](#allowed-packages)

---

## Overview

ASCII-Art converts any printable ASCII string (characters 32–126) into a large graphical representation built entirely from ASCII characters. Each character in your input is looked up in a banner file and rendered as an 8-line tall graphic, with all characters on the same line printed side by side.

```
$ go run . "Hi"
 _    _   _
| |  | | (_)
| |__| |  _
|  __  | | |
| |  | | | |
|_|  |_| |_|
            
            
```

---

## Project Structure

```
ascii-art/
├── go.mod                  # Go module definition
├── main.go                 # Entry point — validates args, loads banner, renders output
├── banner/
│   ├── loader.go           # Reads banner file and builds map[rune][]string lookup table
│   └── loader_test.go      # Unit tests for banner loading
├── render/
│   ├── render.go           # Splits input, renders each segment row by row, returns string
│   └── render_test.go      # Unit tests for rendering logic
└── banners/
    ├── standard.txt        # Standard banner style
    ├── shadow.txt          # Shadow banner style
    └── thinkertoy.txt      # Thinkertoy banner style
```

### Why this structure?

Each package has exactly one responsibility:

- **`main.go`** — thin manager. Validates arguments, loads the banner, prints the result. Does no heavy lifting itself.
- **`banner/`** — knows everything about reading and parsing banner files. Nothing else.
- **`render/`** — knows everything about turning an input string into ASCII art output. Nothing else.

This separation makes the code easy to test, easy to read, and easy to extend.

---

## How It Works

The program runs through a clear pipeline from the moment you execute it to the moment output appears:

### Stage 1 — Argument Validation (`main.go`)

```
go run . "Hello" "shadow"
           ↓        ↓
         input   banner name (optional, defaults to "standard")
```

- If fewer than 2 or more than 3 arguments are provided → print usage to stderr and exit with code 1
- If the input string is empty `""` → exit cleanly with code 0, print nothing
- If no banner name is provided → default to `"standard"`

### Stage 2 — Banner Loading (`banner/loader.go`)

The banner file contains 95 printable ASCII characters (codes 32–126), each represented as 8 lines of ASCII art. Characters are separated by a single empty line, making each character block exactly 9 lines (8 art lines + 1 separator).

```
Line 0:    (empty separator)
Lines 1–8:  SPACE character (ASCII 32)
Line 9:    (empty separator)
Lines 10–17: ! character (ASCII 33)
...and so on
```

For any character with ASCII code `c`, its starting line in the file is:
```
startLine = (c - 32) * 9 + 1
```

The loader reads the entire file into a `[]string` slice using `bufio.Scanner`, then builds a `map[rune][]string` where:
- **Key** = the character as a `rune` (e.g. `'A'`)
- **Value** = a slice of 8 strings (the 8 art lines for that character)

### Stage 3 — Input Splitting (`render/render.go`)

The input may contain the literal two-character sequence `\n` (backslash + n), which signals a new line in the output. The input is split on this sequence into segments:

```
"Hello\nThere"   →  ["Hello", "There"]
"Hello\n\nThere" →  ["Hello", "", "There"]
"\n"             →  ["", ""]  →  trailing "" removed  →  [""]
```

A trailing empty segment is always removed to match expected output behavior.

### Stage 4 — Rendering (`render/render.go`)

For each segment:
- If the segment is `""` → add one blank line to the result
- Otherwise → for each of the 8 rows, concatenate that row from every character's art lines, then add a newline

```
"Hi" at row 0:  H[0] + i[0]  →  " _    _   _ "
"Hi" at row 1:  H[1] + i[1]  →  "| |  | | (_)"
...
"Hi" at row 7:  H[7] + i[7]  →  "            "
```

### Stage 5 — Output (`main.go`)

`render.Render()` returns the full result as a string. `main.go` prints it with `fmt.Print()` (not `fmt.Println()` — the render function already adds newlines).

---

## Banner Files

Three banner styles are included:

| Banner | Description |
|---|---|
| `standard` | Clean, bold block letters (default) |
| `shadow` | Characters with a shadow effect |
| `thinkertoy` | Playful style using symbols like `o`, `-`, `\|` |

Each banner file:
- Contains all 95 printable ASCII characters (codes 32–126)
- Each character is exactly 8 lines tall
- Characters are separated by one empty line
- Files should not be modified

---

## Installation

**Prerequisites:** Go 1.18 or higher

```bash
# Clone the repository
git clone https://zone01.git/ascii-art
cd ascii-art

# Initialize the module (if not already done)
go mod init ascii-art
```

Make sure the `banners/` folder contains `standard.txt`, `shadow.txt` and `thinkertoy.txt`.

---

## Usage

```bash
go run . "<input string>"
go run . "<input string>" "<banner>"
```

| Argument | Required | Description |
|---|---|---|
| `input string` | Yes | The text to render. Use `\n` for new lines. |
| `banner` | No | Banner style: `standard`, `shadow`, or `thinkertoy`. Defaults to `standard`. |

---

## Examples

**Empty input:**
```bash
go run . ""
# no output
```

**Single newline:**
```bash
go run . "\n"
# one blank line
```

**Simple word:**
```bash
go run . "Hello"
 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/
                               
                               
```

**Multiple lines:**
```bash
go run . "Hello\nThere"
# renders "Hello" then "There" on separate 8-line blocks
```

**With empty line between:**
```bash
go run . "Hello\n\nThere"
# renders "Hello", then a blank line, then "There"
```

**With spaces and special characters:**
```bash
go run . "Hello There"
go run . "{Hello There}"
go run . "1Hello 2There"
```

**Using a different banner:**
```bash
go run . "Hello" "shadow"
go run . "Hello" "thinkertoy"
```

---

## Edge Cases

| Input | Behaviour |
|---|---|
| `""` | Prints nothing, exits with code 0 |
| `"\n"` | Prints one blank line |
| `"Hello\n"` | Renders Hello, no extra blank line at end |
| `"Hello\n\nThere"` | Renders Hello, one blank line, then There |
| Character outside ASCII 32–126 | Prints error to stderr, exits with code 1 |
| Missing or wrong banner name | Prints error to stderr, exits with code 1 |
| Wrong number of arguments | Prints usage message to stderr, exits with code 1 |

---

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests for a specific package
go test ascii-art/banner
go test ascii-art/render
```

Expected output:
```
?       ascii-art       [no test files]
ok      ascii-art/banner
ok      ascii-art/render
```

### What is tested

**`banner/loader_test.go`**
- The map contains exactly 95 entries (all printable ASCII characters)
- Each character has exactly 8 lines of art

**`render/render_test.go`**
- Empty input returns empty string
- `\n` input returns a single newline
- A normal character returns the correct 8 lines

---

## Go Concepts Used

| Concept | Where Used |
|---|---|
| `os.Args` | Reading CLI arguments in `main.go` |
| `os.Open` + `defer` | Opening and auto-closing the banner file |
| `bufio.Scanner` | Reading the banner file line by line |
| `map[rune][]string` | Lookup table mapping characters to their art lines |
| `range` on string | Iterating over input as runes, not bytes |
| `strings.Split` | Splitting input on literal `\n` sequence |
| `strings.TrimRight` | Removing `\r` for Windows compatibility |
| Error handling | `if err != nil` pattern throughout |
| Packages and exports | `banner` and `render` packages with exported functions |
| Table-driven tests | Idiomatic Go test pattern in both test files |

---

## Allowed Packages

Only standard Go packages are used:

| Package | Purpose |
|---|---|
| `os` | `os.Args`, `os.Open`, `os.Exit`, `os.Stderr` |
| `fmt` | `fmt.Print`, `fmt.Fprintf` for output and errors |
| `strings` | `strings.Split`, `strings.TrimRight` |
| `bufio` | `bufio.NewScanner` for line-by-line file reading |
| `testing` | Unit test files only |

---

## Author

Built as part of the Zone01 Go curriculum.