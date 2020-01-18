package positioning_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"mars-rover/pkg/positioning"
)

func TestNewRobot_ReturnNonNilRobot(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, positioning.NewRobot())
}

func TestRobot_SetInitialPosition(t *testing.T) {
	t.Parallel()
	newRobot := positioning.NewRobot()
	orientation := positioning.E

	newRobot.SetInitialPosition(1, 2, &orientation)

	assert.Equal(t, newRobot.Position.Coordinate.X, 1)
	assert.Equal(t, newRobot.Position.Coordinate.Y, 2)
	assert.Equal(t, *newRobot.Position.Orientation, orientation)
}

func TestRobot_MoveForwardTowardsEast(t *testing.T) {
	t.Parallel()
	xCoordinate := 1
	yCordinate := 2
	newRobot := positioning.NewRobot()
	orientation := positioning.E
	newRobot.SetInitialPosition(xCoordinate, yCordinate, &orientation)

	err := newRobot.MoveForward(5, 6)

	assert.Nil(t, err)
	assert.Equal(t, newRobot.Position.Coordinate.X, xCoordinate+1)
	assert.Equal(t, newRobot.Position.Coordinate.Y, yCordinate)
}

func TestRobot_MoveForwardTowardsSouth(t *testing.T) {
	t.Parallel()
	xCoordinate := 1
	yCordinate := 2
	newRobot := positioning.NewRobot()
	orientation := positioning.S
	newRobot.SetInitialPosition(xCoordinate, yCordinate, &orientation)

	err := newRobot.MoveForward(5, 6)

	assert.Nil(t, err)
	assert.Equal(t, newRobot.Position.Coordinate.X, xCoordinate)
	assert.Equal(t, newRobot.Position.Coordinate.Y, yCordinate-1)
}

func TestRobot_MoveForwardTowardsWest(t *testing.T) {
	t.Parallel()
	xCoordinate := 1
	yCordinate := 2
	newRobot := positioning.NewRobot()
	orientation := positioning.W
	newRobot.SetInitialPosition(xCoordinate, yCordinate, &orientation)

	err := newRobot.MoveForward(5, 6)

	assert.Nil(t, err)
	assert.Equal(t, newRobot.Position.Coordinate.X, xCoordinate-1)
	assert.Equal(t, newRobot.Position.Coordinate.Y, yCordinate)
}

func TestRobot_MoveForwardTowardsNorth(t *testing.T) {
	t.Parallel()
	xCoordinate := 1
	yCordinate := 2
	newRobot := positioning.NewRobot()
	orientation := positioning.N
	newRobot.SetInitialPosition(xCoordinate, yCordinate, &orientation)

	err := newRobot.MoveForward(5, 6)

	assert.Nil(t, err)
	assert.Equal(t, newRobot.Position.Coordinate.X, xCoordinate)
	assert.Equal(t, newRobot.Position.Coordinate.Y, yCordinate+1)
}
