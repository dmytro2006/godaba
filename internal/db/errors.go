package db

type UnknownDriverError struct {
	err string
}

func (e UnknownDriverError) Error() string {
	return e.err
}
