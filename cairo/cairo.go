// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
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

// Package cairo implements Go bindings for Cairo.  Supports version 1.10 and
// later.
package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"

	"github.com/envoker/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.cairo_gobject_antialias_get_type()), marshalAntialias},
		{glib.Type(C.cairo_gobject_content_get_type()), marshalContent},
		{glib.Type(C.cairo_gobject_fill_rule_get_type()), marshalFillRule},
		{glib.Type(C.cairo_gobject_line_cap_get_type()), marshalLineCap},
		{glib.Type(C.cairo_gobject_line_join_get_type()), marshalLineJoin},
		{glib.Type(C.cairo_gobject_operator_get_type()), marshalOperator},
		{glib.Type(C.cairo_gobject_status_get_type()), marshalStatus},
		{glib.Type(C.cairo_gobject_surface_type_get_type()), marshalSurfaceType},

		// Boxed
		{glib.Type(C.cairo_gobject_context_get_type()), marshalContext},
		{glib.Type(C.cairo_gobject_surface_get_type()), marshalSurface},
	}
	glib.RegisterGValueMarshalers(tm)
}

// Constants

// Antialias is a representation of Cairo's cairo_antialias_t.
type Antialias int

const (
	ANTIALIAS_DEFAULT  Antialias = C.CAIRO_ANTIALIAS_DEFAULT
	ANTIALIAS_NONE     Antialias = C.CAIRO_ANTIALIAS_NONE
	ANTIALIAS_GRAY     Antialias = C.CAIRO_ANTIALIAS_GRAY
	ANTIALIAS_SUBPIXEL Antialias = C.CAIRO_ANTIALIAS_SUBPIXEL
	// ANTIALIAS_FAST     Antialias = C.CAIRO_ANTIALIAS_FAST (since 1.12)
	// ANTIALIAS_GOOD     Antialias = C.CAIRO_ANTIALIAS_GOOD (since 1.12)
	// ANTIALIAS_BEST     Antialias = C.CAIRO_ANTIALIAS_BEST (since 1.12)
)

func marshalAntialias(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Antialias(c), nil
}

// Content is a representation of Cairo's cairo_content_t.
type Content int

const (
	CONTENT_COLOR       Content = C.CAIRO_CONTENT_COLOR
	CONTENT_ALPHA       Content = C.CAIRO_CONTENT_ALPHA
	CONTENT_COLOR_ALPHA Content = C.CAIRO_CONTENT_COLOR_ALPHA
)

func marshalContent(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Content(c), nil
}

// FillRule is a representation of Cairo's cairo_fill_rule_t.
type FillRule int

const (
	FILL_RULE_WINDING  FillRule = C.CAIRO_FILL_RULE_WINDING
	FILL_RULE_EVEN_ODD FillRule = C.CAIRO_FILL_RULE_EVEN_ODD
)

func marshalFillRule(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return FillRule(c), nil
}

// LineCap is a representation of Cairo's cairo_line_cap_t.
type LineCap int

const (
	LINE_CAP_BUTT   LineCap = C.CAIRO_LINE_CAP_BUTT
	LINE_CAP_ROUND  LineCap = C.CAIRO_LINE_CAP_ROUND
	LINE_CAP_SQUARE LineCap = C.CAIRO_LINE_CAP_SQUARE
)

func marshalLineCap(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LineCap(c), nil
}

// LineJoin is a representation of Cairo's cairo_line_join_t.
type LineJoin int

const (
	LINE_JOIN_MITER LineJoin = C.CAIRO_LINE_JOIN_MITER
	LINE_JOIN_ROUND LineJoin = C.CAIRO_LINE_JOIN_ROUND
	LINE_JOIN_BEVEL LineJoin = C.CAIRO_LINE_JOIN_BEVEL
)

func marshalLineJoin(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LineJoin(c), nil
}

// MimeType is a representation of Cairo's CAIRO_MIME_TYPE_*
// preprocessor constants.
type MimeType string

const (
	MIME_TYPE_JP2       MimeType = "image/jp2"
	MIME_TYPE_JPEG      MimeType = "image/jpeg"
	MIME_TYPE_PNG       MimeType = "image/png"
	MIME_TYPE_URI       MimeType = "image/x-uri"
	MIME_TYPE_UNIQUE_ID MimeType = "application/x-cairo.uuid"
)

// Operator is a representation of Cairo's cairo_operator_t.
type Operator int

const (
	OPERATOR_CLEAR          Operator = C.CAIRO_OPERATOR_CLEAR
	OPERATOR_SOURCE         Operator = C.CAIRO_OPERATOR_SOURCE
	OPERATOR_OVER           Operator = C.CAIRO_OPERATOR_OVER
	OPERATOR_IN             Operator = C.CAIRO_OPERATOR_IN
	OPERATOR_OUT            Operator = C.CAIRO_OPERATOR_OUT
	OPERATOR_ATOP           Operator = C.CAIRO_OPERATOR_ATOP
	OPERATOR_DEST           Operator = C.CAIRO_OPERATOR_DEST
	OPERATOR_DEST_OVER      Operator = C.CAIRO_OPERATOR_DEST_OVER
	OPERATOR_DEST_IN        Operator = C.CAIRO_OPERATOR_DEST_IN
	OPERATOR_DEST_OUT       Operator = C.CAIRO_OPERATOR_DEST_OUT
	OPERATOR_DEST_ATOP      Operator = C.CAIRO_OPERATOR_DEST_ATOP
	OPERATOR_XOR            Operator = C.CAIRO_OPERATOR_XOR
	OPERATOR_ADD            Operator = C.CAIRO_OPERATOR_ADD
	OPERATOR_SATURATE       Operator = C.CAIRO_OPERATOR_SATURATE
	OPERATOR_MULTIPLY       Operator = C.CAIRO_OPERATOR_MULTIPLY
	OPERATOR_SCREEN         Operator = C.CAIRO_OPERATOR_SCREEN
	OPERATOR_OVERLAY        Operator = C.CAIRO_OPERATOR_OVERLAY
	OPERATOR_DARKEN         Operator = C.CAIRO_OPERATOR_DARKEN
	OPERATOR_LIGHTEN        Operator = C.CAIRO_OPERATOR_LIGHTEN
	OPERATOR_COLOR_DODGE    Operator = C.CAIRO_OPERATOR_COLOR_DODGE
	OPERATOR_COLOR_BURN     Operator = C.CAIRO_OPERATOR_COLOR_BURN
	OPERATOR_HARD_LIGHT     Operator = C.CAIRO_OPERATOR_HARD_LIGHT
	OPERATOR_SOFT_LIGHT     Operator = C.CAIRO_OPERATOR_SOFT_LIGHT
	OPERATOR_DIFFERENCE     Operator = C.CAIRO_OPERATOR_DIFFERENCE
	OPERATOR_EXCLUSION      Operator = C.CAIRO_OPERATOR_EXCLUSION
	OPERATOR_HSL_HUE        Operator = C.CAIRO_OPERATOR_HSL_HUE
	OPERATOR_HSL_SATURATION Operator = C.CAIRO_OPERATOR_HSL_SATURATION
	OPERATOR_HSL_COLOR      Operator = C.CAIRO_OPERATOR_HSL_COLOR
	OPERATOR_HSL_LUMINOSITY Operator = C.CAIRO_OPERATOR_HSL_LUMINOSITY
)

func marshalOperator(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Operator(c), nil
}

// Status is a representation of Cairo's cairo_status_t.
type Status int

const (
	STATUS_SUCCESS                   Status = C.CAIRO_STATUS_SUCCESS
	STATUS_NO_MEMORY                 Status = C.CAIRO_STATUS_NO_MEMORY
	STATUS_INVALID_RESTORE           Status = C.CAIRO_STATUS_INVALID_RESTORE
	STATUS_INVALID_POP_GROUP         Status = C.CAIRO_STATUS_INVALID_POP_GROUP
	STATUS_NO_CURRENT_POINT          Status = C.CAIRO_STATUS_NO_CURRENT_POINT
	STATUS_INVALID_MATRIX            Status = C.CAIRO_STATUS_INVALID_MATRIX
	STATUS_INVALID_STATUS            Status = C.CAIRO_STATUS_INVALID_STATUS
	STATUS_NULL_POINTER              Status = C.CAIRO_STATUS_NULL_POINTER
	STATUS_INVALID_STRING            Status = C.CAIRO_STATUS_INVALID_STRING
	STATUS_INVALID_PATH_DATA         Status = C.CAIRO_STATUS_INVALID_PATH_DATA
	STATUS_READ_ERROR                Status = C.CAIRO_STATUS_READ_ERROR
	STATUS_WRITE_ERROR               Status = C.CAIRO_STATUS_WRITE_ERROR
	STATUS_SURFACE_FINISHED          Status = C.CAIRO_STATUS_SURFACE_FINISHED
	STATUS_SURFACE_TYPE_MISMATCH     Status = C.CAIRO_STATUS_SURFACE_TYPE_MISMATCH
	STATUS_PATTERN_TYPE_MISMATCH     Status = C.CAIRO_STATUS_PATTERN_TYPE_MISMATCH
	STATUS_INVALID_CONTENT           Status = C.CAIRO_STATUS_INVALID_CONTENT
	STATUS_INVALID_FORMAT            Status = C.CAIRO_STATUS_INVALID_FORMAT
	STATUS_INVALID_VISUAL            Status = C.CAIRO_STATUS_INVALID_VISUAL
	STATUS_FILE_NOT_FOUND            Status = C.CAIRO_STATUS_FILE_NOT_FOUND
	STATUS_INVALID_DASH              Status = C.CAIRO_STATUS_INVALID_DASH
	STATUS_INVALID_DSC_COMMENT       Status = C.CAIRO_STATUS_INVALID_DSC_COMMENT
	STATUS_INVALID_INDEX             Status = C.CAIRO_STATUS_INVALID_INDEX
	STATUS_CLIP_NOT_REPRESENTABLE    Status = C.CAIRO_STATUS_CLIP_NOT_REPRESENTABLE
	STATUS_TEMP_FILE_ERROR           Status = C.CAIRO_STATUS_TEMP_FILE_ERROR
	STATUS_INVALID_STRIDE            Status = C.CAIRO_STATUS_INVALID_STRIDE
	STATUS_FONT_TYPE_MISMATCH        Status = C.CAIRO_STATUS_FONT_TYPE_MISMATCH
	STATUS_USER_FONT_IMMUTABLE       Status = C.CAIRO_STATUS_USER_FONT_IMMUTABLE
	STATUS_USER_FONT_ERROR           Status = C.CAIRO_STATUS_USER_FONT_ERROR
	STATUS_NEGATIVE_COUNT            Status = C.CAIRO_STATUS_NEGATIVE_COUNT
	STATUS_INVALID_CLUSTERS          Status = C.CAIRO_STATUS_INVALID_CLUSTERS
	STATUS_INVALID_SLANT             Status = C.CAIRO_STATUS_INVALID_SLANT
	STATUS_INVALID_WEIGHT            Status = C.CAIRO_STATUS_INVALID_WEIGHT
	STATUS_INVALID_SIZE              Status = C.CAIRO_STATUS_INVALID_SIZE
	STATUS_USER_FONT_NOT_IMPLEMENTED Status = C.CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED
	STATUS_DEVICE_TYPE_MISMATCH      Status = C.CAIRO_STATUS_DEVICE_TYPE_MISMATCH
	STATUS_DEVICE_ERROR              Status = C.CAIRO_STATUS_DEVICE_ERROR
	// STATUS_INVALID_MESH_CONSTRUCTION Status = C.CAIRO_STATUS_INVALID_MESH_CONSTRUCTION (since 1.12)
	// STATUS_DEVICE_FINISHED           Status = C.CAIRO_STATUS_DEVICE_FINISHED (since 1.12)
)

var key_Status = map[Status]string{

	STATUS_SUCCESS:                   "CAIRO_STATUS_SUCCESS",
	STATUS_NO_MEMORY:                 "CAIRO_STATUS_NO_MEMORY",
	STATUS_INVALID_RESTORE:           "CAIRO_STATUS_INVALID_RESTORE",
	STATUS_INVALID_POP_GROUP:         "CAIRO_STATUS_INVALID_POP_GROUP",
	STATUS_NO_CURRENT_POINT:          "CAIRO_STATUS_NO_CURRENT_POINT",
	STATUS_INVALID_MATRIX:            "CAIRO_STATUS_INVALID_MATRIX",
	STATUS_INVALID_STATUS:            "CAIRO_STATUS_INVALID_STATUS",
	STATUS_NULL_POINTER:              "CAIRO_STATUS_NULL_POINTER",
	STATUS_INVALID_STRING:            "CAIRO_STATUS_INVALID_STRING",
	STATUS_INVALID_PATH_DATA:         "CAIRO_STATUS_INVALID_PATH_DATA",
	STATUS_READ_ERROR:                "CAIRO_STATUS_READ_ERROR",
	STATUS_WRITE_ERROR:               "CAIRO_STATUS_WRITE_ERROR",
	STATUS_SURFACE_FINISHED:          "CAIRO_STATUS_SURFACE_FINISHED",
	STATUS_SURFACE_TYPE_MISMATCH:     "CAIRO_STATUS_SURFACE_TYPE_MISMATCH",
	STATUS_PATTERN_TYPE_MISMATCH:     "CAIRO_STATUS_PATTERN_TYPE_MISMATCH",
	STATUS_INVALID_CONTENT:           "CAIRO_STATUS_INVALID_CONTENT",
	STATUS_INVALID_FORMAT:            "CAIRO_STATUS_INVALID_FORMAT",
	STATUS_INVALID_VISUAL:            "CAIRO_STATUS_INVALID_VISUAL",
	STATUS_FILE_NOT_FOUND:            "CAIRO_STATUS_FILE_NOT_FOUND",
	STATUS_INVALID_DASH:              "CAIRO_STATUS_INVALID_DASH",
	STATUS_INVALID_DSC_COMMENT:       "CAIRO_STATUS_INVALID_DSC_COMMENT",
	STATUS_INVALID_INDEX:             "CAIRO_STATUS_INVALID_INDEX",
	STATUS_CLIP_NOT_REPRESENTABLE:    "CAIRO_STATUS_CLIP_NOT_REPRESENTABLE",
	STATUS_TEMP_FILE_ERROR:           "CAIRO_STATUS_TEMP_FILE_ERROR",
	STATUS_INVALID_STRIDE:            "CAIRO_STATUS_INVALID_STRIDE",
	STATUS_FONT_TYPE_MISMATCH:        "CAIRO_STATUS_FONT_TYPE_MISMATCH",
	STATUS_USER_FONT_IMMUTABLE:       "CAIRO_STATUS_USER_FONT_IMMUTABLE",
	STATUS_USER_FONT_ERROR:           "CAIRO_STATUS_USER_FONT_ERROR",
	STATUS_NEGATIVE_COUNT:            "CAIRO_STATUS_NEGATIVE_COUNT",
	STATUS_INVALID_CLUSTERS:          "CAIRO_STATUS_INVALID_CLUSTERS",
	STATUS_INVALID_SLANT:             "CAIRO_STATUS_INVALID_SLANT",
	STATUS_INVALID_WEIGHT:            "CAIRO_STATUS_INVALID_WEIGHT",
	STATUS_INVALID_SIZE:              "CAIRO_STATUS_INVALID_SIZE",
	STATUS_USER_FONT_NOT_IMPLEMENTED: "CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED",
	STATUS_DEVICE_TYPE_MISMATCH:      "CAIRO_STATUS_DEVICE_TYPE_MISMATCH",
	STATUS_DEVICE_ERROR:              "CAIRO_STATUS_DEVICE_ERROR",
}

func StatusToString(status Status) string {

	s, ok := key_Status[status]
	if !ok {
		s = "CAIRO_STATUS_UNDEFINED"
	}

	return s
}

func marshalStatus(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Status(c), nil
}

// SurfaceType is a representation of Cairo's cairo_surface_type_t.
type SurfaceType int

const (
	SURFACE_TYPE_IMAGE          SurfaceType = C.CAIRO_SURFACE_TYPE_IMAGE
	SURFACE_TYPE_PDF            SurfaceType = C.CAIRO_SURFACE_TYPE_PDF
	SURFACE_TYPE_PS             SurfaceType = C.CAIRO_SURFACE_TYPE_PS
	SURFACE_TYPE_XLIB           SurfaceType = C.CAIRO_SURFACE_TYPE_XLIB
	SURFACE_TYPE_XCB            SurfaceType = C.CAIRO_SURFACE_TYPE_XCB
	SURFACE_TYPE_GLITZ          SurfaceType = C.CAIRO_SURFACE_TYPE_GLITZ
	SURFACE_TYPE_QUARTZ         SurfaceType = C.CAIRO_SURFACE_TYPE_QUARTZ
	SURFACE_TYPE_WIN32          SurfaceType = C.CAIRO_SURFACE_TYPE_WIN32
	SURFACE_TYPE_BEOS           SurfaceType = C.CAIRO_SURFACE_TYPE_BEOS
	SURFACE_TYPE_DIRECTFB       SurfaceType = C.CAIRO_SURFACE_TYPE_DIRECTFB
	SURFACE_TYPE_SVG            SurfaceType = C.CAIRO_SURFACE_TYPE_SVG
	SURFACE_TYPE_OS2            SurfaceType = C.CAIRO_SURFACE_TYPE_OS2
	SURFACE_TYPE_WIN32_PRINTING SurfaceType = C.CAIRO_SURFACE_TYPE_WIN32_PRINTING
	SURFACE_TYPE_QUARTZ_IMAGE   SurfaceType = C.CAIRO_SURFACE_TYPE_QUARTZ_IMAGE
	SURFACE_TYPE_SCRIPT         SurfaceType = C.CAIRO_SURFACE_TYPE_SCRIPT
	SURFACE_TYPE_QT             SurfaceType = C.CAIRO_SURFACE_TYPE_QT
	SURFACE_TYPE_RECORDING      SurfaceType = C.CAIRO_SURFACE_TYPE_RECORDING
	SURFACE_TYPE_VG             SurfaceType = C.CAIRO_SURFACE_TYPE_VG
	SURFACE_TYPE_GL             SurfaceType = C.CAIRO_SURFACE_TYPE_GL
	SURFACE_TYPE_DRM            SurfaceType = C.CAIRO_SURFACE_TYPE_DRM
	SURFACE_TYPE_TEE            SurfaceType = C.CAIRO_SURFACE_TYPE_TEE
	SURFACE_TYPE_XML            SurfaceType = C.CAIRO_SURFACE_TYPE_XML
	SURFACE_TYPE_SKIA           SurfaceType = C.CAIRO_SURFACE_TYPE_SKIA
	SURFACE_TYPE_SUBSURFACE     SurfaceType = C.CAIRO_SURFACE_TYPE_SUBSURFACE
	// SURFACE_TYPE_COGL           SurfaceType = C.CAIRO_SURFACE_TYPE_COGL (since 1.12)
)

func marshalSurfaceType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SurfaceType(c), nil
}
