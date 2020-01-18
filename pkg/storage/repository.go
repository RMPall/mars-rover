package storage

import (
	"errors"

	"mars-rover/pkg/positioning"
)

const (
	emptyRobotSuppliedErrMsgFmt = "Empty robot supplied"
)

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Save(robot *positioning.Robot) error {
	if robot == nil {
		return errors.New(emptyRobotSuppliedErrMsgFmt)
	}
	return nil
}
