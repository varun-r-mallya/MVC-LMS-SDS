package models_test

import (
	"fmt"
	"testing"

	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/models"
	"github.com/varun-r-mallya/MVC-LMS-SDS/pkg/types"

	"github.com/joho/godotenv"
)

type RegisterUserTest struct {
	user          types.UserRegister
	expectedError error
	expectedBool  bool
}

var RegisterUserTests = []RegisterUserTest{
	{types.UserRegister{UserName: "like absolutely really really new user", HashedPassword: "password", Salt: "salt", IsAdmin: false}, nil, true},
	{types.UserRegister{UserName: "like absolutely really really new user", HashedPassword: "password", Salt: "salt", IsAdmin: false}, fmt.Errorf("user already exists"), false},
	{types.UserRegister{UserName: "like absolutely really really new user2", HashedPassword: "password", Salt: "salt", IsAdmin: false}, nil, true},
}

func TestRegisterUser(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Errorf("Error loading .env file")
	}

	for _, test := range RegisterUserTests {
		success, err := models.RegisterUser(test.user)
		if err != nil && test.expectedError == nil {
			if err.Error() != test.expectedError.Error() {
				t.Errorf("Expected error %s, but got %s", test.expectedError, err)
			}
		}
		if success != test.expectedBool {
			t.Errorf("Expected bool %t, but got %t", test.expectedBool, success)
		}
	}

}
