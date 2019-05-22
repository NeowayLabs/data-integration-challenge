package models

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestCompany_HasID(t *testing.T) {
	id := uuid.NewV4()
	companyWithID := Company{}
	companyWithID.ID = &id

	tests := []struct {
		name string
		obj  Company
		want bool
	}{
		{
			"No ID",
			Company{},
			false,
		},
		{
			"With ID",
			companyWithID,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.obj.HasID(); got != tt.want {
				t.Errorf("Company.HasID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompany_IsEmpty(t *testing.T) {
	type fields struct {
		ID      *uuid.UUID
		Name    string
		Zip     string
		Website *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company := Company{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				Zip:     tt.fields.Zip,
				Website: tt.fields.Website,
			}
			if got := company.IsEmpty(); got != tt.want {
				t.Errorf("Company.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Normal",
			args{"Aaaaa"},
			false,
		},
		{
			"Empty",
			args{""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("validateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateZip(t *testing.T) {
	type args struct {
		zip string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Normal",
			args{"12345"},
			false,
		},
		{
			"Invalid (letter)",
			args{"a"},
			true,
		},
		{
			"Invalid (length)",
			args{"1"},
			true,
		},
		{
			"Empty",
			args{""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateZip(tt.args.zip); (err != nil) != tt.wantErr {
				t.Errorf("validateZip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateWebsite(t *testing.T) {
	valid := "https://www.senado.leg.br/noticias/tv/"
	invalid := "kek"
	type args struct {
		website *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Normal",
			args{&valid},
			false,
		},
		{
			"Nil",
			args{nil},
			false,
		},
		{
			"Normal",
			args{&invalid},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateWebsite(tt.args.website); (err != nil) != tt.wantErr {
				t.Errorf("validateWebsite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
