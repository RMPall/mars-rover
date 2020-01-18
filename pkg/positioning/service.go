package positioning

import (
	"errors"
	"fmt"

	"mars-rover/pkg/layout"
)

const (
	exceededCoordinateLimitErrMsg = "Coordinate limit of %v has been exceeded in either or both of the supplied inputs"
	invalidRobotInstructionErrMsg = "Unknown or invalid instruction is provided"
	failedToProccessInstruction   = "Failed to process instruction"
)

// Repository interface consists if save. Save saves a robot toa repository.
type Repository interface {
	Save(robot *Robot) error
}

type Interactor struct {
	Repository       Repository
	Robot            *Robot
	LayoutCoordinate *layout.Coordinate
}

// NewInteractor returns a new interactor. If repository is nil , then nil interactor is returned.
func NewInteractor(repository Repository, robot *Robot) *Interactor {
	if repository == nil {
		return nil
	}
	return &Interactor{
		Repository:       repository,
		Robot:            robot,
		LayoutCoordinate: &layout.Coordinate{X: 0, Y: 0},
	}
}

func (i *Interactor) SetLayout(x, y int) {
	i.LayoutCoordinate = layout.NewLayoutCoordinate(x, y)
}

// RobotPlacement places a robot on the specified coordinates with the specified orientation.
// If x or y coordinates are greater than the boundaries of the plateau then return error.
func (i *Interactor) RobotPlacement(x, y int, orientation *Orientation) error {
	if x > layout.MaximumCoordinateValue || y > layout.MaximumCoordinateValue {
		return fmt.Errorf(exceededCoordinateLimitErrMsg, layout.MaximumCoordinateValue)
	}
	i.Robot.SetInitialPosition(x, y, orientation)
	return nil
}

// GiveInstructions provides valid instructions to the robot.
// If an invalid instruction is specified then an error is returned.
func (i *Interactor) GiveInstructions(instructions []string) error {
	for _, instruction := range instructions {
		if !IsInstructionAllowed(Instruction(instruction)) {
			return errors.New(invalidRobotInstructionErrMsg)
		}

		robot, err := i.execute(instruction)
		if err != nil {
			break
		}

		err = i.Repository.Save(robot)
		if err != nil {
			return err
		}

	}
	return nil
}

func (i *Interactor) execute(instruction string) (*Robot, error) {
	switch instruction {
	case left:
		orientation := i.Robot.Position.Orientation.Rotate90DegreesLeft()
		i.Robot.Position.Orientation = &orientation
	case right:
		orientation := i.Robot.Position.Orientation.Rotate90DegreesRight()
		i.Robot.Position.Orientation = &orientation
	case forward:
		err := i.Robot.MoveForward(i.LayoutCoordinate.X, i.LayoutCoordinate.Y)
		if err != nil {
			return i.Robot, err
		}
	default:
		return nil, errors.New(failedToProccessInstruction)

	}
	return i.Robot, nil
}

// Display displays the current robot position.
func (i *Interactor) Display() {
	i.Robot.DisplayRobotStats()
}
