package m8

import (
	"testing"
)

func TestResponse(t *testing.T) {
	t.Parallel()

	got := Response(1)
	want := "It is decidedly so."

	if got != want {
		t.Errorf("got:%s - want:%s", got, want)
	}
}

func TestResponse2(t *testing.T) {
	t.Parallel()

	got := Response(2)
	want := "Most likely."

	if got != want {
		t.Errorf("got:%s - want:%s", got, want)
	}
}

func TestResponse3(t *testing.T) {
	t.Parallel()

	got := Response(3)
	want := "Yes."

	if got != want {
		t.Errorf("got:%s - want:%s", got, want)
	}
}
