package gomarketcap

import (
	"math/rand"
	"testing"
	"time"
)

func TestPad(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{1, "01"},
		{11, "11"},
		{15, "15"},
		{10, "10"},
	}
	for _, test := range tests {
		if got := Pad(test.input); got != test.want {
			t.Errorf("pad(%v) = %s", test.input, got)
		}
	}
}

func randomDate(rng *rand.Rand) (date time.Time) {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rng.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func TestFormatDate(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for index := 0; index < 1000; index++ {
		d := randomDate(rng)
		l := len(FormatDate(d))
		// Can't really test it, we'll just check if the length is right
		if l != 8 {
			t.Errorf("formatDate(%s) = %v", d, l)
		}
	}
}
