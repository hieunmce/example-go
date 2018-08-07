package book

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/luquehuong/example-go/domain"
)

func Test_validationMiddleware_Create(t *testing.T) {
	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.Book) error {
			return nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.Book
	}

	tests := []struct {
		name            string
		args            args
		wantErr         bool
		wantOutput      *domain.Book
		errorStatusCode int
	}{
		{
			name: "valid book",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
			}},
			wantOutput: &domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
			},
		},
		{
			name: "invalid book by missing name",
			args: args{&domain.Book{
				Description: "Curabitur vulputate vestibulum lorem.",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by missing category_id",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by name is empty",
			args: args{&domain.Book{
				Name: "",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "a",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "ab",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "abc",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "abcd",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "abcde",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by description is empty",
			args: args{&domain.Book{
				Description: "",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "a",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "ab",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "abc",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "abcd",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "abcde",
			}},
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
		want    []domain.Book
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
		ctx  context.Context
		book *domain.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.Find(tt.args.ctx, tt.args.book)
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
	type fields struct {
		Service Service
	}
	type args struct {
		p *domain.Book
	}
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.Book) (*domain.Book, error) {
			return p, nil
		},
	}

	defaultCtx := context.Background()

	tests := []struct {
		name            string
		args            args
		wantOutput      *domain.Book
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "valid book",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
			}},
			wantOutput: &domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
			},
		},
		{
			name: "invalid book by missing name",
			args: args{&domain.Book{
				Description: "Curabitur vulputate vestibulum lorem.",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by name is empty",
			args: args{&domain.Book{
				Name: "",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "a",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "ab",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "abc",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "abcd",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght name < 5 characters",
			args: args{&domain.Book{
				Name: "abcde",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by description is empty",
			args: args{&domain.Book{
				Description: "",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "a",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "ab",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "abc",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "abcd",
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by lenght description < 5 characters",
			args: args{&domain.Book{
				Description: "abcde",
			}},
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
				t.Errorf("ValidationMiddleware.Update()  = %v, want output %v, err = %v", gotOutput, tt.wantOutput, err)
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx  context.Context
		book *domain.Book
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
			if err := mw.Delete(tt.args.ctx, tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
