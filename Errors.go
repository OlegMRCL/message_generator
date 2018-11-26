package main

type errorString struct {
	s string
}

func (e *errorString) Error() string{
	return e.s
}

func newError(text string) error {
	return &errorString{text}
}