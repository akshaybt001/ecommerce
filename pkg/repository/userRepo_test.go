package repository

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

func TestUserSignup(t *testing.T) {
	test := []struct {
		name           string
		input          helper.UserReq
		expectedOutput response.UserData
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{
			name: "successful creations",
			input: helper.UserReq{
				Name:     "akshay",
				Email:    "akshaybabut1@gmail.com",
				Mobile:   "8592817810",
				Password: "1234",
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "akshay",
				Email:  "akshaybabut1@gmail.com",
				Mobile: "8592817810",
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"id", "name", "email", "mobile"}).
					AddRow(1, "akshay", "akshaybabut1@gmail.com", "8592817810")

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("akshay", "akshaybabut1@gmail.com", "8592817810", "1234").
					WillReturnRows(row)
			},
			expectedErr: nil,
		},
		{
			name: "duplicate user",
			input: helper.UserReq{
				Name:     "akshay",
				Email:    "akshaybabut1@gmail.com",
				Mobile:   "8592817810",
				Password: "1234",
			},
			expectedOutput: response.UserData{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users(.+)$").
					WithArgs("akshay", "akshaybabut1@gmail.com", "8592817810", "1234").
					WillReturnError(errors.New("email or phone number already used"))
			},
			expectedErr: errors.New("email or phone number already used"),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			userRepository := NewUserRespository(gormDB)

			tt.buildStub(mock)

			actualOutput, actualErr := userRepository.UserSignUp(tt.input)

			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else {
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}

		})
	}
}
