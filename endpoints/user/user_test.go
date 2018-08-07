// +build unit

package user

import (
	"context"
	"testing"

	"github.com/luquehuong/example-go/domain"
	"github.com/luquehuong/example-go/service"
	userService "github.com/luquehuong/example-go/service/user"
)

func TestMakeUpdateEndpoint(t *testing.T) {
	mock := service.Service{
		UserService: &userService.ServiceMock{
			UpdateFunc: func(_ context.Context, p *domain.User) (*domain.User, error) {
				return p, nil
			},
		},
	}

	type args struct {
		req UpdateRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "update user endpoint parsed success",
			args: args{
				UpdateRequest{
					UpdateData{
						ID:    domain.MustGetUUIDFromString("415179ad-8067-4138-9b0d-41e0c68d4376"),
						Name:  "Aliquam feugiat tellus ut neque.",
						Email: "example@gmail.com",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeUpdateEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeCreateEndpoint(t *testing.T) {
	mock := service.Service{
		UserService: &userService.ServiceMock{
			CreateFunc: func(_ context.Context, p *domain.User) error {
				return nil
			},
		},
	}

	type args struct {
		req CreateRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "update user endpoint parsed success",
			args: args{
				CreateRequest{
					CreateData{
						Name:  "Aliquam feugiat tellus ut neque.",
						Email: "example@gmail.com",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeCreateEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeFindUserEndPoint(t *testing.T) {
	mock := service.Service{
		UserService: &userService.ServiceMock{
			FindFunc: func(_ context.Context, p *domain.User) (*domain.User, error) {
				return p, nil
			},
		},
	}
	type args struct {
		req FindRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "find user endpoint parsed success",
			args: args{
				FindRequest{
					UserID: domain.MustGetUUIDFromString("cfa930f4-0f37-4d61-9314-5c2cb0993e44"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeFindEndPoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeFindAllEndpoint(t *testing.T) {
	mock := service.Service{
		UserService: &userService.ServiceMock{
			FindAllFunc: func(_ context.Context) ([]domain.User, error) {
				return []domain.User{}, nil
			},
		},
	}
	type args struct {
		req FindAllRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "find all user endpoint parsed success",
			args: args{
				FindAllRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeFindAllEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeDeleteEndpoint(t *testing.T) {
	mock := service.Service{
		UserService: &userService.ServiceMock{
			DeleteFunc: func(_ context.Context, p *domain.User) error {
				return nil
			},
		},
	}

	type args struct {
		req DeleteRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "delete user endpoint parsed success",
			args: args{
				DeleteRequest{
					UserID: domain.MustGetUUIDFromString("cfa930f4-0f37-4d61-9314-5c2cb0993e44"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeDeleteEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
