package Event

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestAddEvents(t *testing.T) {
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
			if err := AddEvents(tt.args.db); (err != nil) != tt.wantErr {
			}
		})
	}
}

func TestLoadEvents(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want []Event
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadEvents(tt.args.db); !reflect.DeepEqual(got, tt.want) {
			}
		})
	}
}
