package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
import "C"

import (
	"unsafe"
)

/*
 * GdkEvent
 */

// Event is a representation of GDK's GdkEvent.
type Event struct {
	GdkEvent *C.GdkEvent
}

// native returns a pointer to the underlying GdkEvent.
func (v *Event) native() *C.GdkEvent {
	if v == nil {
		return nil
	}
	return v.GdkEvent
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Event) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalEvent(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Event{(*C.GdkEvent)(unsafe.Pointer(c))}, nil
}

func (v *Event) free() {
	C.gdk_event_free(v.native())
}

/*
 * GdkEventKey
 */

// EventKey is a representation of GDK's GdkEventKey.
type EventKey struct {
	*Event
}

// Native returns a pointer to the underlying GdkEventKey.
func (v *EventKey) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventKey) native() *C.GdkEventKey {
	return (*C.GdkEventKey)(unsafe.Pointer(v.Event.native()))
}

func (v *EventKey) KeyVal() uint {
	c := v.native().keyval
	return uint(c)
}

/*
 * GdkEventButton
 */

type EventButton struct {
	*Event
}

func (e *EventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(e.native()))
}

func (e *EventButton) native() *C.GdkEventButton {
	return (*C.GdkEventButton)(unsafe.Pointer(e.Event.native()))
}

func (e *EventButton) Pos() (x, y int) {
	x = int(e.native().x)
	y = int(e.native().y)
	return
}

/*
 * GdkEventMotion
 */

type EventMotion struct {
	*Event
}

func (e *EventMotion) Native() uintptr {
	return uintptr(unsafe.Pointer(e.native()))
}

func (e *EventMotion) native() *C.GdkEventMotion {
	return (*C.GdkEventMotion)(unsafe.Pointer(e.Event.native()))
}

func (e *EventMotion) Pos() (x, y int) {
	x = int(e.native().x)
	y = int(e.native().y)
	return
}

func (e *EventMotion) State() int {
	return int(e.native().state)
}
