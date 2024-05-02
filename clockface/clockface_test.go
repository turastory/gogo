package clockface

import (
	"math"
	"testing"
	"time"
)

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2024, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func TestSecondsToRadian(t *testing.T) {
	cases := []struct {
		time time.Time
		want float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 15), 0.5 * math.Pi},
		{simpleTime(0, 0, 30), 1.0 * math.Pi},
		{simpleTime(0, 0, 45), 1.5 * math.Pi},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsToRadian(c.time)

			if !roughlyEqualFloat64(got, c.want) {
				t.Fatalf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestMinutesToRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * math.Pi / (30 * 60)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesToRadian(c.time)
			if got != c.angle {
				t.Fatalf("got %v, want %v", got, c.angle)
			}
		})
	}
}

func TestHoursToRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(3, 0, 0), math.Pi / 2},
		{simpleTime(18, 0, 0), math.Pi},
		{simpleTime(0, 7, 0), 7 * math.Pi / 12 / 30},
		{simpleTime(0, 0, 5), 5 * math.Pi / 12 / 30 / 60},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursToRadian(c.time)
			if got != c.angle {
				t.Fatalf("got %v, want %v", got, c.angle)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, -1}},
		{simpleTime(0, 0, 30), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, 1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, -1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(18, 0, 0), Point{0, 1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
