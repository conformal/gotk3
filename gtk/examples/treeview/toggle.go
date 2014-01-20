package main

import (
	"log"
	"github.com/conformal/gotk3/gtk"
	"github.com/conformal/gotk3/glib"
)

const (
	TOGGLE_COLUMN = 0
)

func initWindow() *gtk.Window {
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(200, 100)
	win.SetTitle("TreeView Toggle Example")
	win.Connect("destroy", gtk.MainQuit)
	return win
}

func initTreeView(ls *gtk.ListStore) *gtk.TreeView {
	tv, _ := gtk.TreeViewNewWithModel(ls)

	toggleRenderer, _ := gtk.CellRendererToggleNew()
	doneCol, _ := gtk.TreeViewColumnNewWithAttribute("Toggle", toggleRenderer, "active", TOGGLE_COLUMN)
	tv.AppendColumn(doneCol)
	return tv
}

func addRow(ls *gtk.ListStore, val bool) {
	var iter gtk.TreeIter
	ls.Append(&iter)
	if err := ls.Set(&iter, []int{TOGGLE_COLUMN}, []interface{}{val}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	gtk.Init(nil)
	win := initWindow()
	ls, _ := gtk.ListStoreNew(glib.TYPE_BOOLEAN)
	tv := initTreeView(ls)
	win.Add(tv)

	addRow(ls, true)
	addRow(ls, false)
	addRow(ls, true)

	win.ShowAll()
	gtk.Main()
}
