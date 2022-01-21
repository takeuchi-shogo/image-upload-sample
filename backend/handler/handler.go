package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Uploader(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("MultipartForm Error: ", err)
	}
	files := form.File["file"]

	for _, file := range files {
		err := c.SaveUploadedFile(file, "img/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}
