package Campaign

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestAddCampaign(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddCampaign(tt.args.db); (err != nil) != tt.wantErr {
			}
		})
	}
}

func TestLoadCampaign(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want []Campaign
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadCampaign(tt.args.db); !reflect.DeepEqual(got, tt.want) {
			}
		})
	}
}
