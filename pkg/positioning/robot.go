package positioning

import (
	"fmt"
)

const (
	Disabled = "Disabled"
	Active   = "Active"
)

type Robot struct {
	ID          int
	Position    Position
	Instruction Instruction
	Status      string
}

type Position struct {
	Coordinate  *Coordinate
	Orientation *Orientation
}

// NewRobot returns a new robot with initial coordinates set to 0, 0 .
func NewRobot() *Robot {
	return &Robot{
		Position: Position{
			Coordinate: &Coordinate{
				X: 0,
				Y: 0,
			},
		},
		Status: Disabled,
	}
}

// SetInitialPosition sets initial position of the robot including orientation.
func (r *Robot) SetInitialPosition(x, y int, orientation *Orientation) {
	r.Position.Coordinate.X = x
	r.Position.Coordinate.Y = y
	r.Position.Orientation = orientation
	r.Status = Active
}

// MoveForward is an instruction to the robot to move forward one grid point, and maintain the same heading.
func (r *Robot) MoveForward(layoutCoordX, layoutCoordY int) error {
	var newCoordinate = *r.Position.Coordinate
	switch *r.Position.Orientation {
	case N:
		newCoordinate.Y = r.Position.Coordinate.Y + 1
	case E:
		newCoordinate.X = r.Position.Coordinate.X + 1
	case S:
		newCoordinate.Y = r.Position.Coordinate.Y - 1
	case W:
		newCoordinate.X = r.Position.Coordinate.X - 1
	default:
		break
	}
	r.Position.Coordinate = &newCoordinate
	return nil
}

// DisplayRobotStats displays a robot's final position in the rectangular grid.
func (r *Robot) DisplayRobotStats() {
	fmt.Println(r.Position.Coordinate.X, r.Position.Coordinate.Y, r.Position.Orientation.String())
}
