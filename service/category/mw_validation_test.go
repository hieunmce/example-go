package category

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/luantranminh/example-go/domain"
)

func Test_validationMiddleware_Create(t *testing.T) {
	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.Category) error {
			return nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.Category
	}

	tests := []struct {
		name            string
		args            args
		wantErr         error
		errorStatusCode int
	}{
		{
			name: "valid category",
			args: args{&domain.Category{
				Name: "this is a valid category",
			}},
			wantErr:         nil,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name:            "invalid category by missing name",
			args:            args{&domain.Category{}},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid category by name'length < 5",
			args: args{&domain.Category{
				Name: "a",
			}},
			wantErr:         ErrNameTooShort,
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
				if tt.wantErr != err {
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

func Test_validationMiddleware_FindAll(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Category
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validationMiddleware_Find(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx      context.Context
		category *domain.Category
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Category
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.Find(tt.args.ctx, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validationMiddleware_Update(t *testing.T) {
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.Category) (*domain.Category, error) {
			return p, nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		category *domain.Category
	}
	tests := []struct {
		name            string
		args            args
		wantOutput      *domain.Category
		wantErr         error
		errorStatusCode int
	}{
		{
			name: "valid user",
			args: args{&domain.Category{
				Name: "Nameisvalid",
			}},
			wantOutput: &domain.Category{
				Name: "Nameisvalid",
			},
		},
		{
			name:            "invalid category by missing name",
			args:            args{&domain.Category{}},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},

		{
			name: "invalid category by new name to short",
			args: args{&domain.Category{
				Name: "No",
			}},
			wantErr:         ErrNameTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}

			got, err := mw.Update(defaultCtx, tt.args.category)
			if err != nil {
				if tt.wantErr != err {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("validationMiddleware.Update() = %v, want %v", got, tt.wantOutput)
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	serverMock := &ServiceMock{
		DeleteFunc: func(_ context.Context, p *domain.Category) error {
			return nil
		},
	}
	defaultCtx := context.Background()
	type args struct {
		category *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "sucess delete",
			args:    args{&domain.Category{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serverMock,
			}
			if err := mw.Delete(defaultCtx, tt.args.category); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
