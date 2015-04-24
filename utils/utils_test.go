package utils

import (
	"errors"
	"testing"
)

func TestToParamString(t *testing.T) {
	cases := [][2]string{
		{"OrderItem", "order_item"},
		{"order item", "order_item"},
		{"Order Item", "order_item"},
	}
	for _, c := range cases {
		if got := ToParamString(c[0]); got != c[1] {
			t.Errorf("ToParamString(%q) = %q; want %q", c[0], got, c[1])
		}
	}
}

func TestPatchUrl(t *testing.T) {
	var cases = []struct {
		original string
		input    []interface{}
		want     string
		err      error
	}{
		{
			original: "http://qor.com/admin/orders?locale=global&q=dotnet&test=1#test",
			input:    []interface{}{"locale", "cn"},
			want:     "http://qor.com/admin/orders?locale=cn&q=dotnet&test=1#test",
		},
		{
			original: "http://qor.com/admin/orders?locale=global&q=dotnet&test=1#test",
			input:    []interface{}{"locale", ""},
			want:     "http://qor.com/admin/orders?q=dotnet&test=1#test",
		},
		{
			original: "http://qor.com/admin/orders?locale=global&q=dotnet&test=1#test",
			input:    []interface{}{"locale", 1},
			err:      errors.New("1 type is int, want string"),
		},
	}
	for _, c := range cases {
		// u, _ := url.Parse(c.original)
		// context := Context{Context: &qor.Context{Request: &http.Request{URL: u}}}
		got, err := PatchURL(c.original, c.input...)
		if c.err != nil {
			if err == nil || err.Error() != c.err.Error() {
				t.Errorf("got error %s; want %s", err, c.err)
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if got != c.want {
				t.Errorf("context.PatchURL = %s; c.want %s", got, c.want)
			}
		}
	}
}
