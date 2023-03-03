package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubUserService struct {
}

type getUsersType func() []GormUser
type getUserByIdType func(id string) GormUser
type storeUserMock func(user User) GormUser
type deleteUserType func(id string) GormUser

var getUsersMock getUsersType
var getByIdMock getUserByIdType
var storeMock storeUserMock
var deleteMock deleteUserType

func (u *StubUserService) getUsers() []GormUser {
	return getUsersMock()
}

func (u *StubUserService) getUserById(id string) GormUser {
	return getByIdMock(id)
}

func (userService *StubUserService) storeUser(user User) GormUser {
	return storeMock(user)
}

func (userService *StubUserService) deleteUserById(id string) GormUser {
	return deleteMock(id)
}

type getUserStruct struct {
	name      string
	function  getUsersType
	userCount int
}

func TestGetAllUsers(t *testing.T) {
	var cases []getUserStruct = []getUserStruct{
		{
			name: "empty response",
			function: func() []GormUser {
				return []GormUser{}
			},
			userCount: 0,
		},
		{
			name: "response with users",
			function: func() []GormUser {
				return []GormUser{{Username: "1", Password: "1"}, {Username: "2", Password: "2"}}
			},
			userCount: 2,
		},
	}
	for _, Case := range cases {
		t.Run(Case.name, func(t *testing.T) {
			getUsersMock = Case.function

			var userService IUserService = &StubUserService{}
			userController := UserController{userService: userService}
			router := setupRouter(userController)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users", nil)
			router.ServeHTTP(w, req)

			var users []GormUser
			json.Unmarshal(w.Body.Bytes(), &users)

			assert.Equal(t, 200, w.Code)
			assert.Equal(t, Case.userCount, len(users))
		})
	}
}

type getUserByIdStruct struct {
	name       string
	function   getUserByIdType
	resultUser GormUser
}

func TestGetUserById(t *testing.T) {
	var user GormUser
	var foundedUser GormUser = GormUser{Username: "1", Password: "1"}
	var cases []getUserByIdStruct = []getUserByIdStruct{
		{
			name: "empty response",
			function: func(id string) GormUser {
				return user
			},
			resultUser: user,
		},
		{
			name: "response with user",
			function: func(id string) GormUser {
				return foundedUser
			},
			resultUser: foundedUser,
		},
	}
	for _, Case := range cases {
		t.Run(Case.name, func(t *testing.T) {
			getByIdMock = Case.function

			var userService IUserService = &StubUserService{}
			userController := UserController{userService: userService}
			router := setupRouter(userController)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/1", nil)
			router.ServeHTTP(w, req)

			var actualUser GormUser
			json.Unmarshal(w.Body.Bytes(), &actualUser)

			assert.Equal(t, 200, w.Code)
			assert.Equal(t, Case.resultUser, actualUser)
		})
	}

}
