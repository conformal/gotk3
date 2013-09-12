/*
 * Copyright (c) 2013 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
Go bindings for GDK-PixBuf 3.  Supports version 2.0 and later.
*/
package gdkpixbuf

// #cgo pkg-config: gdk-pixbuf-2.0
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include "gdkpixbuf.go.h"
import "C"
import (
	"errors"
	"github.com/conformal/gotk3/gdk"
	"github.com/conformal/gotk3/glib"
	"runtime"
	"unsafe"
)

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constants
 */

type InterpType C.GdkInterpType

const (
	INTERP_NEAREST InterpType = iota
	INTERP_TILES
	INTERP_BILINEAR
	INTERP_HYPER
)

type PixbufRotation C.GdkPixbufRotation

const (
	PIXBUF_ROTATE_NONE             PixbufRotation = 0
	PIXBUF_ROTATE_COUNTERCLOCKWISE                = 90
	PIXBUF_ROTATE_UPSIDEDOWN                      = 180
	PIXBUF_ROTATE_CLOCKWISE                       = 270
)

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

/*
 * GdkPixbufLoader
 */

// PixbufLoader is a representation of GDK's GdkPixbufLoader.
// Users of PixbufLoader are expected to call Close() when they are finished.
type PixbufLoader struct {
	*glib.Object
}

// Native() returns a pointer to the underlying GdkWindow.
func (v *PixbufLoader) Native() *C.GdkPixbufLoader {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbufLoader(p)
}

// PixbufLoaderNew() is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() (*PixbufLoader, error) {
	c := C.gdk_pixbuf_loader_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := &PixbufLoader{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)

	return l, nil
}

// Write() is a wrapper around gdk_pixbuf_loader_write().
func (v *PixbufLoader) Write(data []byte) (n int, err error) {
	// n is set to 0 on error, and set to len(data) otherwise.
	// This is a tiny hacky to satisfy io.Writer and io.WriteCloser,
	// which would allow access to all io and ioutil goodies,
	// and play along nice with go environment.

	if len(data) == 0 {
		return 0, nil
	}

	var cerr *C.GError
	ok := gobool(C.gdk_pixbuf_loader_write(v.Native(), (*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)), &cerr))
	if !ok {
		defer C.g_error_free(cerr)
		return 0, errors.New(C.GoString((*C.char)(C.error_get_message(cerr))))
	}

	return len(data), nil
}

// Close() is a wrapper around gdk_pixbuf_loader_close().
func (v *PixbufLoader) Close() error {
	var cerr *C.GError

	if ok := gobool(C.gdk_pixbuf_loader_close(v.Native(), &cerr)); !ok {
		defer C.g_error_free(cerr)
		return errors.New(C.GoString((*C.char)(C.error_get_message(cerr))))
	}
	return nil
}

// SetSize() is a wrapper around gdk_pixbuf_loader_set_size().
func (v *PixbufLoader) SetSize(width, height int) {
	C.gdk_pixbuf_loader_set_size(v.Native(), C.int(width), C.int(height))
}

// GetPixbuf() is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (v *PixbufLoader) GetPixbuf() (*gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_loader_get_pixbuf(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &gdk.Pixbuf{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// TODO GdkPixbufFormat, GdkPixbufAnimation,
// gdk_pixbuf_loader_new_with_type, gdk_pixbuf_loader_new_with_mime_type

// PixbufScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func PixbufScaleSimple(src *gdk.Pixbuf, width, height int, interp InterpType) (*gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_scale_simple(src.Native(), C.int(width), C.int(height), C.GdkInterpType(interp))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &gdk.Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufScale is a wrapper around gdk_pixbuf_scale().
func PixbufScale(src, dst *gdk.Pixbuf, x, y, width, height int, offsetX, offsetY, scaleX, scaleY float64, interp InterpType) {
	C.gdk_pixbuf_scale(src.Native(), dst.Native(), C.int(x), C.int(y), C.int(width), C.int(height), C.double(offsetX), C.double(offsetY), C.double(scaleX), C.double(scaleY), C.GdkInterpType(interp))
}

// PixbufRotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func PixbufRotateSimple(src *gdk.Pixbuf, angle PixbufRotation) (*gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_rotate_simple(src.Native(), C.GdkPixbufRotation(angle))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &gdk.Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufFlip is a wrapper around gdk_pixbuf_flip().
func PixbufFlip(src *gdk.Pixbuf, horizontal bool) (*gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_flip(src.Native(), gbool(horizontal))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &gdk.Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}
