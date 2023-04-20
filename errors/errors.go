package error

import "github.com/pkg/errors"

var (
	ErrInvalidBody = errors.New("invalid body is passed")
	ErrUrl         = errors.New("invalid url is passed")
	ErrNotFound    = errors.New("url not found")
	ErrCollision   = errors.New("collision detected")
	ErrInsecureUrl = errors.New("insecure url")
	ErrEmptyUrl    = errors.New("empty url is passed")
)
