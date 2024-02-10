package main

import (
	"fileToPdf_microService/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Is Working",
		})
	})

	r.POST("/convertFileToPdf", controllers.ConvertFileToPdf)
	r.GET("/getFile/:fileName", controllers.GetFile)
	r.POST("/convertFile", controllers.ConvertAndReturnFile)

	// The port 3000 has been set in the .env file
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
