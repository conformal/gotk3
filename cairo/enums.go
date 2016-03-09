package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

type Format int // cairo_format_t

const (
	FORMAT_INVALID   Format = C.CAIRO_FORMAT_INVALID
	FORMAT_ARGB32    Format = C.CAIRO_FORMAT_ARGB32
	FORMAT_RGB24     Format = C.CAIRO_FORMAT_RGB24
	FORMAT_A8        Format = C.CAIRO_FORMAT_A8
	FORMAT_A1        Format = C.CAIRO_FORMAT_A1
	FORMAT_RGB16_565 Format = C.CAIRO_FORMAT_RGB16_565
	FORMAT_RGB30     Format = C.CAIRO_FORMAT_RGB30
)
