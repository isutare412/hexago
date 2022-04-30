package derror_test

import (
	"fmt"
	"testing"

	"github.com/isutare412/hexago/gateway/pkg/derror"
	"github.com/stretchr/testify/assert"
)

func TestBoundError(t *testing.T) {
	err := errorWithBind()
	sErr := derror.Unbind(err)
	assert.Error(t, err)
	assert.Equal(t, sErr, derror.ErrServiceUnavailable)
}

func TestUnboundError(t *testing.T) {
	err := errorWithoutBind()
	sErr := derror.Unbind(err)
	assert.Error(t, err)
	assert.NoError(t, sErr)
}

func errorWithBind() error {
	err := func() error {
		err := fmt.Errorf("error occured")
		return derror.Bind(err, derror.ErrServiceUnavailable)
	}()
	return fmt.Errorf("forward: %w", err)
}

func errorWithoutBind() error {
	err := func() error {
		return fmt.Errorf("error occured")
	}()
	return fmt.Errorf("forward: %w", err)
}
