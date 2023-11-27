package nilsafer

import (
	"testing"
)

func TestValueOrZero(t *testing.T) {
	t.Run("safe string success with value", func(t *testing.T) {
		var s = "hello"
		result := ValueOrZero(&s)
		want := "hello"
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
	t.Run("safe string success with zero value", func(t *testing.T) {
		var s *string
		result := ValueOrZero(s)
		want := ""
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
	t.Run("safe boolean success with value", func(t *testing.T) {
		var b = true
		result := ValueOrZero(&b)
		want := true
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
	t.Run("safe boolean success with zero value", func(t *testing.T) {
		var b *bool
		result := ValueOrZero(b)
		want := false
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
	t.Run("safe number success with value", func(t *testing.T) {
		var n = 1
		result := ValueOrZero(&n)
		want := 1
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
	t.Run("safe number success with zero value", func(t *testing.T) {
		var n *int
		result := ValueOrZero(n)
		want := 0
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
	t.Run("safe object success with value", func(t *testing.T) {
		var o = struct {
			Name string
		}{
			Name: "hello",
		}
		result := ValueOrZero(&o)
		want := struct {
			Name string
		}{
			Name: "hello",
		}
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
	t.Run("safe object success with zero value", func(t *testing.T) {
		var o *struct {
			Name string
		}
		result := ValueOrZero(o)
		want := struct {
			Name string
		}{
			Name: "",
		}
		if result != want {
			t.Errorf("Unexpected result %v, want %v", result, want)
		}
	})
}
