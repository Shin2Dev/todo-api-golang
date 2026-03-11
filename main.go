// PROGRAM UTAMA
package main

// IMPORT GIN FRAMEWORK (TEST)
import "github.com/gin-gonic/gin"

// FUNGSI MAIN
func main() {

	// membuat router server
	router := gin.Default()

	// endpoint pertama
	router.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "server berjalan",
		})

	})

	// menjalankan server di port 8080
	router.Run(":8080")
}
