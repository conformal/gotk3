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
package main

import (
	"fmt"
	"github.com/conformal/gotk3/gtk"
	"github.com/conformal/gotk3/gtk/gtkx"
	"log"
	"os"
	"os/exec"
)

func main() {
	/*
	   We set an environment variable so that we don't get bothered by xterm accessibility warnings
	   like "** (xterm:16917): WARNING **: Couldn't connect to accessibility bus: Failed to connect to socket /tmp/dbus-QznzMfGEXB: Connection refused"
	   See http://debbugs.gnu.org/cgi/bugreport.cgi?bug=15154
	*/
	os.Setenv("NO_AT_BRIDGE", "1")

	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new socket:
	s, err := gtkx.SocketNew()
	if err != nil {
		log.Fatal("Unable to create socket:", err)
	}

	//Adding the socket to the window:
	win.Add(s)

	// Getting the socketId:
	sId := s.GetId()
	fmt.Printf("Our socket: %v\n", sId)

	// Embedding something in the socket:
	xterm := exec.Command("xterm", "-into", gtkx.WindowIdToString(sId))

	go func() {
		xterm.Run()
	}()

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
