// Copyright (c) 2014 Jakob Runge <sicarius@g4t3.de>
//
// This file originated from: http://opensource.conformal.com/
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
package gtkx

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include "gtkx.go.h"
import "C"
import (
	"errors"
	"github.com/conformal/gotk3/gdk"
	"github.com/conformal/gotk3/glib"
	"github.com/conformal/gotk3/gtk"
	"runtime"
	"strconv"
	"unsafe"
)

// We need the same nilPtrErr as ../gtk.go:
var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
  This package deals with Sockets and Plugs, which wrap the concept of GtkSocket and GtkPlug as described in [1,2].
  To deal with both, it is necessary to identify them via their C.Window,
  that acts as an id and motivates our WindowId type.
  [1]: https://developer.gnome.org/gtk3/stable/GtkSocket.html
  [2]: https://developer.gnome.org/gtk3/stable/GtkPlug.html
*/
type (
	// Socket is a representation of GTK's GtkSocket.
	Socket struct {
		gtk.Container
	}

	// Plug is a representation of GTK's GtkPlug.
	Plug struct {
		gtk.Window
	}

	/*
	   The windowId that gtk uses, which is accessible as C.Window is a word32.
	   We handle this by using a typealias.
	*/
	WindowId uint32
)

// Native returns a pointer to the underlying GtkSocket.
func (v *Socket) Native() *C.GtkSocket {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSocket(p)
}

func wrapSocket(obj *glib.Object) *Socket {
	return &Socket{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}
}

// SocketNew is a wrapper around gtk_socket_new().
func SocketNew() (*Socket, error) {
	c := C.gtk_socket_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSocket(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// AddId is a wrapper around gtk_socket_add_id().
func (v *Socket) AddId(w WindowId) {
	C.gtk_socket_add_id(v.Native(), C.Window(w))
}

// GetId is a wrapper around gtk_socket_get_id().
func (v *Socket) GetId() WindowId {
	id := C.gtk_socket_get_id(v.Native())
	return WindowId(id)
}

// GetPlugWindow is a wrapper around gtk_socket_get_plug_window().
func (v *Socket) GetPlugWindow() (*gdk.Window, error) {
	c := C.gtk_socket_get_plug_window(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := &gdk.Window{obj}
	w.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

func (v *Plug) Native() *C.GtkPlug {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkPlug(p)
}

func wrapPlug(obj *glib.Object) *Plug {
	return &Plug{gtk.Window{gtk.Bin{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}}}
}

// PlugNew is a wrapper around gtk_plug_new().
func PlugNew(socketId WindowId) (*Plug, error) {
	c := C.gtk_plug_new(C.Window(socketId))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := wrapPlug(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// GetId is a wrapper around gtk_plug_get_id().
func (v *Plug) GetId() WindowId {
	id := C.gtk_plug_get_id(v.Native())
	return WindowId(id)
}

// GetEmbedded is a wrapper around gtk_plug_get_embedded().
func (v *Plug) GetEmbedded() bool {
	c := C.gtk_plug_get_embedded(v.Native())
	//Sadly we don't have gobool() outside gtk.go:
	if c != 0 {
		return true
	}
	return false
}

// GetSocketWindow is a wrapper around gtk_plug_get_socket_window().
func (v *Plug) GetSocketWindow() (*gdk.Window, error) {
	c := C.gtk_plug_get_socket_window(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := &gdk.Window{obj}
	w.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

func WindowIdToString(w WindowId) string {
	u := uint32(w)
	return strconv.FormatUint(uint64(u), 10)
}

func StringToWindowId(s string) (WindowId, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return WindowId(u), nil
}
