package pagination

import "errors"

var (
	ErrorMaxPage     = errors.New("Chosen page more then total page")
	ErrorPage        = errors.New("Page must be greater than zero")
	ErrorPageEmpty   = errors.New("Page cannot be empty")
	ErrorPageInvalid = errors.New("Page invalid, must be number")
)
