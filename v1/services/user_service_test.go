package services

import "testing"

func TestUserService_UserByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		u        *UserService
		args     args
		wantData string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{}
			gotData, err := u.UserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotData != tt.wantData {
				t.Errorf("UserService.UserByID() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
