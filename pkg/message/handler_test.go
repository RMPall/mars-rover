package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"mars-rover/pkg/message"
	"mars-rover/pkg/positioning"
	"mars-rover/pkg/storage"
)

func TestNewHandler_ReturnNilHandler(t *testing.T) {
	t.Parallel()

	assert.Nil(t, message.NewHandler(nil))
}

func TestNewHandler_ReturnValidHandler(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, message.NewHandler(positioning.NewInteractor(storage.NewRepository(), positioning.NewRobot())))
}
