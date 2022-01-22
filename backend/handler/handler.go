package handler

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Uploader(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("MultipartForm Error: ", err)
	}
	// s := form.Value
	// fmt.Println("s: ", s)
	files := form.File["file"]
	fmt.Println("file name: ", files[0].Filename)
	fmt.Println("header: ", files[0].Header)
	fmt.Println("size: ", files[0].Size)

	f, err := files[0].Open()
	fmt.Println("err: ", err)
	fmt.Println("f: ", f)
	// 画像をアップロードしてURLを取得したい
	for _, file := range files {
		err := c.SaveUploadedFile(file, "img/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	_, format, err := image.DecodeConfig(f)
	if err != nil {
		fmt.Println("multipart err: ", err)
	}
	fmt.Println("format: ", format)

	c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}
