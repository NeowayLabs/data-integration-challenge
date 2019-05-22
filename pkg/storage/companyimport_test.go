package storage

import (
	"testing"

	"github.com/jean-lopes/data-integration-challenge/pkg/configs"
)

func TestLoadFromCSV(t *testing.T) {
	str, err := OpenPostgreSQLStorage(configs.PgConfig{})
	if err != nil {
		t.Fatalf("Error: %v, %v", err, configs.PgConfig{})
	}
	defer str.Close()

	type args struct {
		store Company
		path  string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO errouu
		//		{
		//			"Invalid path",
		//			args{str, "lul.csv"},
		//			true,
		//		},
		{
			"normal",
			args{str, "q1_catalog.csv"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str.Clean()
			if err := LoadFromCSV(tt.args.store, "../../test/"+tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadFromCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
