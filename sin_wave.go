// steve.mynott@gmail.com 20150823
package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
	maths "math"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	canvas := driver.CreateCanvas(math.Size{W: 800, H: 600})

	for i := 0; i < 800; i++ {
		var fy float64 = maths.Sin(float64(i) * (maths.Pi / 180))
		y := (int(fy*100) + 240)
		plot(canvas, i, y)
	}

	canvas.Complete()

	image := theme.CreateImage()
	image.SetCanvas(canvas)

	window := theme.CreateWindow(800, 600, "")
	window.AddChild(image)
	window.OnClose(driver.Terminate)

}

func plot(c gxui.Canvas, x int, y int) {
	brush := gxui.CreateBrush(gxui.White)
	c.DrawRect(math.CreateRect(x, y, x+1, y+1), brush)
}

func main() {
	gl.StartDriver(appMain)
}
