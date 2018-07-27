package category

import (
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"

	testutil "github.com/luantranminh/example-go/config/database/pg/util"
	"github.com/luantranminh/example-go/domain"
)

func Test_pgService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{&domain.Category{
				Name: "Create category",
			}},
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
	return
}

func Test_pgService_Update(t *testing.T) {
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
		t.Fatalf("Failed to create in update  category by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Category
		wantErr error
	}{
		{
			name: "success update",
			args: args{&domain.Category{
				Model: domain.Model{ID: category.ID},
				Name:  "nonameno",
			}},
			want: &domain.Category{
				Model: domain.Model{ID: category.ID},
				Name:  "fictiona",
			},
		},
		{
			name: "update is failed",
			args: args{&domain.Category{
				Model: domain.Model{ID: fakeCategoryID},
				Name:  "categorytest",
			}},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			got, err := s.Update(context.Background(), tt.args.p)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("pgService.Update() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
			}

			if (got != nil && tt.want != nil) && (got.ID != tt.want.ID) {
				t.Errorf("pgService.Update() got = %v, want %v", got, tt.want)
			}
		})
	}
	return
}

func Test_pgService_Find(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
		p   *domain.Category
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
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.Find(tt.args.in0, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_FindAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
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
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindAll(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_Delete(t *testing.T) {
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
		t.Fatalf("Failed to create category by error %v", err)
	}

	// dumpCategory := domain.Category{}
	// err = testDB.Create(&dumpCategory).Error
	// if err != nil {
	// 	t.Fatalf("Failed to create category by error %v", err)
	// }

	// dumpBook1 := domain.Book{
	// 	Name:       "This book need deleted",
	// 	CategoryID: category.ID,
	// }
	// err = testDB.Create(&dumpBook1).Error
	// if err != nil {
	// 	t.Fatalf("Failed to create book by error %v", err)
	// }

	// dumpBook2 := domain.Book{
	// 	Name:       "This book need deleted",
	// 	CategoryID: category.ID,
	// }
	// err = testDB.Create(&dumpBook2).Error
	// if err != nil {
	// 	t.Fatalf("Failed to create book by error %v", err)
	// }

	// dumpBook3 := domain.Book{
	// 	Name:       "Do not delete this trap",
	// 	CategoryID: dumpCategory.ID,
	// }
	// err = testDB.Create(&dumpBook3).Error
	// if err != nil {
	// 	t.Fatalf("Failed to create book by error %v", err)
	// }

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success delete",
			args: args{&domain.Category{
				Name:  "category",
				Model: domain.Model{ID: category.ID},
			}},
		},
		{
			name: "deletion failed: category not exists",
			args: args{&domain.Category{
				Model: domain.Model{ID: fakeCategoryID},
				Name:  "This is trap",
			}},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			err := s.Delete(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Delete() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
