package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type IUserService interface {
	getUsers() []GormUser
	getUserById(id string) GormUser
	storeUser(user User) GormUser
	deleteUserById(id string) GormUser
}

type UserService struct {
	DB *gorm.DB
}

func (userService *UserService) getUsers() []GormUser {
	var users []GormUser
	userService.DB.Find(&users)
	return users
}

func (userService *UserService) getUserById(id string) GormUser {
	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err.Error())
	}
	var user GormUser
	userService.DB.First(&user, userId)
	return user
}

func (userService *UserService) storeUser(user User) GormUser {
	gormUser := GormUser{Username: user.Username, Password: user.Password}
	result := userService.DB.Create(&gormUser)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return gormUser
}

func (userService *UserService) deleteUserById(id string) GormUser {
	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err.Error())
	}
	var user GormUser
	userService.DB.Delete(&user, userId)
	return user
}

func initUserService() IUserService {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SERVER"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't connect to DB via Gorm:", err.Error())
		panic(err)
	}

	var userService IUserService = &UserService{db}

	return userService

}
