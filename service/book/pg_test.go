package book

import (
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	testutil "github.com/luquehuong/example-go/config/database/pg/util"
	"github.com/luquehuong/example-go/domain"
)

func Test_pgService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create dummy table category by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.Book{
					Name:        "Create New Book 1",
					Category_id: category.ID,
					Description: "description",
				},
			},
		},
		{
			name: "Category_id is not correct",
			args: args{
				&domain.Book{
					Name:        "Create New Book 1",
					Category_id: fakeCategoryID,
					Description: "description",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if err := s.Create(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pgService_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
		p   *domain.Book
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
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.Update(tt.args.in0, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
		p   *domain.Book
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
			s := &pgService{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.in0, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
