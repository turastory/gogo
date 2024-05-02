package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/turastory/gogo/clockface"
)

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2024, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func TestSVGWriterForSecondHands(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			var b bytes.Buffer
			clockface.SVGWriter(&b, c.time)

			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf(`Expected to find a line %+v, in the SVG. Got %+v`, c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterForMinuteHands(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
	}

	for _, c := range cases {
		var b bytes.Buffer
		clockface.SVGWriter(&b, c.time)

		svg := Svg{}
		xml.Unmarshal(b.Bytes(), &svg)

		if !containsLine(c.line, svg.Line) {
			t.Errorf(`Expected to find a line %+v, in the SVG. Got %+v`, c.line, svg.Line)
		}
	}
}

func TestSVGWriterForHourHands(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 100}},
		{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
	}

	for _, c := range cases {
		var b bytes.Buffer
		clockface.SVGWriter(&b, c.time)

		svg := Svg{}
		xml.Unmarshal(b.Bytes(), &svg)

		if !containsLine(c.line, svg.Line) {
			t.Errorf(`Expected to find a line %+v, in the SVG. Got %+v`, c.line, svg.Line)
		}
	}

}

func containsLine(want Line, gots []Line) bool {
	for _, line := range gots {
		if line == want {
			return true
		}
	}
	return false
}
