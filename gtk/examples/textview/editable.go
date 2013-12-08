package main
import (
"github.com/weberc2/gotk3/gtk"
)
func main() {
	gtk.Init(nil)

	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.SetDefaultSize(500, 200)
	w.SetTitle("TextView Get/SetEditable example")
	hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	w.Add(hbox)

	tv1, _ := gtk.TextViewNew()
	tv1.SetEditable(false)
	buf, _ := tv1.GetBuffer()
	buf.SetText("Try and edit me. Betcha Can't")
	hbox.PackStart(tv1, true, true, 0)

	tv2, _ := gtk.TextViewNew()
	tv2.SetEditable(true)
	buf, _ = tv2.GetBuffer()
	buf.SetText("You can edit me though.")
	hbox.PackStart(tv2, true, true, 0)

	w.ShowAll()

	gtk.Main()
}
