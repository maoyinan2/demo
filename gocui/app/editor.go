package app

import "github.com/jroimartin/gocui"

type lineEditor struct {
	gocuiEditor gocui.Editor
}

var le lineEditor

// Edit 安装handler
func (e *lineEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch key {
	// 禁用换行
	case gocui.KeyEnter:
		return

	// 禁用右箭头换行
	case gocui.KeyArrowRight:
		x, _ := v.Cursor()
		if x >= len(v.ViewBuffer())-1 {
			return
		}

	case gocui.KeyHome:
		v.SetCursor(0, 0)
		return

	case gocui.KeyEnd:
		v.SetCursor(len(v.ViewBuffer())-1, 0)
		return
	}

	e.gocuiEditor.Edit(v, key, ch, mod)
}
