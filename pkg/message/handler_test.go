package message_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"mars-rover/pkg/message"
	"mars-rover/pkg/positioning"
	"mars-rover/pkg/positioning/mocks"
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

func TestProcessor(t *testing.T) {
	t.Parallel()
	type fields struct {
	}
	type args struct {
		repository positioning.Repository
		robot      *positioning.Robot
		text       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty message specified",
			fields:  fields{},
			args:    args{repository: nil, robot: nil, text: ""},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "invalid instruction is provided",
			fields:  fields{},
			args:    args{repository: &mocks.RepoMock{}, robot: positioning.NewRobot(), text: "1 something"},
			wantErr: true,
			errMsg:  "Cannot provide instruction",
		},
		{
			name:    "sets coordinate",
			fields:  fields{},
			args:    args{repository: &mocks.RepoMock{}, robot: positioning.NewRobot(), text: "1 2"},
			wantErr: false,
			errMsg:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := positioning.NewInteractor(tt.args.repository, tt.args.robot)
			handler := message.NewHandler(interactor)
			err := handler.Processor(tt.args.text)

			if tt.wantErr && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Processor() %s: got = %v, want = %v", tt.name, err, tt.errMsg)
				return
			}

			if !tt.wantErr {
				assert.Nil(t, err)
			}
		})
	}
}
