package inner

import "testing"

func Test_Valid(t *testing.T) {
	cases := []struct {
		input int
		want  bool
	}{
		{
			input: 0,
			want:  true,
		},
		{
			input: 67,
			want:  false,
		},
		{
			input: 1,
			want:  true,
		},
		{
			input: 2,
			want:  true,
		},
		{
			input: 3,
			want:  true,
		},
		{
			input: 25,
			want:  true,
		},
		{
			input: 26,
			want:  true,
		},
	}

	for _, c := range cases {
		got := Valid(c.input)
		if got != c.want {
			t.Fatalf("Got: %v, Want: %v", got, c.want)
		}
	}
}