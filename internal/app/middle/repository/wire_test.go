// +build wireinject

package repository

import (
	"reflect"
	"testing"
)

func TestCreateMiddleRepository(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    MiddleRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateMiddleRepository(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMiddleRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMiddleRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
