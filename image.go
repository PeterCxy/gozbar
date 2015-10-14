// ZBar image bindings for golang.
// Read the ZBar documents for details
package zbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

import "image"
import "image/color"
import "unsafe"
import "runtime"

type Image struct {
	image *C.zbar_image_t
	raw []uint8
}

// Create an ZBar image object from a Golang image.
// To scan the image, call a Scanner.
func FromImage(image image.Image) *Image {
	ret := new(Image)
	ret.image = C.zbar_image_create()
	bounds := image.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	C.zbar_image_set_size(ret.image, C.uint(w), C.uint(h))

	ret.raw = make([]uint8, w * h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			m := color.GrayModel.Convert(image.At(x, y)).(color.Gray)

			ret.raw[y * w + x] = m.Y
		}
	}

	C.zbar_image_set_format(ret.image, 0x30303859) // Y800 (grayscale)
	C.zbar_image_set_data(ret.image, unsafe.Pointer(&ret.raw[0]), C.ulong(w * h), nil)

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
