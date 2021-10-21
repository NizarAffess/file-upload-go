package handlersfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HandleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal("Error while uploading file: ", err)
		return
	}

	dst := "files/" + file.Filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Fatalln("Error while saving file to files directory: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{file.Filename: "Successfully uploaded"})
	fmt.Println("Upload successful")

}
