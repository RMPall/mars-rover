package integration_tests

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DATA-DOG/godog"
	"mars-rover/pkg/layout"
	"mars-rover/pkg/positioning"
	"mars-rover/pkg/storage"
)

type instructor struct {
	coordinate *layout.Coordinate
	interactor *positioning.Interactor
}

func newInstructor() *instructor {
	interactor := positioning.NewInteractor(storage.NewRepository(), positioning.NewRobot())
	return &instructor{
		interactor: interactor,
	}
}

func (i *instructor) theLayoutOfThePlateauIsDetermined() error {
	i.coordinate = layout.NewLayoutCoordinate(20, 20)
	return nil
}

func (i *instructor) iSendARequestToPositionTheRobotWithinThePlateauBoundaries(details string) error {
	message := strings.Split(strings.Replace(details, " ", "", -1), "")
	if len(message) == 0 {
		return nil
	}

	x, err := strconv.Atoi(message[0])
	if err != nil {
		return err
	}

	y, err := strconv.Atoi(message[1])
	if err != nil {
		return err
	}

	orientation := positioning.StringToOrientation[message[2]]
	err = i.interactor.RobotPlacement(x, y, &orientation)
	return err
}

func (i *instructor) iSendASetOfInstructionsToTheRobot(instruction string) error {
	message := strings.Split(strings.Replace(instruction, " ", "", -1), "")
	if len(message) == 0 {
		return nil
	}

	err := i.interactor.GiveInstructions(message)
	return err
}

func (i *instructor) iAmAbleToSeeTheFinalPositionOfTheRobotInThePlateau(output string) error {
	message := strings.Split(strings.Replace(output, " ", "", -1), "")
	if len(message) == 0 {
		return nil
	}
	x, err := strconv.Atoi(message[0])
	if err != nil {
		return err
	}

	if i.interactor.Robot.Position.Coordinate.X != x {
		return fmt.Errorf("expected: '%v', got: '%v'", i.interactor.Robot.Position.Coordinate.X, x)
	}

	y, err := strconv.Atoi(message[1])
	if err != nil {
		return err
	}

	if i.interactor.Robot.Position.Coordinate.Y != y {
		return fmt.Errorf("expected: '%v', got: '%v'", i.interactor.Robot.Position.Coordinate.Y, y)
	}

	orientation := positioning.StringToOrientation[message[2]]

	if *i.interactor.Robot.Position.Orientation != orientation {
		return fmt.Errorf("expected: '%v', got: '%v'", *i.interactor.Robot.Position.Orientation, orientation)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	// Before & After
	var i *instructor = newInstructor()
	s.Step(`^the layout of the plateau is determined$`, i.theLayoutOfThePlateauIsDetermined)
	s.Step(`^I send a request "([^"]*)" to position the robot within the plateau boundaries$`, i.iSendARequestToPositionTheRobotWithinThePlateauBoundaries)
	s.Step(`^I send a set of "([^"]*)" instructions to the robot$`, i.iSendASetOfInstructionsToTheRobot)
	s.Step(`^I am able to see the final "([^"]*)" position of the robot in the plateau$`, i.iAmAbleToSeeTheFinalPositionOfTheRobotInThePlateau)
}
