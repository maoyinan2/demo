package app

import (
	"log"

	"github.com/jroimartin/gocui"
)

// App 应用程序
type App struct {
	gui         *gocui.Gui
	currentView int
	state       *state
}

// state 界面输入
type state struct {
	// Settings map[string]bool
}

// New 创建app实例
func New() *App {
	a := new(App)

	var err error
	a.gui, err = gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}

	// 默认设置
	a.state = new(state)
	// a.state.Settings = map[string]bool{
	// 	"TLDSubstitutions": false,
	// }

	a.initGui()

	return a
}

// Loop starts the GUI loop
func (a *App) Loop() {
	if err := a.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// initGui 初始化GUI
func (a *App) initGui() {
	le = lineEditor{gocui.DefaultEditor}

	// Defaults
	a.gui.Cursor = true
	a.gui.InputEsc = false // esc不作为输入字符
	a.gui.BgColor = gocui.ColorDefault
	a.gui.FgColor = gocui.ColorDefault

	// Set Layout function
	a.gui.SetManagerFunc(a.Layout)

	a.currentView = -1

	// Set Keybindings
	a.setKeyBindings()
}

// Close closes the app
func (a *App) Close() {
	a.gui.Close()
}

// updateState saves the current state of views
func (a *App) updateState() {
	// a.state.Parts1 = a.parseLine(viewPart1)
	// a.state.Parts2 = a.parseLine(viewPart2)
	// a.state.Tlds = a.parseLine(viewTLD)
	// a.state.Domains = a.parseLine(viewDomain)
}

// updateViews updates the views based on the current state
func (a *App) updateViews() {
	// if a.state.Settings["TLDSubstitutions"] {
	// 	a.writeView(viewSettings, "[X] TLD substitutions")
	// } else {
	// 	a.writeView(viewSettings, "[ ] TLD substitutions")
	// }
}
