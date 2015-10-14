// ZBar symbols
package zbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

type Symbol struct {
	symbol *C.zbar_symbol_t
}

// Get the next symbol
func (this *Symbol) Next() *Symbol {
	n := C.zbar_symbol_next(this.symbol)

	if n == nil {
		return nil
	}

	return &Symbol {
		symbol: n,
	}
}

// Get the data of this symbol
func (this *Symbol) Data() string {
	s := C.zbar_symbol_get_data(this.symbol)

	if s == nil {
		return ""
	}

	return C.GoString(s)
}

// Iterate over all symbols after this symbol.
// f will be called with each symbol's data as the argument
func (this *Symbol) Each(f func(string)) {
	t := this

	for {
		f(t.Data())

		t = t.Next()

		if t == nil {
			break
		}
	}
}
