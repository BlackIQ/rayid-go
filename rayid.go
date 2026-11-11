package rayid

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Type string

const (
	Digit Type = "digit"
	Lower Type = "lower"
	Upper Type = "upper"
	All   Type = "all"
)

var values = map[Type]string{
	Digit: "0123456789",
	Lower: "abcdefghijklmnopqrstuvwxyz",
	Upper: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

func init() {
	values[All] = values[Digit] + values[Lower] + values[Upper]
}

type RayID struct {
	typ Type
}

func New(t Type) (*RayID, error) {
	if _, ok := values[t]; !ok {
		return nil, fmt.Errorf("invalid type %q. expected one of: digit, lower, upper, all", t)
	}

	return &RayID{typ: t}, nil
}

func MustNew(t Type) *RayID {
	r, err := New(t)

	if err != nil {
		panic(err)
	}

	return r
}

func (r *RayID) Characters() string {
	return values[r.typ]
}

func (r *RayID) Gen(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than zero")
	}

	chars := r.Characters()
	var sb strings.Builder

	sb.Grow(length)

	if r.typ == Digit {
		firstChars := strings.ReplaceAll(chars, "0", "")
		c, err := randomChar(firstChars)

		if err != nil {
			return "", err
		}

		sb.WriteByte(c)

		for i := 1; i < length; i++ {
			c, err := randomChar(chars)

			if err != nil {
				return "", err
			}

			sb.WriteByte(c)
		}
	} else {
		for i := 0; i < length; i++ {
			c, err := randomChar(chars)

			if err != nil {
				return "", err
			}

			sb.WriteByte(c)
		}
	}

	return sb.String(), nil
}

func (r *RayID) MustGen(length int) string {
	id, err := r.Gen(length)

	if err != nil {
		panic(err)
	}

	return id
}

func randomChar(charset string) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

	if err != nil {
		return 0, err
	}

	return charset[n.Int64()], nil
}
