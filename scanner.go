// ZBar scanner binding for golang
// Install ZBar library and headers first
package zbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

import "runtime"

type Scanner struct {
	scanner *C.zbar_image_scanner_t
}

func NewScanner() *Scanner {
	r := &Scanner{
		scanner: C.zbar_image_scanner_create(),
	}
	runtime.SetFinalizer(r, (*Scanner).Destroy)
	return r
}

// Set configuration.
// Read the ZBar docs for details.
func (this *Scanner) SetConfig(symbology C.zbar_symbol_type_t, config C.zbar_config_t, value int) *Scanner {
	C.zbar_image_scanner_set_config(this.scanner, symbology, config, C.int(value))
	return this
}

// Scan image. Create an Image object first.
// To get the result, call Image.First() and Symbol.Each()
// >0: Successful
// 0: No symbols found
// -1: error occurred
func (this *Scanner) Scan(img *Image) int {
	return int(C.zbar_scan_image(this.scanner, img.image))
}

// Destroy this Scanner
func (this *Scanner) Destroy() {
	C.zbar_image_scanner_destroy(this.scanner)
}
