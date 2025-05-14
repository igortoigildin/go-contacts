package subscriber_tests

import (
	"context"
	"testing"

	mocks "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	usecase "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases"
)

func Test_usecase_MakeFriendRequest(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	type args struct {
		ctx       context.Context
		orderInfo *usecase.FriendRequestDTO
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
				orderInfo: &usecase.FriendRequestDTO{
					Username:       "Alice",
					TargetUsername: "Bob",
				},
			},
			mock: func(t *testing.T) usecase.Deps {
				repoMock := mocks.NewSubscriberRepository(t)

				repoMock.On("MakeFriendRequest", ctx, mock.Anything).Return(nil).Maybe()
				return usecase.Deps{
					SubscriberRepository: repoMock,
				}
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := usecase.NewUsecase(tt.mock(t))

			err := uc.MakeFriendRequest(tt.args.ctx, tt.args.orderInfo)
			tt.wantErr(t, err)
		})
	}
}
