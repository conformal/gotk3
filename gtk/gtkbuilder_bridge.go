// Copyright (c) 2014 Joan Garcia i Silano - https://github.com/joangs
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
package gtk

/*
#cgo pkg-config: gtk+-3.0
#include <stdlib.h>
#include <gtk/gtk.h>

*/
import "C"
import (
	"github.com/conformal/gotk3/glib"
	"log"
	"reflect"
	"strings"
	"unsafe"
)

//export dispatchEvent
func dispatchEvent(pbuilder unsafe.Pointer, csignal *C.char, rval *C.GValue, nparams C.guint, size C.int, params *C.GValue) {
	b := (*Builder)(pbuilder)
	signal := C.GoString(csignal)
	signal = strings.ToUpper(signal[:1]) + signal[1:]

	values := []reflect.Value{}

	for c := uint64(0); c < (uint64)(nparams); c++ {
		addr := uintptr(unsafe.Pointer(params)) + uintptr(c*uint64(size))
		v := (*C.GValue)(unsafe.Pointer(addr))

		t := glib.Type(v.g_type)
		if m, err := t.GetMarshaller(); err != nil {
			log.Printf("Error: %s", err.Error())
			return
		} else {
			if gov, err := m(uintptr(unsafe.Pointer(v))); err != nil {
				log.Printf("Error: %s", err.Error())
				return
			} else {
				values = append(values, reflect.ValueOf(gov))
			}
		}
	}

	log.Printf("%s(%v) %d", signal, values, nparams)

	for _, iface := range b.callbacks {
		v := reflect.ValueOf(iface)
		m := v.MethodByName(signal)

		if m.IsValid() {
			m.Call(values)
			return
		}
	}

	log.Printf("Warning: Signal '%s' not defined", signal)
}
