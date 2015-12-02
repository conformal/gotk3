package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
import "C"

import (
	"unsafe"
)

type EventButton struct {
	event *C.GdkEventButton
}

func (e *EventButton) native() *C.GdkEventButton {
	return e.event
}

func (e *EventButton) FromNative(ptr uintptr) {
	e.event = (*C.GdkEventButton)(unsafe.Pointer(ptr))
}

func (e *EventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(e.event))
}

func (e *EventButton) Pos() (x, y int) {
	x = int(e.native().x)
	y = int(e.native().y)
	return
}

/*
struct GdkEventMotion {
  GdkEventType type;
  GdkWindow *window;
  gint8 send_event;
  guint32 time;
  gdouble x;
  gdouble y;
  gdouble *axes;
  guint state;
  gint16 is_hint;
  GdkDevice *device;
  gdouble x_root, y_root;
};
*/
type EventMotion struct {
	event *C.GdkEventMotion
}

func (e *EventMotion) native() *C.GdkEventMotion {
	return e.event
}

func (e *EventMotion) FromNative(ptr uintptr) {
	e.event = (*C.GdkEventMotion)(unsafe.Pointer(ptr))
}

func (e *EventMotion) Pos() (x, y int) {
	x = int(e.native().x)
	y = int(e.native().y)
	return
}

func (e *EventMotion) State() int {
	return int(e.native().state)
}
