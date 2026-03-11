// PROGRAM UTAMA
package main

// IMPORT GIN FRAMEWORK (TEST)

import (
	"todo-api-golang/database"
	"todo-api-golang/routes"

	"github.com/gin-gonic/gin"
)

// Menyimpan variabel untuk menyimpan data todo sementara (NANTI PAKAI DATABASE)
// FUNGSI MAIN
func main() {

	// membuat router server
	router := gin.Default()

	database.ConnectDatabase()

	routes.SetupTodoRoutes(router)

	// menjalankan server di port 8080
	router.Run(":8080")
}
