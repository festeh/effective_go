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

func (u *Url) Port() string { 
	idx := strings.Index(u.Host, ":")
	if idx == -1 {
		return ""
	}
	return u.Host[idx+1:]
}

func (u *Url) Hostname() string {
	idx := strings.Index(u.Host, ":")
	if idx == -1 {
		return u.Host
	}
	return u.Host[:idx]
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

func (u *Url) String() string {
	var builder strings.Builder
	if u.Scheme != "" {
		builder.WriteString(u.Scheme)
		builder.WriteString("://")
	}
	if u.Host != "" {
		builder.WriteString(u.Host)
	}
	if u.Path != "" {
		builder.WriteString(u.Path)
	}
	return builder.String()
}

func (u *Url) SlowString() string {
	s := ""
	if u.Scheme != "" {
		s += u.Scheme + "://"
	}
	if u.Host != "" {
		s += u.Host
	}
	if u.Path != "" {
		s += u.Path
	}
	return s
}
