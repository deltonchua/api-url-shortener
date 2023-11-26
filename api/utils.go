package api

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/url"
)

func isURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	u, parseErr := url.Parse(input)
	if parseErr != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func nanoid() (string, error) {
	const (
		alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		length   = 8
	)
	return gonanoid.Generate(alphabet, length)
}
