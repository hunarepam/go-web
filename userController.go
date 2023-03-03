package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService IUserService
}

func (userController *UserController) getAllUsers(c *gin.Context) {
	users := userController.userService.getUsers()
	c.JSON(200, users)
}

func (userController *UserController) getUserById(c *gin.Context) {
	userId := c.Param("userId")
	users := userController.userService.getUserById(userId)
	c.JSON(200, users)
}

func (userController *UserController) storeUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser := userController.userService.storeUser(user)
	c.JSON(201, createdUser)
}

func (userController *UserController) deleteUserById(c *gin.Context) {
	userId := c.Param("userId")
	users := userController.userService.deleteUserById(userId)
	c.JSON(204, users)
}

func getUserContoller() UserController {
	var userController = UserController{userService: initUserService()}
	return userController
}
