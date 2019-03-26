package zbar

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"testing"

	//iconv "github.com/djimenez/iconv-go"
)

func TestBarcode(t *testing.T) {
	f, err := os.Open("testdata/barcode.png")

	if err != nil {
		t.Fail()
		return
	}

	i, _ := png.Decode(f)

	img := FromImage(i)

	s := NewScanner()
	s.SetConfig(0, CFG_ENABLE, 1)

	s.Scan(img)
	defer s.Destroy()

	img.First().Each(func(str string) {
		fmt.Println(str)

		if str != "9876543210128" {
			t.Fail()
		}
	})
}

func TestQRCode(t *testing.T) {
	f, err := os.Open("testdata/qr.jpg")

	if err != nil {
		t.Fail()
		return
	}

	i, _ := jpeg.Decode(f)
	img := FromImage(i)

	s := NewScanner()
	s.SetConfig(0, CFG_ENABLE, 1)
	s.Scan(img)
	defer s.Destroy()

	img.First().Each(func(str string) {
		// Charset decoding
		//str, _ = iconv.ConvertString(str, "utf-8", "iso-8859-1")
		//str, _ = iconv.ConvertString(str, "iso-8859-1", "utf-8")
		fmt.Println(str)

		if str != "ZBar big law good! ZBar大法好!" {
			t.Fail()
		}
	})
}

func TestPhoto(t *testing.T) {
	f, err := os.Open("/mnt/sdcard/Download/sessfo-qrcode.jpg")

	if err != nil {
		t.Fail()
		return
	}

	i, _ := jpeg.Decode(f)
	img := FromImage(i)

	s := NewScanner()
	s.SetConfig(0, CFG_ENABLE, 1)
	s.Scan(img)
	defer s.Destroy()

	img.First().Each(func(str string) {
		// Charset decoding
		//str, _ = iconv.ConvertString(str, "utf-8", "iso-8859-1")
		//str, _ = iconv.ConvertString(str, "iso-8859-1", "utf-8")
		fmt.Println(str)

		if str != "http://www.searchenginestrategies.com/sanfrancisco/share.html" {
			t.Fail()
		}
	})
}
