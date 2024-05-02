package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength   = 90
	minuteHandLength   = 80
	hourHandLength     = 50
	secondsInHalfClock = 30
	secondsInClock     = secondsInHalfClock * 2
	minutesInHalfClock = 30
	minutesInClock     = minutesInHalfClock * 2
	hoursInHalfClock   = 6
	hoursInClock       = hoursInHalfClock * 2
)

func (a *Point) Add(b *Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func (a *Point) Scale(factor float64) Point {
	return Point{a.X * factor, a.Y * factor}
}

func secondsToRadian(time time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(time.Second()))
}

func minutesToRadian(tm time.Time) (radian float64) {
	radian = math.Pi / (minutesInHalfClock / float64(tm.Minute()))
	radian += secondsToRadian(tm) / minutesInClock
	return
}

func hoursToRadian(tm time.Time) (radian float64) {
	radian = math.Pi / (hoursInHalfClock / float64(tm.Hour()%hoursInClock))
	radian += minutesToRadian(tm) / hoursInClock
	return
}

func secondHandPoint(tm time.Time) Point {
	return angleToPoint(secondsToRadian(tm))
}

func minuteHandPoint(tm time.Time) Point {
	return angleToPoint(minutesToRadian(tm))
}

func hourHandPoint(tm time.Time) Point {
	return angleToPoint(hoursToRadian(tm))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), -math.Cos(angle)}
}

var base = Point{150, 150}

func secondHand(tm time.Time) Point {
	return makeHand(secondHandPoint(tm), secondHandLength)
}

func minuteHand(tm time.Time) Point {
	return makeHand(minuteHandPoint(tm), minuteHandLength)
}

func hourHand(tm time.Time) Point {
	return makeHand(hourHandPoint(tm), hourHandLength)
}

func makeHand(direction Point, scale float64) Point {
	p := direction.Scale(scale)
	return base.Add(&p)
}

func SVGWriter(w io.Writer, time time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)

	p := secondHand(time)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="stroke:#000;stroke-width:2px;"/>`, p.X, p.Y)

	p = minuteHand(time)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="stroke:#000;stroke-width:2px;"/>`, p.X, p.Y)

	p = hourHand(time)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="stroke:#000;stroke-width:2px;"/>`, p.X, p.Y)

	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
