package derror

import "errors"

type boundError struct {
	error
	secondary error
}

func Bind(primary, secondary error) error {
	return boundError{
		error:     primary,
		secondary: secondary,
	}
}

func Unbind(err error) error {
	var berr boundError
	if !errors.As(err, &berr) {
		return nil
	}
	return berr.secondary
}
