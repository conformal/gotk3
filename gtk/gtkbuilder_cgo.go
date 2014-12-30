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


typedef void (*callbackFunc) ();

#define BUILDER_DATA(ptr) ((struct gtkbuilder_data *)(ptr))

// TODO: Free data when GtkBuilder is freed.
struct gtkbuilder_data {
	char *signal;
	void *builder;
};


void gtkbuilder_handler(GClosure *closure,
                    GValue *return_value,
                    guint nparams,
                    const GValue *params,
                    gpointer invocation_hint,
                    gpointer marshal_data) {

	const GValue **p;
	int c;

	p = (const GValue **)malloc(nparams * sizeof(GValue*));
	for (c = 0; c < nparams; c++) {
		p[c] = params+c;
	}

	dispatchEvent(BUILDER_DATA(marshal_data)->builder, BUILDER_DATA(marshal_data)->signal, return_value, nparams, sizeof(GValue), params);
	free(p);
}

void gtkbuilder_connector(GtkBuilder *builder,
                          GObject *object,
                          const gchar *signal_name,
                          const gchar *handler_name,
                          GObject *connect_object,
                          GConnectFlags flags,
                          gpointer user_data)
{
	struct gtkbuilder_data *data;
	GClosure *c;

	if (!(data = (struct gtkbuilder_data *)malloc(sizeof(struct gtkbuilder_data)))) {
			fprintf(stderr, "ERROR: Not enough memory.\n");
			return;
	}

	data->signal = g_strdup_printf("%s", handler_name);
	data->builder = user_data;
	printf("Connect '%s' w/ %s\n", signal_name, data->signal);

	c = g_closure_new_simple(sizeof(GClosure), data);
	g_closure_set_meta_marshal(c, data, gtkbuilder_handler);
	g_signal_connect_closure(object, signal_name, c, FALSE);
}

*/
import "C"
import (
	"unsafe"
)

// Bind callbacks methods using reflection. Callbacks must be exported (first letter in upper case)
func (b *Builder) BindCallbacks(ifaces ...interface{}) {
	for _, v := range ifaces {
		b.callbacks = append(b.callbacks, v)
	}

	C.gtk_builder_connect_signals_full(b.native(), C.callbackFunc(C.gtkbuilder_connector), C.gpointer(unsafe.Pointer(b)))

}
