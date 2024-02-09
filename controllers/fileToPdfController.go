package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func ConvertFileToPdf(c *gin.Context) {

	// Get the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get program directory",
		})
		return
	}

	// Get the file path
	var fileDir string
	if os.Getenv("GO_ENV") == "production" {
		fileDir = filepath.Join(dir, "files")
	} else {
		fileDir = filepath.Join("files")
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(fileDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create image directory",
		})
		return
	}

	// Get the file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get file",
		})
		return
	}

	// Get the file extension
	extension := filepath.Ext(file.Filename)
	name := file.Filename[0 : len(file.Filename)-len(extension)]

	// Check if the file is a pdf
	if filepath.Ext(file.Filename) == ".pdf" {

		// Create the file path
		filePath := filepath.Join(fileDir, name+extension)

		// Save the file
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save the file",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"fileName": name + ".pdf",
		})
		return
	}

	// Create the file path
	officeFilePath := filepath.Join(fileDir, name+extension)
	pdfFilePath := filepath.Join(fileDir, name+".pdf")

	// Save office file
	if err := c.SaveUploadedFile(file, officeFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save the file",
		})
		return
	}

	// Convert the file to pdf
	cmd := exec.Command("unoconv", "--output", pdfFilePath, "--format", "pdf", officeFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		err = os.Remove(officeFilePath)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to convert to PDF", "e": extension,
		})
		return
	}

	// Remove the office file
	err = os.Remove(officeFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to remove the office file",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"fileName": name + ".pdf",
	})
}

func GetFile(c *gin.Context) {

	fileName := c.Param("fileName")

	file := filepath.Join("files", fileName+".pdf")

	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "file not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get file",
		})
		return
	}

	c.File(file)
}

func ConvertAndReturnFile(c *gin.Context) {

	// Get the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get program directory",
		})
		return
	}

	// Get the file path
	var fileDir string
	if os.Getenv("GO_ENV") == "production" {
		fileDir = filepath.Join(dir, "files")
	} else {
		fileDir = filepath.Join("files")
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(fileDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create image directory",
		})
		return
	}

	// Get the file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get file",
		})
		return
	}

	// Get the file extension
	extension := filepath.Ext(file.Filename)
	name := file.Filename[0 : len(file.Filename)-len(extension)]

	// Create the file path
	officeFilePath := filepath.Join(fileDir, name+extension)
	pdfFilePath := filepath.Join(fileDir, name+".pdf")

	// Save office file
	if err := c.SaveUploadedFile(file, officeFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save the file",
		})
		return
	}

	// Convert the file to pdf
	cmd := exec.Command("unoconv", "--output", pdfFilePath, "--format", "pdf", officeFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		err = os.Remove(officeFilePath)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to convert to PDF", "e": extension,
		})
		return
	}

	// Remove the office file
	err = os.Remove(officeFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to remove the office file",
		})
		return
	}

	// Read the PDF file
	pdfFile, err := ioutil.ReadFile(pdfFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read the PDF file",
		})
		return
	}

	// Delete the PDF file after a delay
	/*
		go func() {
			time.Sleep(5 * time.Second)
			err := os.Remove(pdfFilePath)
			if err != nil {
				return
			}
		}()
	*/

	// Return the PDF file
	c.Data(http.StatusOK, "application/pdf", pdfFile)

}
