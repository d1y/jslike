package jsstring

import (
	"testing"
)

func TestStrings(t *testing.T) {
	var str JSString = "Hello World"
	t.Run("test At function", func(t *testing.T) {
		var last = str.At(-1)
		if last.NotEqual("d") {
			t.FailNow()
		}
	})

	t.Run("test Includes function", func(t *testing.T) {
		if !str.Includes("or", 7) {
			t.FailNow()
		}
	})

	t.Run("test padend/start", func(t *testing.T) {
		var result = str.PadEnd(42, "*|&")
		if result.Length() != 42 {
			t.FailNow()
		}
		var result1 = str.PadStart(18, "*")
		if result1.At(-1) != "d" || result1.Length() != 18 {
			t.FailNow()
		}
	})

	t.Run("test slice demo", func(t *testing.T) {
		var result = str.Slice(4, 6)
		if len(result) != 2 || result.NotEqual("o ") {
			t.FailNow()
		}
	})

	t.Run("test substring demo", func(t *testing.T) {
		var result = str.Substring(4, 6)
		if result.Length() != 2 || result.NotEqual("o ") {
			t.FailNow()
		}
	})

	t.Run("test trim", func(t *testing.T) {
		var str JSString = "  s  dfds v  "
		str = str.Trim()
		if str.At(-1).Equal(" ") {
			t.FailNow()
		}
	})
}
