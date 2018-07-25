// +build unit

package user

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/neverdiefc/example-go/domain"
)

func Test_validationMiddleware_Update(t *testing.T) {
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.User) (*domain.User, error) {
			return p, nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.User
	}
	tests := []struct {
		name            string
		args            args
		wantOutput      *domain.User
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "valid user",
			args: args{&domain.User{
				Name:  "Curabitur vulputate vestibulum lorem.",
				Email: "example@gmail.com",
			}},
			wantOutput: &domain.User{
				Name:  "Curabitur vulputate vestibulum lorem.",
				Email: "example@gmail.com",
			},
		},
		{
			name: "invalid user by missing name",
			args: args{&domain.User{
				Email: "example@gmail.com",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid user by missing email",
			args: args{&domain.User{
				Name: "Curabitur vulputate vestibulum lorem.",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid user by wrong email format",
			args: args{&domain.User{
				Name:  "Curabitur vulputate vestibulum lorem.",
				Email: "wrong email format",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name:            "invalid user by missing attribute",
			args:            args{&domain.User{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			gotOutput, err := mw.Update(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr == false {
					t.Errorf("validationMiddleware.Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Update() error %v doesn't implement StatusCode()", err)
				}
				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Update() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}

			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("ValidationMiddleware.Update() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_validationMiddleware_Create(t *testing.T) {
	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.User) error {
			return nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.User
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "valid user",
			args: args{&domain.User{
				Name:  "Curabitur vulputate vestibulum lorem.",
				Email: "example@gmail.com",
			}},
		},
		{
			name:            "invalid user by missing name",
			args:            args{&domain.User{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid user by missing email",
			args: args{&domain.User{
				Name: "Curabitur vulputate vestibulum lorem.",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid user by wrong email format",
			args: args{&domain.User{
				Name:  "Curabitur vulputate vestibulum lorem.",
				Email: "wrong email format",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name:            "invalid user by missing attribute",
			args:            args{&domain.User{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			err := mw.Create(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr == false {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Create() error %v doesn't implement StatusCode()", err)
				}
				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Create() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}
		})
	}
}

func Test_validationMiddleware_Find(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
		p   *domain.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput *domain.User
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			gotOutput, err := mw.Find(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("validationMiddleware.Find() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_validationMiddleware_FindAll(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput []domain.User
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			gotOutput, err := mw.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("validationMiddleware.FindAll() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
		p   *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			if err := mw.Delete(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
