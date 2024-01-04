package url

import (
	"testing"
)

func TestParse(t *testing.T) {
	want := "http"                              // arrange
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

func TestPort(t *testing.T) {
	data := map[string]struct {
		host string
		want string
	}{
		"no port":   {"www.google.com", ""},
		"with port": {"www.google.com:80", "80"},
		"localhost": {"localhost:8080", "8080"},
		"443":       {"www.google.com:443", "443"},
	}
	for name, test := range data {
		t.Run(name, func(t *testing.T) {
			u := &Url{Host: test.host}
			if got := u.Port(); got != test.want {
				t.Errorf("expected  %s, got %s", test.want, got)
			}
		})
	}
}

func TestHostname(t *testing.T) {
	data := map[string]struct {
		host string
		want string
	}{
		"no port":   {"www.google.com", "www.google.com"},
		"with port": {"www.google.com:80", "www.google.com"},
		"localhost": {"localhost:8080", "localhost"},
	}
	for name, test := range data {
		t.Run(name, func(t *testing.T) {
			u := &Url{Host: test.host}
			if got := u.Hostname(); got != test.want {
				t.Fatalf("expected  %s, got %s", test.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	data := map[string]struct {
		url  *Url
		want string
	}{
		"no scheme":   {&Url{Host: "www.google.com"}, "www.google.com"},
		"with scheme": {&Url{Scheme: "http", Host: "www.google.com"}, "http://www.google.com"},
		"with path":   {&Url{Scheme: "http", Host: "www.google.com", Path: "/go"}, "http://www.google.com/go"},
	}
	for name, test := range data {
		t.Run(name, func(t *testing.T) {
			if got := test.url.String(); got != test.want {
				t.Fatalf("expected  %s, got %s", test.want, got)
			}
		})
	}
}

func BenchmarkString(b *testing.B) {
	b.ReportAllocs()
	u := &Url{Scheme: "http", Host: "www.google.com", Path: "/go"}
	for i := 0; i < b.N; i++ {
		u.String()
	}
}

func BenchmarkSlowString(b *testing.B) {
	b.ReportAllocs()
	u := &Url{Scheme: "http", Host: "www.google.com", Path: "/go"}
	for i := 0; i < b.N; i++ {
		u.SlowString()
	}
}
