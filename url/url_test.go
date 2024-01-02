package url

import "testing"

func TestParse(t *testing.T) {
	want := "http"                           // arrange
	u, err := Parse("http://www.google.com/go") // act
	if err != nil {
		t.Fatal(err)
	}
	if got := u.Scheme; got != want { // assert
		t.Fatalf("expected scheme http, got %s", got)
	}
	if got, want := u.Host, "www.google.com"; got != want {
		t.Fatalf("expected host %s, got %s", want, got)
	}
	if got, want := u.Path, "/go"; got != want {
		t.Fatalf("expected path %s, got %s", want, got)
	}
}
