/* 一个ui仪表盘
- Premode widget 预定义widges
- Custom widget 自定义widges
- Grid or absolute coordinates
- Keyboard, mouse, and terminal resizing events 事件
- Colors and styling
*/

package main

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

/* demo 任务
1, Premode展示
2，布局
3，重绘（数据更新 & 配置）
4，事件（Resize）
5，自定义widget）-》 echart & 饿了么控件
6，配置
*/
func main() {
	// Init
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// 1, Premode展示
	/* 12栅格系统：*/
	grid0 := ui.NewGrid()
	grid0.Border = true
	// gridPremode.BorderStyle = Theme
	termWidth, termHeight := ui.TerminalDimensions()
	grid0.SetRect(0, 0, termWidth, termHeight)

	// -------------barchart------------
	bc := widgets.NewBarChart()
	bc.Border = false
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	// bc.BarWidth = 3
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen, ui.ColorMagenta}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue), ui.NewStyle(ui.ColorRed)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
	bc.BarGap = 1
	// bc.NumFormatter = func(n float64) string { return fmt.Sprintf("num:%g", n) }

	// bc.SetRect(0, 0, termWidth, termHeight)
	// ui.Render(bc)

	// -------------sparkline：体积小和数据密度高的图表-看不懂-----------
	// sl0 := widgets.NewSparkline()
	// sl0.Title = "Sparkline"
	// sl0.LineColor = ui.ColorGreen // 数据块颜色
	// // sl0.Data = []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	// sl0.Data = []float64{1, 1}
	// slg0 := widgets.NewSparklineGroup(sl0)
	// slg0.Title = "Sparkline Group 0"
	// // slg0.Border = false
	// slg0.BorderStyle.Fg = ui.ColorCyan

	// -------------Gauge 计量-----------
	g0 := widgets.NewGauge()
	g0.Title = "Gauge 0"
	g0.Percent = 75
	// g0.Block.Inner.Max.Y = 2
	g0.Border = false

	g1 := widgets.NewGauge()
	g1.Title = "Gauge 1"
	g1.Percent = 30
	g1.BarColor = ui.ColorRed
	g1.Border = false
	g1.Label = fmt.Sprintf("%v%%(100MBs free)", g1.Percent)

	// // -------------List------------------
	l := widgets.NewList()
	l.Title = "List"
	l.Rows = []string{
		"[0] github.com/gizak/termui/v3",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] foo",
		"[8] bar",
		"[9] baz",
		"[10] 新中国成立70周年.go",
		"[11] Go语言",
		"[12] 蒋毅勋",
		"[13] Baker",
	}
	// l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.Border = false
	l.WrapText = false
	l.SelectedRow = 10
	l.SelectedRowStyle = ui.NewStyle(ui.ColorCyan)

	grid0.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, bc),
			ui.NewCol(1.0/2,
				ui.NewRow(1.0/3, g0),
				ui.NewRow(1.0/3, g1),
			),
		),
		ui.NewRow(1.0/2, l),
	)
	// grid0.Set(
	// 	ui.NewRow(1.0/2, g1),
	// )

	ui.Render(grid0)

	/* custom widgets 参照v3/widgets
	 */

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<Resize>":
				paylaod := e.Payload.(ui.Resize)
				// bc.SetRect(0, 0, paylaod.Width, paylaod.Height)
				grid0.SetRect(0, 0, paylaod.Width, paylaod.Height)
			case "q", "<C-c>":
				return
			case "j":
				l.ScrollDown()
			case "k":
				l.ScrollUp()
			}
		case <-ticker:
			ui.Render(grid0)
			// ui.Render(bc)
		}

	}
}
