package main

import (
	"bytes"
	"math/rand"
	"time"

	mg "github.com/erkkah/margaid"
	"github.com/janpfeifer/gonb/gonbui"
)

func mgPlot(width, height int) string {
	// Используем mg.Series{} вместо mg.NewSeries()
	randomSeries := mg.NewSeries()
	rand.Seed(time.Now().Unix())
	for i := float64(0); i < 10; i++ {
		randomSeries.Add(mg.MakeValue(i+1, 200*rand.Float64()))
	}

	testSeries := mg.NewSeries() //на mg.Series{}
	multiplier := 2.1
	v := 0.33
	for i := float64(0); i < 10; i++ {
		v *= multiplier
		testSeries.Add(mg.MakeValue(i+1, v))
	}

	diagram := mg.New(width, height,
		mg.WithAutorange(mg.XAxis, testSeries),
		mg.WithAutorange(mg.YAxis, testSeries),
		mg.WithAutorange(mg.Y2Axis, testSeries),
		mg.WithProjection(mg.YAxis, mg.Log),
		mg.WithInset(70),
		mg.WithPadding(2),
		mg.WithColorScheme(90),
		mg.WithBackgroundColor("#f8f8f8"),
	)

	diagram.Line(testSeries, mg.UsingAxes(mg.XAxis, mg.YAxis), mg.UsingMarker("square"), mg.UsingStrokeWidth(1))
	diagram.Smooth(testSeries, mg.UsingAxes(mg.XAxis, mg.Y2Axis), mg.UsingStrokeWidth(3.14))
	diagram.Smooth(randomSeries, mg.UsingAxes(mg.XAxis, mg.YAxis), mg.UsingMarker("filled-circle"))
	diagram.Axis(testSeries, mg.XAxis, diagram.ValueTicker('f', 0, 10), false, "X")
	diagram.Axis(testSeries, mg.YAxis, diagram.ValueTicker('f', 1, 2), true, "Y")

	diagram.Frame()
	diagram.Title("A diagram of sorts 📊 📈")
	buf := bytes.NewBuffer(nil)
	diagram.Render(buf)
	return buf.String()
}

func main() {
	gonbui.DisplaySVG(mgPlot(640, 480))
}
