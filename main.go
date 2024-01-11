package main

import (
	"fileToPdf_microService/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Is Working",
		})
	})

	r.POST("/convertFileToPdf", controllers.ConvertFileToPdf)
	r.GET("/getFile/:fileName", controllers.GetFile)

	// The port 3000 has been set in the .env file
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
