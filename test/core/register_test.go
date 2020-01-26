package test

import (
	"authorization-service/core"
	"errors"
	"testing"
)

// NewRepository : return new stub
func NewRepository(resultSaveUser error, resultGetUser core.User, resultGetUserError error) core.UserRepository {
	return &repositoryStub{resultSaveUser, resultGetUser, resultGetUserError}
}

type repositoryStub struct {
	resultSaveUser     error
	resultGetUser      core.User
	resultGetUserError error
}

func (r repositoryStub) SaveUser(user core.User) error {
	return r.resultSaveUser
}

func (r repositoryStub) GetUserByEmail(user core.User) (core.User, error) {
	return r.resultGetUser, r.resultGetUserError
}

func Test_RegisterWithError(t *testing.T) {
	repo := NewRepository(errors.New("error creating new user"), core.User{}, nil)
	srv := core.NewUserService(repo)
	expectedResult := false
	expectedError := errors.New("error creating new user")
	result, err := srv.UserRegister(&core.User{})
	if result != expectedResult {
		t.Error("saved user not posible")
	}
	if err.Error() != expectedError.Error() {
		t.Error("saved user no error ocurred")
	}
}

func Test_RegisterUserExistingUser(t *testing.T) {
	repo := NewRepository(nil, core.User{Email: "email@gmail.com", Password: "123", Name: "name", Age: 18}, nil)
	srv := core.NewUserService(repo)
	expectedResult := false
	expectedError := errors.New("user already exist")
	result, err := srv.UserRegister(&core.User{Email: "email@gmail.com", Password: "123", Name: "name", Age: 18})
	if result != expectedResult {
		t.Error("saved user not posible")
	}
	if err.Error() != expectedError.Error() {
		t.Error("saved user no error ocurred")
	}
}

func Test_RegisterUserNoUsers(t *testing.T) {
	repo := NewRepository(nil, core.User{}, nil)
	srv := core.NewUserService(repo)
	expectedResult := true
	result, err := srv.UserRegister(&core.User{Email: "email@gmail.com", Password: "123", Name: "name", Age: 18})
	if result != expectedResult {
		t.Error("saved user not posible")
	}
	if err != nil {
		t.Error("saved user error ocurred")
	}
}