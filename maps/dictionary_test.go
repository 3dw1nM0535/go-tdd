package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	assertString := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q, expected %q", got, want)
		}
	}

	t.Run("Search for a known key in the dictionary", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is a test"

		assertString(t, got, want)
	})

	t.Run("Search for an unknown key in the dictionary", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrorKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	dict.Add("test", "this is an add test")

	want := "this is an add test"
	got, err := dict.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}
