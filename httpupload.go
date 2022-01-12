package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.Static("/file", "saved")
	router.POST("/many", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			err := c.SaveUploadedFile(file, "saved/"+file.Filename)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(c.PostForm("key"))
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.POST("/one", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(file.Filename)

		err = c.SaveUploadedFile(file, "saved/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8000")
}
