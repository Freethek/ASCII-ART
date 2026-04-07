package banner

import "testing"

func TestLoad(t *testing.T) {
	bannerMap := Load("../banners/standard.txt")

	if len(bannerMap) !=95{
		t.Errorf("Inomplete Graphic Respresentation Of All Ascii")
	}

	if len(bannerMap['A']) != 8 {
		t.Errorf("incomplete Graphic Representation")
	}
}