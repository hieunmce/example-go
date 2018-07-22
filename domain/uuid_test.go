// +build unit

package domain

import (
	"database/sql/driver"
	"errors"
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUUIDFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    UUID
		wantErr error
	}{
		{
			name:    "correct uuid",
			args:    args{"a98484cb-cc66-4687-8e66-837e5997c427"},
			want:    UUID(uuid.Must(uuid.FromString("a98484cb-cc66-4687-8e66-837e5997c427"))),
			wantErr: nil,
		},
		{
			name:    "correct uuid zero value",
			args:    args{"00000000-0000-0000-0000-000000000000"},
			want:    UUID(uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000000"))),
			wantErr: nil,
		},
		{
			name:    "incorrect uuid by wrong character",
			args:    args{"a98484cb-cc66-4687-8e66-837e5997c42l"},
			wantErr: errors.New("encoding/hex: invalid byte: U+006C 'l'"),
		},
		{
			name:    "incorrect uuid by invalid length",
			args:    args{"a98484cb-cc66-4687-8e66-837e5997c42"},
			wantErr: errors.New("uuid: incorrect UUID length: a98484cb-cc66-4687-8e66-837e5997c42"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UUIDFromString(tt.args.s)
			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("UUIDFromString() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if tt.wantErr.Error() != err.Error() {
					t.Errorf("UUIDFromString() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUIDFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_IsZero(t *testing.T) {
	tests := []struct {
		name string
		u    UUID
		want bool
	}{
		{
			name: "zero uuid",
			u:    UUID(uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000000"))),
			want: true,
		},
		{
			name: "normal uuid",
			u:    UUID(uuid.Must(uuid.FromString("c2186152-0e6e-437f-bf16-d0b291e98100"))),
			want: false,
		},
		{
			name: "nil uuid",
			u:    UUID{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (&tt.u).IsZero(); got != tt.want {
				t.Errorf("UUID.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		u       UUID
		want    []byte
		wantErr bool
	}{
		{
			name:    "success marshal json",
			u:       UUID(uuid.Must(uuid.FromString("c2186152-0e6e-437f-bf16-d0b291e98100"))),
			want:    []byte("\"c2186152-0e6e-437f-bf16-d0b291e98100\""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestUUID_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		u       UUID
		args    args
		wantErr bool
	}{
		{
			name:    "success unmarshal json",
			u:       UUID{},
			args:    args{[]byte("\"c2186152-0e6e-437f-bf16-d0b291e98100\"")},
			wantErr: false,
		},
		{
			name:    "fail unmarshal json by length",
			u:       UUID{},
			args:    args{[]byte("\"c2186152-0e6e-437f-bf16-d0b291e9810\"")},
			wantErr: true,
		},
		{
			name:    "fail unmarshal json by wrong character",
			u:       UUID{},
			args:    args{[]byte("\"c2186152-0e6e-437f-bf16-d0b291e981l\"")},
			wantErr: true,
		},
		{
			name:    "fail unmarshal json by missing double quote character",
			u:       UUID{},
			args:    args{[]byte("c2186152-0e6e-437f-bf16-d0b291e9810")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := (&tt.u).UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UUID.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUUID_Value(t *testing.T) {
	tests := []struct {
		name    string
		u       UUID
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "correct uuid",
			u:       UUID(uuid.Must(uuid.FromString("c2186152-0e6e-437f-bf16-d0b291e98100"))),
			want:    driver.Value("c2186152-0e6e-437f-bf16-d0b291e98100"),
			wantErr: false,
		},
		{
			name:    "zero uuid",
			u:       UUID(uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000000"))),
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_Scan(t *testing.T) {
	type args struct {
		b interface{}
	}
	tests := []struct {
		name    string
		u       *UUID
		args    args
		wantErr bool
	}{
		{
			name:    "success scan uuid",
			u:       &UUID{},
			args:    args{[]byte("5d38884d-2d3c-4662-8e1a-a91464a0d248")},
			wantErr: false,
		},
		{
			name:    "failed scan uuid by length",
			u:       &UUID{},
			args:    args{[]byte("5d38884d-2d3c-4662-8e1a-a91464a0d24")},
			wantErr: true,
		},
		{
			name:    "failed scan uuid by invalid character",
			u:       &UUID{},
			args:    args{[]byte("5d38884d-2d3c-4662-8e1a-a91464a0d24l")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Scan(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UUID.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
