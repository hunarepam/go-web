package main

import (
	//"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load configuration from .env file")
	}
}

func setupRouter(userController UserController) *gin.Engine {
	router := gin.Default()

	router.GET("/users", userController.getAllUsers)
	router.GET("/users/:userId", userController.getUserById)
	router.POST("/users", userController.storeUser)
	router.DELETE("/users/:userId", userController.deleteUserById)

	return router
}

func main() {
	/*connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")
	defer conn.Close()

	count, err := ReadEmployees(conn)
	if err != nil {
		log.Fatal("ReadEmployees failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", count) */

	router := setupRouter(getUserContoller())

	router.Run()

}
