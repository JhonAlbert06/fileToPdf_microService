# File To Pdf MicroService

## Project Description

The "fileToPdf_microService" project is a web service implemented in Go (Golang) using the Gin framework. Its main purpose is to convert files of various formats (such as office documents) to PDF files. The service exposes two main endpoints: one for file conversion and another for retrieving converted files.

### Main File: main.go

The main file of the project configures a web server using Gin and defines three routes:

1. **GET "/"**: A test route that returns a JSON message indicating that the service is working.

2. **POST "/convertFileToPdf"**: A route that accepts files via the POST method and converts them to PDF files. It uses the `controllers.ConvertFileToPdf` handler to manage the conversion logic.

3. **GET "/getFile/:fileName"**: A route that allows the download of previously converted PDF files. It uses the `controllers.GetFile` handler to manage this operation.

### Controller: controllers.go

The controller contains two main functions:

1. **ConvertFileToPdf(c *gin.Context)**: This function handles receiving a file, determining its type, and converting it to a PDF file if it is not already. It uses the "unoconv" command for the conversion. Additionally, it manages the storage and deletion of temporary files.

2. **GetFile(c *gin.Context)**: This function manages the download of previously converted PDF files. It takes the filename as a parameter from the URL and returns the corresponding file if it exists.

### Important Notes

Dependencies:

* Linux
  `sudo apt-get install libreoffice`
  `sudo apt-get install unoconv`

* Mac
  `brew install libreoffice`
  `brew install unoconv`

* Windows
  https://www.libreoffice.org/download/download/
