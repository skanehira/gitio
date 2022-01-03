package cmd

import "testing"

func TestGetShortenURL(t *testing.T) {
	input := "https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests"
	got, err := getShortenURL(input)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	want := "https://git.io/JSKd3"

	if got != want {
		t.Fatalf("unexpected url. want: %s, got: %s", want, got)
	}
}

func TestGetShortenURLFail(t *testing.T) {
	input := "https://google.com"
	_, err := getShortenURL(input)
	want := "Must be a GitHub.com URL."
	if err.Error() != want {
		t.Fatalf("unexpected error. want: %s, got: %s", want, err)
	}
}
