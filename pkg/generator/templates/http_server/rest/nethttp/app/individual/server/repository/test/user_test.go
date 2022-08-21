package repository_test

import (
	"context"
	"reflect"
	"testing"

	"GO_TADY_PACKAGE_NAME/app/individual/server/entity"
	"GO_TADY_PACKAGE_NAME/app/individual/server/repository"
	"GO_TADY_PACKAGE_NAME/app/pkg/rdb"
)

const wantErr, noErr = true, false

func TestCreate(t *testing.T) {
	cases := map[string]struct {
		input   *entity.User
		want    *entity.User
		wantErr bool
	}{
		"success": {
			&entity.User{ID: 1, Name: "John Titer"},
			&entity.User{ID: 1, Name: "John Titer"},
			noErr,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			db := rdb.NewSQLHandler()
			repo := repository.NewUserRepository(db)

			got, err := repo.Create(context.TODO(), tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("FAIL: wantErr: %v, but got: %v", tt.wantErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("FAIL: wantErr: %v, but got: %v", tt.wantErr, err)
				}
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("FAIL: want: %v, but got: %v", tt.want, tt.want)
			}
		})
	}
}
