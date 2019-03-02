package main

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	for _, tc := range []struct {
		t    entityType
		in   string
		want string
	}{
		{
			t:    typeHTML,
			in:   `<text>foo&bar</text>`,
			want: `&lt;text&gt;foo&amp;bar&lt;/text&gt;`,
		},
		{
			t:    typeURL,
			in:   `абвгд / abcde`,
			want: `%D0%B0%D0%B1%D0%B2%D0%B3%D0%B4+%2F+abcde`,
		},
		{
			t:    typeBase64,
			in:   `{"foo": "bar"}`,
			want: `eyJmb28iOiAiYmFyIn0=`,
		},
	} {
		got, err := encode(tc.in, tc.t)
		if err != nil {
			t.Error(err)
		}

		if got != tc.want {
			t.Errorf("want %q, got %q", tc.want, got)
		}

		got, err = decode(got, tc.t)
		if err != nil {
			t.Error(err)
		}

		if got != tc.in {
			t.Errorf("want %q, got %q", tc.in, got)
		}
	}
}
