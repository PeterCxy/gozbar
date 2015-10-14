package zbar

import (
	"testing"
	"fmt"
	"os"
	"image/png"
)

func TestBarcode(t *testing.T) {
	f, err := os.Open("testdata/barcode.png")

	if err != nil {
		t.Fail()
	}

	i, _ := png.Decode(f)

	img := FromImage(i)

	s := NewScanner()
	s.SetConfig(0, CFG_ENABLE, 1)

	s.Scan(img)

	img.First().Each(func(str string) {
		fmt.Println(str)

		if str != "9876543210128" {
			t.Fail()
		}
	})
}
