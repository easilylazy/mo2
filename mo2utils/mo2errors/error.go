package mo2errors

import "fmt"

type Mo2Errors struct {
	ErrorCode int
	ErrorTip  string
}

func (e Mo2Errors) Error() string {
	return fmt.Sprintf("%v: %v", e.ErrorCode, e.ErrorTip)
}
func (e *Mo2Errors) SetErrorTip(s string) {
	e.ErrorTip = s
}
func (e *Mo2Errors) Init(c int, s string) {
	e.ErrorCode = c
	e.ErrorTip = s
}
func (e *Mo2Errors) InitCode(c int) {
	e.ErrorCode = c
	e.ErrorTip = CodeText(c)
}
func (e Mo2Errors) IsError() (error bool) {
	error = true
	if e.ErrorCode == Mo2NoError {
		error = false
	}
	return
}

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(c int, s string) error {
	return &Mo2Errors{c, s}
}

// NewCode returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func NewCode(c int) error {
	return &Mo2Errors{c, CodeText(c)}
}
