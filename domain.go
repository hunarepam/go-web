package main

type User struct {
	Id       int
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
