package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"main.go/pkg/api/middleware"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	mockUser "main.go/pkg/usecase/mockUsecase"
)

func TestUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)

	userUseCase := mockUser.NewMockUserUseCase(ctrl)
	cartUseCase := mockUser.NewMockCartUsecase(ctrl)
	UserHandler := NewUserHandler(userUseCase, cartUseCase)

	testData := []struct {
		name             string
		loginData        helper.LoginReq
		buildStub        func(userUsecase mockUser.MockUserUseCase)
		expectedCode     int
		expectedResponse response.Response
		expectedError    error
	}{
		{
			name: "valid login",
			loginData: helper.LoginReq{
				Email:    "akshaybabut1@gmail.com",
				Password: "1234",
			},
			buildStub: func(userUsecase mockUser.MockUserUseCase) {
				userUsecase.EXPECT().UserLogin(helper.LoginReq{
					Email:    "akshaybabut1@gmail.com",
					Password: "1234",
				}).Times(1).
					Return("validToken", nil)
			},
			expectedCode: 200,
			expectedResponse: response.Response{
				StatusCode: 200,
				Message:    "login successfully",
				Data:       nil,
				Errors:     nil,
			},
			expectedError: nil,
		},
		{
			name: "invalid login",
			loginData: helper.LoginReq{
				Email:    "invalid@example.com",
				Password: "invalidPassword",
			},
			buildStub: func(userUsecase mockUser.MockUserUseCase) {
				userUsecase.EXPECT().UserLogin(helper.LoginReq{
					Email:    "invalid@example.com",
					Password: "invalidPassword",
				}).Times(1).
					Return("", errors.New("invlid credentials"))
			},
			expectedCode: 400,
			expectedResponse: response.Response{
				StatusCode: 400,
				Message:    "failed to login",
				Data:       nil,
				Errors:     "invalid credentials",
			},
			expectedError: errors.New("invalid credentials"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)
			engine := gin.Default()
			recorder := httptest.NewRecorder()
			engine.POST("/user/login", UserHandler.UserLogin)
			var body []byte
			body, err := json.Marshal(tt.loginData)
			assert.NoError(t, err)
			url := "/user/login"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			engine.ServeHTTP(recorder, req)
			var actual response.Response
			err = json.Unmarshal(recorder.Body.Bytes(), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
		})
	}
}

func TestViewProfile(t *testing.T) {
	ctrl := gomock.NewController(t)

	userUseCase := mockUser.NewMockUserUseCase(ctrl)
	cartUseCase := mockUser.NewMockCartUsecase(ctrl)
	UserHandler := NewUserHandler(userUseCase, cartUseCase)

	testData := []struct {
		name             string
		userID           int64
		buildStub        func(userUsecase mockUser.MockUserUseCase)
		expectedCode     int
		expectedResponse response.Response
		expectedData     response.Userprofile
		expectedError    error
	}{
		{
			name:   "valid profile",
			userID: 1,
			buildStub: func(userUsecase mockUser.MockUserUseCase) {
				userUseCase.EXPECT().ViewProfile(gomock.Any()).Times(1).
					Return(response.Userprofile{
						Name:   "TestUser",
						Email:  "test@example.com",
						Mobile: "8592817810",
						Address: response.Address{
							House_number: "3044",
							Street:       "mgo",
							City:         "calicut",
							District:     "calicut",
							Landmark:     "MG",
							Pincode:      679534,
							IsDefault:    true,
						},
					}, nil)
			},
			expectedCode: 200,
			expectedResponse: response.Response{
				StatusCode: 200,
				Message:    "profile",
				Data: response.Userprofile{
					Name:   "TestUser",
					Email:  "test@example.com",
					Mobile: "8592817810",
					Address: response.Address{
						House_number: "3044",
						Street:       "mgo",
						City:         "calicut",
						District:     "calicut",
						Landmark:     "MG",
						Pincode:      679534,
						IsDefault:    true,
					},
				},
				Errors: nil,
			},
			expectedData: response.Userprofile{
				Name:   "TestUser",
				Email:  "test@example.com",
				Mobile: "8592817810",
				Address: response.Address{
					House_number: "3044",
					Street:       "mgo",
					City:         "calicut",
					District:     "calicut",
					Landmark:     "MG",
					Pincode:      679534,
					IsDefault:    true,
				},
			},
			expectedError: nil,
		},
		{
			name:   "invalid profile",
			userID: 2,
			buildStub: func(userUsecase mockUser.MockUserUseCase) {
				userUseCase.EXPECT().ViewProfile(gomock.Any()).Times(1).Return(response.Userprofile{}, errors.New("user not found"))
			},
			expectedCode: 400,
			expectedResponse: response.Response{
				StatusCode: 400,
				Message:    "Can't find userprofile",
				Data:       nil,
				Errors:     "user not found",
			},
			expectedData:  response.Userprofile{},
			expectedError: errors.New("user not found"),
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)
			engine := gin.Default()
			recorder := httptest.NewRecorder()
			engine.GET("/user/profile/", middleware.TestUserAuth, UserHandler.ViewProfile)
			url := "/user/profile/"
			req := httptest.NewRequest(http.MethodGet, url, nil)
			engine.ServeHTTP(recorder, req)
			var actual response.Response
			err := json.Unmarshal(recorder.Body.Bytes(), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
		})
	}
}
