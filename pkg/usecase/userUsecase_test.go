package usecase

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	mockUser "main.go/pkg/repository/mockRepo"
)

type eqCreateParamsMatcher struct {
	arg      helper.UserReq
	password string
}

func (e eqCreateParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(helper.UserReq)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}
func (e eqCreateParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)

}
func EqCreateParams(arg helper.UserReq, password string) gomock.Matcher {
	return eqCreateParamsMatcher{arg, password}
}
func TestUserSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := mockUser.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo)
	testData := []struct {
		name           string
		input          helper.UserReq
		buildStub      func(userRepo mockUser.MockUserRepository)
		expectedOutput response.UserData
		expectedError  error
	}{
		{
			name: "new user",
			input: helper.UserReq{
				Name:     "akshay",
				Email:    "akshaybabut1@gmail.com",
				Mobile:   "8592817810",
				Password: "1234",
			},
			buildStub: func(userRepo mockUser.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(
					EqCreateParams(helper.UserReq{
						Name:     "akshay",
						Email:    "akshaybabut1@gmail.com",
						Mobile:   "8592817810",
						Password: "1234",
					}, "1234")).Times(1).Return(response.UserData{
					Id:     1,
					Name:   "akshay",
					Email:  "akshaybabut1@gmail.com",
					Mobile: "8592817810",
				}, nil)
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "akshay",
				Email:  "akshaybabut1@gmail.com",
				Mobile: "8592817810",
			},
			expectedError: nil,
		},
		{
			name: "already exists",
			input: helper.UserReq{
				Name:     "akshay",
				Email:    "akshaybabut1@gmail.com",
				Mobile:   "8592817810",
				Password: "1234",
			},
			buildStub: func(userRepo mockUser.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(EqCreateParams(helper.UserReq{
					Name:     "akshay",
					Email:    "akshaybabut1@gmail.com",
					Mobile:   "8592817810",
					Password: "1234",
				}, "1234")).Times(1).Return(response.UserData{}, errors.New("user already exists"))

			},
			expectedOutput: response.UserData{},
			expectedError:  errors.New("user already exists"),
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userRepo)
			actualUser, err := userUseCase.UserSignUp(tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, actualUser, tt.expectedOutput)
		})
	}
}
