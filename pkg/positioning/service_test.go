package positioning_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mars-rover/pkg/positioning"
	"mars-rover/pkg/positioning/mocks"
)

func TestNewInteractor(t *testing.T) {
	t.Parallel()
	type fields struct {
	}
	type args struct {
		repository positioning.Repository
		robot      *positioning.Robot
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "returns nil interactor",
			fields:  fields{},
			args:    args{repository: nil, robot: nil},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "returns valid interactor",
			fields:  fields{},
			args:    args{repository: &mocks.RepoMock{}, robot: &positioning.Robot{}},
			wantErr: false,
			errMsg:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := positioning.NewInteractor(tt.args.repository, tt.args.robot)
			if tt.args.repository == nil {
				assert.Nil(t, interactor)
				return
			}
			assert.NotNil(t, interactor)

		})
	}
}

func TestGiveInstructions(t *testing.T) {
	t.Parallel()
	orientation := positioning.Orientation(2)
	type fields struct {
	}
	type args struct {
		instruction []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "invalid instruction",
			fields:  fields{},
			args:    args{instruction: []string{"something"}},
			wantErr: true,
			errMsg:  "invalid instruction",
		},
		{
			name:    "valid instruction",
			fields:  fields{},
			args:    args{instruction: []string{string(positioning.AllowedInstructions[0])}},
			wantErr: false,
			errMsg:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRobot := positioning.NewRobot()
			newRobot.SetInitialPosition(1, 2, &orientation)
			orientation := positioning.N
			newRobot.Position.Orientation = &orientation
			repoMock := &mocks.RepoMock{}
			interactor := positioning.NewInteractor(repoMock, newRobot)
			repoMock.On("Save", mock.Anything).Return(nil)

			err := interactor.GiveInstructions(tt.args.instruction)

			if tt.wantErr && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("GetInstruction() %s: got = %v, want = %v", tt.name, err, tt.errMsg)
				return
			}

			if !tt.wantErr {
				assert.Nil(t, err)
			}

		})
	}
}

func TestInteractor_RobotPlacement_returnsError(t *testing.T) {
	t.Parallel()
	interactor := positioning.NewInteractor(&mocks.RepoMock{}, positioning.NewRobot())
	orientation := positioning.E

	assert.NotNil(t, interactor.RobotPlacement(100, 0, &orientation))
}

func TestInteractor_RobotPlacement_returnsNilError(t *testing.T) {
	t.Parallel()
	interactor := positioning.NewInteractor(&mocks.RepoMock{}, positioning.NewRobot())
	orientation := positioning.E

	assert.Nil(t, interactor.RobotPlacement(5, 0, &orientation))
}
