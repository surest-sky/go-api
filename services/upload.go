package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-api/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Upload(c *gin.Context) {
	path := utils.UploadPath()
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	dst := filepath.Join(path, cast.ToString(time.Now().Unix()) + "-" + file.Filename)
	_ = c.SaveUploadedFile(file, dst)

	url := filepath.Join(os.Getenv("APP_URL"), dst)
	c.JSON(200, map[string]interface{}{
		"message": "success",
		"code":    http.StatusCreated,
		"data": map[string]interface{}{
			"url": url,
		},
	})
}