package main

import (
	"errors"
	"os"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("Error standalone", func(t *testing.T) {
		var e1 error = OopsNonwrapped()

		assertError(t, e1, ErrConst)
		assertNotError(t, e1, nil)
	})

	t.Run("Error wraps os.ErrInvalid", func(t *testing.T) {
		var e1 error = OopsInvalid()

		assertError(t, e1, ErrConst)
		assertError(t, e1, os.ErrInvalid)
		assertUnwrap(t, e1, os.ErrInvalid)
	})

	t.Run("Error wraps os.ErrExist", func(t *testing.T) {
		var e1 error = OopsExists()

		assertError(t, e1, ErrConst)
		assertError(t, e1, os.ErrExist)
		assertUnwrap(t, e1, os.ErrExist)
	})

	t.Run("Error as os.ErrInvalid", func(t *testing.T) {
		e1 := OopsInvalid()

		assertError(t, e1, ErrConst)
		assertError(t, e1, os.ErrInvalid)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("wrong error: expected %q, got %q", want, got)
	}
}

func assertNotError(t *testing.T, got, want error) {
	t.Helper()

	if errors.Is(got, want) {
		t.Errorf("wrong error: expected %q, got %q", want, got)
	}
}

func assertUnwrap(t *testing.T, e, want error) {
	t.Helper()

	got := errors.Unwrap(e)

	if got != want {
		t.Errorf("unwrapped wrong error: got %q, want %q", got, want)
	}
}
