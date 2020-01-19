package message

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"mars-rover/pkg/positioning"
)

const (
	InitCoordinateMessage = iota + 2
	InitRobotPosition

	failedToParseErrMsgFmt   = "failed to parse message, due to %v"
	initalCoordsNotSetErrMsg = "initial coordinates not set"
	unknownOrientationErrMsg = "unknown orientation specified"
)

// Handler struct consists of an interactor , which can be used to interact with the usecases like
// defining the layout of the plateau, passing instructions for the robot etc.
type Handler struct {
	interactor *positioning.Interactor
}

// NewHandler returns a handler.If a nil interactor is provded, then
// a nil handler is returned.
func NewHandler(interactor *positioning.Interactor) *Handler {
	if interactor == nil {
		return nil
	}
	return &Handler{interactor: interactor}
}

// Processor processes the input text and passes the information to the business layers.
func (h *Handler) Processor(text string) error {
	message := strings.Split(strings.Replace(text, " ", "", -1), "")
	if len(message) == 0 {
		return nil
	}

	switch {
	case (len(message) == InitCoordinateMessage):
		horizontalAxis, verticalAxis, err := getXAndYAxisFromText(message[0], message[1])
		if err != nil {
			return err
		}

		h.interactor.SetLayout(horizontalAxis, verticalAxis)

	case (len(message) == InitRobotPosition):
		h.interactor.Robot = positioning.NewRobot()
		if h.interactor.LayoutCoordinate.X == 0 || h.interactor.LayoutCoordinate.Y == 0 {
			return errors.New(initalCoordsNotSetErrMsg)
		}
		horizontalAxis, verticalAxis, err := getXAndYAxisFromText(message[0], message[1])
		if err != nil {
			return err
		}

		orientation := positioning.StringToOrientation[message[2]]
		if orientation == positioning.Unknown {
			return errors.New(unknownOrientationErrMsg)
		}
		return h.interactor.RobotPlacement(horizontalAxis, verticalAxis, &orientation)

	default:
		err := h.interactor.GiveInstructions(message)
		if err != nil {
			return err
		}
		h.interactor.Display()
	}

	return nil
}

func getXAndYAxisFromText(horizontal, vertical string) (int, int, error) {
	horizontalAxis, err := strconv.Atoi(horizontal)
	if err != nil {
		return 0, 0, fmt.Errorf(failedToParseErrMsgFmt, err)
	}

	verticalAxis, err := strconv.Atoi(vertical)
	if err != nil {
		return 0, 0, fmt.Errorf(failedToParseErrMsgFmt, err)
	}

	return horizontalAxis, verticalAxis, nil
}
