package app

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

const (
	viewPart1 = "parts1"
)

// viewProperties 视图属性
type viewProperties struct {
	title    string
	text     string
	x1       float64 //x1，y1，x2，y2是百分比
	y1       float64
	x2       float64
	y2       float64
	editor   gocui.Editor
	editable bool
	modal    bool
}

// views 所有视图
var views = []string{
	viewPart1,
}

// selectableViews 可被选择的视图
var selectableViews = []string{
	viewPart1,
}

// vp 视图属性map
var vp = map[string]viewProperties{
	viewPart1: {
		title:    "Parts 1",
		text:     "",
		x1:       0.0,
		y1:       0.0,
		x2:       1,
		y2:       0.1,
		editor:   &le,
		editable: true,
		modal:    false,
	},
}

// Layout sets up the views
func (a *App) Layout(g *gocui.Gui) error {
	for _, v := range views {
		if err := a.initView(v); err != nil {
			return err
		}
	}

	// Set the first view on the first run
	if a.currentView == -1 {
		a.currentView = 0
		a.setSelectableView(a.currentView)
	}

	a.updateState()
	a.updateViews()

	return nil
}

// initView 初始化view
func (a *App) initView(viewName string) error {
	maxX, maxY := a.gui.Size()

	p := vp[viewName]

	if p.modal {
		// Don't init modals
		return nil
	}

	x1 := int(p.x1 * float64(maxX))
	y1 := int(p.y1 * float64(maxY))
	x2 := int(p.x2*float64(maxX)) - 1
	y2 := int(p.y2*float64(maxY)) - 1

	return a.createView(viewName, x1, y1, x2, y2)
}

// createview creates a new view
func (a *App) createView(viewName string, x1, y1, x2, y2 int) error {
	if v, err := a.gui.SetView(viewName, x1, y1, x2, y2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		p := vp[viewName]
		v.Title = p.title
		v.Editor = p.editor
		v.Editable = p.editable

		a.writeView(viewName, p.text)
	}

	return nil
}

// setSelectableView set the focus to the view specified by id
func (a *App) setSelectableView(id int) error {
	if err := a.setView(selectableViews[id]); err != nil {
		return err
	}
	a.currentView = id

	return nil
}

// setView set the focus to the view specified by name
func (a *App) setView(name string) error {
	if _, err := a.gui.SetCurrentView(name); err != nil {
		return err
	}

	return nil
}

// clearView clears a view
func (a *App) clearView(name string) {
	v, _ := a.gui.View(name)
	v.Clear()
}

// closeView closes a view
func (a *App) closeView(name string) {
	a.gui.DeleteView(name)
	a.setView(viewPart1)
}

// writeView writes string to view
func (a *App) writeView(name, text string) {
	v, _ := a.gui.View(name)
	v.Clear()
	fmt.Fprint(v, text)
	v.SetCursor(len(text), 0)
}

// showModal shows a modal dialog on top of other views
func (a *App) showModal(name, text string, width, height float64) {
	p := vp[name]
	p.text = text
	vp[name] = p

	maxX, maxY := a.gui.Size()

	modalWidth := int(float64(maxX) * width)
	modalHeight := int(float64(maxY) * height)

	x1 := (maxX - modalWidth) / 2
	x2 := x1 + modalWidth
	y1 := (maxY - modalHeight) / 2
	y2 := y1 + modalHeight

	a.createView(name, x1, y1, x2, y2)
	a.setView(name)
}
