package url

import (
	"errors"
	"strings"
)

type Url struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(url string) (*Url, error) {
	idx := strings.Index(url, "://")
	if idx == -1 {
		return nil, errors.New("invalid url: missing scheme")
	}
	scheme, rest := url[:idx], url[idx+3:]
	host, path := rest, ""
	if idx = strings.Index(rest, "/"); idx != -1 {
		host, path = rest[:idx], rest[idx:]
	}
	return &Url{Scheme: scheme, Host: host, Path: path}, nil
}
