package rayid

import (
	"testing"
	"unicode"
)

func TestDigit(t *testing.T) {
	r := MustNew(Digit)
	id := r.MustGen(20)

	if len(id) != 20 {
		t.Fatalf("expected length 20, got %d", len(id))
	}

	if id[0] == '0' {
		t.Error("digit id should not start with zero")
	}

	for _, c := range id {
		if !unicode.IsDigit(c) {
			t.Errorf("unexpected character: %c", c)
		}
	}
}

func TestLower(t *testing.T) {
	r := MustNew(Lower)
	id := r.MustGen(15)

	for _, c := range id {
		if c < 'a' || c > 'z' {
			t.Errorf("unexpected character: %c", c)
		}
	}
}

func TestInvalidType(t *testing.T) {
	_, err := New("invalid")
	if err == nil {
		t.Error("expected error for invalid type")
	}
}
