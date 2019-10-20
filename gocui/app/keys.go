package app

import (
	"github.com/jroimartin/gocui"
)

type keyConfig struct {
	views   *[]string                           //绑定的视图
	key     interface{}                         //绑定的键
	mod     gocui.Modifier                      //alt
	handler func(*gocui.Gui, *gocui.View) error //处理函数
}

// setKeyBindings sets up the keyboard shortcuts
func (a *App) setKeyBindings() error {
	var kc = []keyConfig{
		{
			&selectableViews,
			gocui.KeyCtrlQ,
			gocui.ModNone,
			a.quit,
		},
	}

	for _, shortcut := range kc {
		for _, view := range *shortcut.views {
			if err := a.gui.SetKeybinding(
				view,
				shortcut.key,
				shortcut.mod,
				shortcut.handler,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

// quit handles quit keyboard shortcut
func (a *App) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
