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
	"github.com/conformal/gotk3/glib"
	"runtime"
	"strconv"
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
	PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
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
 * GdkPixbuf
 */

// Pixbuf is a representation of GDK's GdkPixbuf.
type Pixbuf struct {
	*glib.Object
}

func wrapPixbuf(obj *glib.Object) *Pixbuf {
	return &Pixbuf{obj}
}

// Native() returns a pointer to the underlying GdkPixbuf.
func (v *Pixbuf) Native() *C.GdkPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbuf(p)
}

// PixbufNewFromFile() is a wrapper around gdk_pixbuf_new_from_file().
func PixbufNewFromFile(fileName string) (*Pixbuf, error) {
	var cerr *C.GError
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_pixbuf_new_from_file(cstr, &cerr)
	if c == nil {
		defer C.g_error_free(cerr)
		errstr := C.GoString((*C.char)(C.error_get_message(cerr)))
		return nil, errors.New(errstr)
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPixbuf(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufNewFromFileAtSize() is a wrapper around
// gdk_pixbuf_new_from_file_at_size().
func PixbufNewFromFileAtSize(fileName string, width, height int) (*Pixbuf, error) {
	var cerr *C.GError
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_pixbuf_new_from_file_at_size(cstr, C.int(width),
		C.int(height), &cerr)
	if c == nil {
		defer C.g_error_free(cerr)
		errstr := C.GoString((*C.char)(C.error_get_message(cerr)))
		return nil, errors.New(errstr)
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPixbuf(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufNewFromFileAtScale() is a wrapper around
// gdk_pixbuf_new_from_file_at_scale().
func PixbufNewFromFileAtScale(fileName string, width, height int, preserveAspectRatio bool) (*Pixbuf, error) {
	var cerr *C.GError
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_pixbuf_new_from_file_at_scale(cstr, C.int(width),
		C.int(height), gbool(preserveAspectRatio), &cerr)
	if c == nil {
		defer C.g_error_free(cerr)
		errstr := C.GoString((*C.char)(C.error_get_message(cerr)))
		return nil, errors.New(errstr)
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPixbuf(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// ScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func (v *Pixbuf) ScaleSimple(width, height int, interp InterpType) (*Pixbuf, error) {
	c := C.gdk_pixbuf_scale_simple(v.Native(), C.int(width), C.int(height),
		C.GdkInterpType(interp))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPixbuf(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// Scale is a wrapper around gdk_pixbuf_scale().
func (v *Pixbuf) Scale(dst *Pixbuf, x, y, width, height int, offsetX, offsetY, scaleX, scaleY float64, interp InterpType) {
	C.gdk_pixbuf_scale(v.Native(), dst.Native(), C.int(x), C.int(y),
		C.int(width), C.int(height), C.double(offsetX),
		C.double(offsetY), C.double(scaleX), C.double(scaleY),
		C.GdkInterpType(interp))
}

// RotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func (v *Pixbuf) RotateSimple(angle PixbufRotation) (*Pixbuf, error) {
	c := C.gdk_pixbuf_rotate_simple(v.Native(), C.GdkPixbufRotation(angle))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPixbuf(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// Flip is a wrapper around gdk_pixbuf_flip().
func (v *Pixbuf) Flip(horizontal bool) (*Pixbuf, error) {
	c := C.gdk_pixbuf_flip(v.Native(), gbool(horizontal))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPixbuf(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// SavePNG is a wrapper around gdk_pixbuf_save().
// Compression is a number between 0...9
func (v *Pixbuf) SavePNG(path string, compression int) error {
	cpath := C.CString(path)
	ccompression := C.CString(strconv.Itoa(compression))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(ccompression))

	var cerr *C.GError
	c := C._gdk_pixbuf_save_png(v.Native(), cpath, &cerr, ccompression);
	if !gobool(c) {
		defer C.g_error_free(cerr)
		errstr := C.GoString((*C.char)(C.error_get_message(cerr)))
		return errors.New(errstr)
	}
	return nil
}

// SaveJPEG is a wrapper around gdk_pixbuf_save().
// Quality is a number between 0...100
func (v *Pixbuf) SaveJPEG(path string, quality int) error {
	cpath := C.CString(path)
	cquality := C.CString(strconv.Itoa(quality))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(cquality))

	var cerr *C.GError
	c := C._gdk_pixbuf_save_jpeg(v.Native(), cpath, &cerr, cquality);
	if !gobool(c) {
		defer C.g_error_free(cerr)
		errstr := C.GoString((*C.char)(C.error_get_message(cerr)))
		return errors.New(errstr)
	}
	return nil
}

// GetWidth is a wrapper around gdk_pixbuf_get_width().
func (v *Pixbuf) GetWidth() int {
	return int(C.gdk_pixbuf_get_width(v.Native()))
}

// GetHeight is a wrapper around gdk_pixbuf_get_height().
func (v *Pixbuf) GetHeight() int {
	return int(C.gdk_pixbuf_get_height(v.Native()))
}

// TODO gdk_pixbuf_get_file_info, resource and stream related functions.

/*
 * GdkPixbufLoader
 */

// PixbufLoader is a representation of GDK's GdkPixbufLoader.
// Users of PixbufLoader are expected to call Close() when they are finished.
type PixbufLoader struct {
	*glib.Object
}

// Native() returns a pointer to the underlying GdkPixbufLoader.
func (v *PixbufLoader) Native() *C.GdkPixbufLoader {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbufLoader(p)
}

func wrapPixbufLoader(obj *glib.Object) *PixbufLoader {
	return &PixbufLoader{obj}
}

// PixbufLoaderNew() is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() (*PixbufLoader, error) {
	c := C.gdk_pixbuf_loader_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapPixbufLoader(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l, nil
}

// Write() is a wrapper around gdk_pixbuf_loader_write().  The
// function signature differs from the C equivalent to satisify the
// io.Writer interface.
func (v *PixbufLoader) Write(data []byte) (n int, err error) {
	// n is set to 0 on error, and set to len(data) otherwise.
	// This is a tiny hacky to satisfy io.Writer and io.WriteCloser,
	// which would allow access to all io and ioutil goodies,
	// and play along nice with go environment.

	if len(data) == 0 {
		return 0, nil
	}

	var cerr *C.GError
	c := C.gdk_pixbuf_loader_write(v.Native(),
		(*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)),
		&cerr)
	if !gobool(c) {
		defer C.g_error_free(cerr)
		errstr := C.GoString((*C.char)(C.error_get_message(cerr)))
		return 0, errors.New(errstr)
	}

	return len(data), nil
}

// Close is a wrapper around gdk_pixbuf_loader_close().  An error is
// returned instead of a bool like the native C function to support the
// io.Closer interface.
func (v *PixbufLoader) Close() error {
	var cerr *C.GError

	if ok := gobool(C.gdk_pixbuf_loader_close(v.Native(), &cerr)); !ok {
		defer C.g_error_free(cerr)
		errstr := C.GoString((*C.char)(C.error_get_message(cerr)))
		return errors.New(errstr)
	}
	return nil
}

// SetSize is a wrapper around gdk_pixbuf_loader_set_size().
func (v *PixbufLoader) SetSize(width, height int) {
	C.gdk_pixbuf_loader_set_size(v.Native(), C.int(width), C.int(height))
}

// GetPixbuf is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (v *PixbufLoader) GetPixbuf() (*Pixbuf, error) {
	c := C.gdk_pixbuf_loader_get_pixbuf(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPixbuf(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// TODO GdkPixbufFormat, GdkPixbufAnimation,
// gdk_pixbuf_loader_new_with_type, gdk_pixbuf_loader_new_with_mime_type
