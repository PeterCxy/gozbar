// ZBar image bindings for golang.
// Read the ZBar documents for details
package zbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

import "image"
import "image/draw"
//import "image/color"
import "unsafe"
import "runtime"

type Image struct {
	image *C.zbar_image_t
	gray *image.Gray
}

// Create an ZBar image object from a Golang image.
// To scan the image, call a Scanner.
func FromImage(img image.Image) *Image {
	ret := new(Image)
	ret.image = C.zbar_image_create()
	bounds := img.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	// Create a grayscale image
	ret.gray = image.NewGray(bounds)
	draw.Draw(
		ret.gray, bounds, img,
		image.ZP, draw.Over)

	C.zbar_image_set_format(ret.image, C.ulong(0x30303859)) // Y800 (grayscale)
	C.zbar_image_set_size(ret.image, C.uint(w), C.uint(h))
	C.zbar_image_set_data(ret.image, unsafe.Pointer(&ret.gray.Pix[0]), C.ulong(len(ret.gray.Pix)), nil)

	// finalizer
	runtime.SetFinalizer(ret, (*Image).Destroy)

	return ret
}

// Get the first scanned symbol of this image.
// To iterate over the symbols, use Symbol.Each() function
func (this *Image) First() *Symbol {
	s := C.zbar_image_first_symbol(this.image)

	if s == nil {
		return nil
	}

	return &Symbol {
		symbol: s,
	}
}

// Destroy this object
func (this *Image) Destroy() {
	C.zbar_image_destroy(this.image)
}