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

// #cgo pkg-config: cairo
// #include <stdlib.h>
// #include <cairo.h>
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

// Type conversions

func cairobool(b bool) C.cairo_bool_t {
	if b {
		return C.cairo_bool_t(1)
	}
	return C.cairo_bool_t(0)
}

func gobool(b C.cairo_bool_t) bool {
	if b != 0 {
		return true
	}
	return false
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

// Content is a representation of Cairo's cairo_content_t.
type Content int

const (
	CONTENT_COLOR       Content = C.CAIRO_CONTENT_COLOR
	CONTENT_ALPHA       Content = C.CAIRO_CONTENT_ALPHA
	CONTENT_COLOR_ALPHA Content = C.CAIRO_CONTENT_COLOR_ALPHA
)

// FillRule is a representation of Cairo's cairo_fill_rule_t.
type FillRule int

const (
	FILL_RULE_WINDING  FillRule = C.CAIRO_FILL_RULE_WINDING
	FILL_RULE_EVEN_ODD FillRule = C.CAIRO_FILL_RULE_EVEN_ODD
)

// LineCap is a representation of Cairo's cairo_line_cap_t.
type LineCap int

const (
	LINE_CAP_BUTT   LineCap = C.CAIRO_LINE_CAP_BUTT
	LINE_CAP_ROUND  LineCap = C.CAIRO_LINE_CAP_ROUND
	LINE_CAP_SQUARE LineCap = C.CAIRO_LINE_CAP_SQUARE
)

// LineJoin is a representation of Cairo's cairo_line_join_t.
type LineJoin int

const (
	LINE_JOIN_MITER LineJoin = C.CAIRO_LINE_JOIN_MITER
	LINE_JOIN_ROUND LineJoin = C.CAIRO_LINE_JOIN_ROUND
	LINE_JOIN_BEVEL LineJoin = C.CAIRO_LINE_JOIN_BEVEL
)

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

/*
 * cairo_t
 */

// Context is a representation of Cairo's cairo_t.
type Context struct {
	context *C.cairo_t
}

// Native returns a pointer to the underlying cairo_t.
func (v *Context) Native() *C.cairo_t {
	if v == nil {
		return nil
	}
	return v.context
}

func wrapContext(context *C.cairo_t) *Context {
	return &Context{context}
}

// Create is a wrapper around cairo_create().
func Create(target *Surface) *Context {
	c := C.cairo_create(target.Native())
	ctx := wrapContext(c)
	runtime.SetFinalizer(ctx, (*Context).destroy)
	return ctx
}

// reference is a wrapper around cairo_reference().
func (v *Context) reference() {
	v.context = C.cairo_reference(v.Native())
}

// destroy is a wrapper around cairo_destroy().
func (v *Context) destroy() {
	C.cairo_destroy(v.Native())
}

// Status is a wrapper around cairo_status().
func (v *Context) Status() Status {
	c := C.cairo_status(v.Native())
	return Status(c)
}

// Save is a wrapper around cairo_save().
func (v *Context) Save() {
	C.cairo_save(v.Native())
}

// Restore is a wrapper around cairo_restore().
func (v *Context) Restore() {
	C.cairo_restore(v.Native())
}

// GetTarget is a wrapper around cairo_get_target().
func (v *Context) GetTarget() *Surface {
	c := C.cairo_get_target(v.Native())
	s := wrapSurface(c)
	s.reference()
	runtime.SetFinalizer(s, (*Surface).destroy)
	return s
}

// PushGroup is a wrapper around cairo_push_group().
func (v *Context) PushGroup() {
	C.cairo_push_group(v.Native())
}

// PushGroupWithContent is a wrapper around cairo_push_group_with_content().
func (v *Context) PushGroupWithContent(content Content) {
	C.cairo_push_group_with_content(v.Native(), C.cairo_content_t(content))
}

// TODO(jrick) PopGroup (depends on Pattern)

// PopGroupToSource is a wrapper around cairo_pop_group_to_source().
func (v *Context) PopGroupToSource() {
	C.cairo_pop_group_to_source(v.Native())
}

// GetGroupTarget is a wrapper around cairo_get_group_target().
func (v *Context) GetGroupTarget() *Surface {
	c := C.cairo_get_group_target(v.Native())
	s := wrapSurface(c)
	s.reference()
	runtime.SetFinalizer(s, (*Surface).destroy)
	return s
}

// SetSourceRGB is a wrapper around cairo_set_source_rgb().
func (v *Context) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(v.Native(), C.double(red), C.double(green),
		C.double(blue))
}

// SetSourceRGBA is a wrapper around cairo_set_source_rgba().
func (v *Context) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(v.Native(), C.double(red), C.double(green),
		C.double(blue), C.double(alpha))
}

// TODO(jrick) SetSource (depends on Pattern)

// SetSourceSurface is a wrapper around cairo_set_source_surface().
func (v *Context) SetSourceSurface(surface *Surface, x, y float64) {
	C.cairo_set_source_surface(v.Native(), surface.Native(), C.double(x),
		C.double(y))
}

// TODO(jrick) GetSource (depends on Pattern)

// SetAntialias is a wrapper around cairo_set_antialias().
func (v *Context) SetAntialias(antialias Antialias) {
	C.cairo_set_antialias(v.Native(), C.cairo_antialias_t(antialias))
}

// GetAntialias is a wrapper around cairo_get_antialias().
func (v *Context) GetAntialias() Antialias {
	c := C.cairo_get_antialias(v.Native())
	return Antialias(c)
}

// SetDash is a wrapper around cairo_set_dash().
func (v *Context) SetDash(dashes []float64, offset float64) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&dashes))
	cdashes := (*C.double)(unsafe.Pointer(header.Data))
	C.cairo_set_dash(v.Native(), cdashes, C.int(header.Len),
		C.double(offset))
}

// GetDashCount is a wrapper around cairo_get_dash_count().
func (v *Context) GetDashCount() int {
	c := C.cairo_get_dash_count(v.Native())
	return int(c)
}

// GetDash is a wrapper around cairo_get_dash().
func (v *Context) GetDash() (dashes []float64, offset float64) {
	dashCount := v.GetDashCount()
	cdashes := (*C.double)(C.calloc(8, C.size_t(dashCount)))
	var coffset C.double
	C.cairo_get_dash(v.Native(), cdashes, &coffset)
	header := (*reflect.SliceHeader)((unsafe.Pointer(&dashes)))
	header.Data = uintptr(unsafe.Pointer(cdashes))
	header.Len = dashCount
	header.Cap = dashCount
	return dashes, float64(coffset)
}

// SetFillRule is a wrapper around cairo_set_fill_rule().
func (v *Context) SetFillRule(fillRule FillRule) {
	C.cairo_set_fill_rule(v.Native(), C.cairo_fill_rule_t(fillRule))
}

// GetFillRule is a wrapper around cairo_get_fill_rule().
func (v *Context) GetFillRule() FillRule {
	c := C.cairo_get_fill_rule(v.Native())
	return FillRule(c)
}

// SetLineCap is a wrapper around cairo_set_line_cap().
func (v *Context) SetLineCap(lineCap LineCap) {
	C.cairo_set_line_cap(v.Native(), C.cairo_line_cap_t(lineCap))
}

// GetLineCap is a wrapper around cairo_get_line_cap().
func (v *Context) GetLineCap() LineCap {
	c := C.cairo_get_line_cap(v.Native())
	return LineCap(c)
}

// SetLineJoin is a wrapper around cairo_set_line_join().
func (v *Context) SetLineJoin(lineJoin LineJoin) {
	C.cairo_set_line_join(v.Native(), C.cairo_line_join_t(lineJoin))
}

// GetLineJoin is a wrapper around cairo_get_line_join().
func (v *Context) GetLineJoin() LineJoin {
	c := C.cairo_get_line_join(v.Native())
	return LineJoin(c)
}

// SetLineWidth is a wrapper around cairo_set_line_width().
func (v *Context) SetLineWidth(width float64) {
	C.cairo_set_line_width(v.Native(), C.double(width))
}

// GetLineWidth is a wrapper cairo_get_line_width().
func (v *Context) GetLineWidth() float64 {
	c := C.cairo_get_line_width(v.Native())
	return float64(c)
}

// SetMiterLimit is a wrapper around cairo_set_miter_limit().
func (v *Context) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(v.Native(), C.double(limit))
}

// GetMiterLimit is a wrapper around cairo_get_miter_limit().
func (v *Context) GetMiterLimit() float64 {
	c := C.cairo_get_miter_limit(v.Native())
	return float64(c)
}

// SetOperator is a wrapper around cairo_set_operator().
func (v *Context) SetOperator(op Operator) {
	C.cairo_set_operator(v.Native(), C.cairo_operator_t(op))
}

// GetOperator is a wrapper around cairo_get_operator().
func (v *Context) GetOperator() Operator {
	c := C.cairo_get_operator(v.Native())
	return Operator(c)
}

// SetTolerance is a wrapper around cairo_set_tolerance().
func (v *Context) SetTolerance(tolerance float64) {
	C.cairo_set_tolerance(v.Native(), C.double(tolerance))
}

// GetTolerance is a wrapper around cairo_get_tolerance().
func (v *Context) GetTolerance() float64 {
	c := C.cairo_get_tolerance(v.Native())
	return float64(c)
}

// Clip is a wrapper around cairo_clip().
func (v *Context) Clip() {
	C.cairo_clip(v.Native())
}

// ClipPreserve is a wrapper around cairo_clip_preserve().
func (v *Context) ClipPreserve() {
	C.cairo_clip_preserve(v.Native())
}

// ClipExtents is a wrapper around cairo_clip_extents().
func (v *Context) ClipExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 C.double
	C.cairo_clip_extents(v.Native(), &cx1, &cy1, &cx2, &cy2)
	return float64(cx1), float64(cy1), float64(cx2), float64(cy2)
}

// InClip is a wrapper around cairo_in_clip().
func (v *Context) InClip(x, y float64) bool {
	c := C.cairo_in_clip(v.Native(), C.double(x), C.double(y))
	return gobool(c)
}

// ResetClip is a wrapper around cairo_reset_clip().
func (v *Context) ResetClip() {
	C.cairo_reset_clip(v.Native())
}

// TODO(jrick) CopyRectangleList (depends on RectangleList)

// Fill is a wrapper around cairo_fill().
func (v *Context) Fill() {
	C.cairo_fill(v.Native())
}

// FillPreserve is a wrapper around cairo_fill_preserve().
func (v *Context) FillPreserve() {
	C.cairo_fill_preserve(v.Native())
}

// FillExtents is a wrapper around cairo_fill_extents().
func (v *Context) FillExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 C.double
	C.cairo_fill_extents(v.Native(), &cx1, &cy1, &cx2, &cy2)
	return float64(cx1), float64(cy1), float64(cx2), float64(cy2)
}

// InFill is a wrapper around cairo_in_fill().
func (v *Context) InFill(x, y float64) bool {
	c := C.cairo_in_fill(v.Native(), C.double(x), C.double(y))
	return gobool(c)
}

// TODO(jrick) Mask (depends on Pattern)

// MaskSurface is a wrapper around cairo_mask_surface().
func (v *Context) MaskSurface(surface *Surface, surfaceX, surfaceY float64) {
	C.cairo_mask_surface(v.Native(), surface.Native(), C.double(surfaceX),
		C.double(surfaceY))
}

// Paint is a wrapper around cairo_paint().
func (v *Context) Paint() {
	C.cairo_paint(v.Native())
}

// PaintWithAlpha is a wrapper around cairo_paint_with_alpha().
func (v *Context) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(v.Native(), C.double(alpha))
}

// Stroke is a wrapper around cairo_stroke().
func (v *Context) Stroke() {
	C.cairo_stroke(v.Native())
}

// StrokePreserve is a wrapper around cairo_stroke_preserve().
func (v *Context) StrokePreserve() {
	C.cairo_stroke_preserve(v.Native())
}

// StrokeExtents is a wrapper around cairo_stroke_extents().
func (v *Context) StrokeExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 C.double
	C.cairo_stroke_extents(v.Native(), &cx1, &cy1, &cx2, &cy2)
	return float64(cx1), float64(cy1), float64(cx2), float64(cy2)
}

// InStroke is a wrapper around cairo_in_stroke().
func (v *Context) InStroke(x, y float64) bool {
	c := C.cairo_in_stroke(v.Native(), C.double(x), C.double(y))
	return gobool(c)
}

// CopyPage is a wrapper around cairo_copy_page().
func (v *Context) CopyPage() {
	C.cairo_copy_page(v.Native())
}

// ShowPage is a wrapper around cairo_show_page().
func (v *Context) ShowPage() {
	C.cairo_show_page(v.Native())
}

// TODO(jrick) SetUserData (depends on UserDataKey and DestroyFunc)

// TODO(jrick) GetUserData (depends on UserDataKey)

/*
 * cairo_surface_t
 */

// Surface is a representation of Cairo's cairo_surface_t.
type Surface struct {
	surface *C.cairo_surface_t
}

// Native returns a pointer to the underlying cairo_surface_t.
func (v *Surface) Native() *C.cairo_surface_t {
	if v == nil {
		return nil
	}
	return v.surface
}

func wrapSurface(surface *C.cairo_surface_t) *Surface {
	return &Surface{surface}
}

// NewSurface creates a gotk3 cairo Surface from a pointer to a
// C cairo_surface_t.  This is primarily designed for use with other
// gotk3 packages and should be avoided by applications.
func NewSurface(s *C.cairo_surface_t, needsRef bool) *Surface {
	surface := wrapSurface(s)
	if needsRef {
		surface.reference()
	}
	runtime.SetFinalizer(surface, (*Surface).destroy)
	return surface
}

// CreateSimilar is a wrapper around cairo_surface_create_similar().
func (v *Surface) CreateSimilar(content Content, width, height int) *Surface {
	c := C.cairo_surface_create_similar(v.Native(),
		C.cairo_content_t(content), C.int(width), C.int(height))
	s := wrapSurface(c)
	runtime.SetFinalizer(s, (*Surface).destroy)
	return s
}

// TODO cairo_surface_create_similar_image (since 1.12)

// CreateForRectangle is a wrapper around cairo_surface_create_for_rectangle().
func (v *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	c := C.cairo_surface_create_for_rectangle(v.Native(), C.double(x),
		C.double(y), C.double(width), C.double(height))
	s := wrapSurface(c)
	runtime.SetFinalizer(s, (*Surface).destroy)
	return s
}

// reference is a wrapper around cairo_surface_reference().
func (v *Surface) reference() {
	v.surface = C.cairo_surface_reference(v.Native())
}

// destroy is a wrapper around cairo_surface_destroy().
func (v *Surface) destroy() {
	C.cairo_surface_destroy(v.Native())
}

// Status is a wrapper around cairo_surface_status().
func (v *Surface) Status() Status {
	c := C.cairo_surface_status(v.Native())
	return Status(c)
}

// Flush is a wrapper around cairo_surface_flush().
func (v *Surface) Flush() {
	C.cairo_surface_flush(v.Native())
}

// TODO(jrick) GetDevice (requires Device bindings)

// TODO(jrick) GetFontOptions (require FontOptions bindings)

// TODO(jrick) GetContent (requires Content bindings)

// MarkDirty is a wrapper around cairo_surface_mark_dirty().
func (v *Surface) MarkDirty() {
	C.cairo_surface_mark_dirty(v.Native())
}

// MarkDirtyRectangle is a wrapper around cairo_surface_mark_dirty_rectangle().
func (v *Surface) MarkDirtyRectangle(x, y, width, height int) {
	C.cairo_surface_mark_dirty_rectangle(v.Native(), C.int(x), C.int(y),
		C.int(width), C.int(height))
}

// SetDeviceOffset is a wrapper around cairo_surface_set_device_offset().
func (v *Surface) SetDeviceOffset(x, y float64) {
	C.cairo_surface_set_device_offset(v.Native(), C.double(x), C.double(y))
}

// GetDeviceOffset is a wrapper around cairo_surface_get_device_offset().
func (v *Surface) GetDeviceOffset() (x, y float64) {
	var xOffset, yOffset C.double
	C.cairo_surface_get_device_offset(v.Native(), &xOffset, &yOffset)
	return float64(xOffset), float64(yOffset)
}

// SetFallbackResolution is a wrapper around
// cairo_surface_set_fallback_resolution().
func (v *Surface) SetFallbackResolution(xPPI, yPPI float64) {
	C.cairo_surface_set_fallback_resolution(v.Native(), C.double(xPPI),
		C.double(yPPI))
}

// GetFallbackResolution is a wrapper around
// cairo_surface_get_fallback_resolution().
func (v *Surface) GetFallbackResolution() (xPPI, yPPI float64) {
	var x, y C.double
	C.cairo_surface_get_fallback_resolution(v.Native(), &x, &y)
	return float64(x), float64(y)
}

// GetType is a wrapper around cairo_surface_get_type().
func (v *Surface) GetType() SurfaceType {
	c := C.cairo_surface_get_type(v.Native())
	return SurfaceType(c)
}

// TODO(jrick) SetUserData (depends on UserDataKey and DestroyFunc)

// TODO(jrick) GetUserData (depends on UserDataKey)

// CopyPage is a wrapper around cairo_surface_copy_page().
func (v *Surface) CopyPage() {
	C.cairo_surface_copy_page(v.Native())
}

// ShowPage is a wrapper around cairo_surface_show_page().
func (v *Surface) ShowPage() {
	C.cairo_surface_show_page(v.Native())
}

// HasShowTextGlyphs is a wrapper around cairo_surface_has_show_text_glyphs().
func (v *Surface) HasShowTextGlyphs() bool {
	c := C.cairo_surface_has_show_text_glyphs(v.Native())
	return gobool(c)
}

// TODO(jrick) SetMimeData (depends on DestroyFunc)

// GetMimeData is a wrapper around cairo_surface_get_mime_data().  The
// returned mimetype data is returned as a Go byte slice.
func (v *Surface) GetMimeData(mimeType MimeType) []byte {
	cstr := C.CString(string(mimeType))
	defer C.free(unsafe.Pointer(cstr))
	var data *C.uchar
	var length C.ulong
	C.cairo_surface_get_mime_data(v.Native(), cstr, &data, &length)
	return C.GoBytes(unsafe.Pointer(data), C.int(length))
}

// TODO(jrick) SupportsMimeType (since 1.12)

// TODO(jrick) MapToImage (since 1.12)

// TODO(jrick) UnmapImage (since 1.12)
