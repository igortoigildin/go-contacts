package subscriber_tests

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	dto "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases"
	usecase "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases"
)

func Test_usecase_AcceptFriendRequest(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	type args struct {
		ctx         context.Context
		requestInfo *dto.FriendAcceptDTO
	}

	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		mock    func(t *testing.T) usecase.Deps
	}{
		{
			name: "Test 1. Positive",
			args: args{
				ctx: ctx, // dummy
				requestInfo: &dto.FriendAcceptDTO{
					Username:       "Alice",
					TargetUsername: "Bob",
				},
			},
			mock: func(t *testing.T) usecase.Deps {
				repoMock := mocks.NewSubscriberRepository(t)

				repoMock.On("AcceptRequest", ctx, mock.Anything).Return(nil)
				return usecase.Deps{
					SubscriberRepository: repoMock,
				}
			},
			wantErr: assert.NoError,
		},
		{
			name: "Test 2. Negative",
			args: args{
				ctx: ctx, // dummy
				requestInfo: &dto.FriendAcceptDTO{
					Username:       "Alice",
					TargetUsername: "Bob",
				},
			},
			mock: func(t *testing.T) usecase.Deps {
				repoMock := mocks.NewSubscriberRepository(t)

				repoMock.On("AcceptRequest", ctx, mock.Anything).Return(errors.New("some error")).Once()
				return usecase.Deps{
					SubscriberRepository: repoMock,
				}
			},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := usecase.NewUsecase(tt.mock(t))

			_, err := uc.AcceptFriendRequest(ctx, tt.args.requestInfo)

			tt.wantErr(t, err)
		})
	}
}
