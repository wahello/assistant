package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/tsundata/assistant/api/pb"
)

func TestUser_Authorization(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(app.User)
	if err != nil {
		t.Fatal(err)
	}
	rdb.Set(context.Background(), "user:auth:token", "test", time.Hour)

	repo := new(mocks.UserRepository)
	u := NewUser(rdb, repo)
	type args struct {
		ctx     context.Context
		payload *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{"case1", u, args{context.Background(), &pb.TextRequest{Text: "test"}}, &pb.StateReply{State: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Authorization(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Authorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Authorization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetRole(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(app.User)
	if err != nil {
		t.Fatal(err)
	}
	rdb.Set(context.Background(), "user:auth:token", "test", time.Hour)

	repo := new(mocks.UserRepository)
	repo.
		On("GetRole", mock.AnythingOfType("int")).
		Return(model.Role{Profession: "super"}, nil)
	u := NewUser(rdb, repo)
	type args struct {
		in0     context.Context
		payload *pb.RoleRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.RoleReply
		wantErr bool
	}{
		{"case1", u, args{context.Background(), &pb.RoleRequest{Id: 1}}, &pb.RoleReply{
			Role: &pb.Role{
				Profession: "super",
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetRole(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || got.Role.Profession != tt.want.Role.Profession {
				t.Errorf("User.GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
