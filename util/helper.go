package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
)

var typeExt = []string{".jpg", ".png", ".jpeg"}

func UploadFile(c *gin.Context, fileName, pathFolder string) (string, error) {
	exist, err := Exists(pathFolder)
	if !exist {
		err := os.Mkdir(pathFolder, 0777)
		if err != nil {
			return "can't create folder " + pathFolder, err
		}
	}

	file, err := c.FormFile(fileName)
	if err != nil {
		return "form file undefined", err
	}

	fileSize := file.Size / 1024

	if fileSize > 500 {
		return fmt.Sprintf("file size is %vkb and its too large, max 500kb", fileSize), fmt.Errorf("file size is too large")
	}

	fileType := strings.ToLower(filepath.Ext(file.Filename))
	imgAllow := InArray(fileType, typeExt)

	if imgAllow == false {
		return "file type not support", fmt.Errorf("file type unknown")
	}

	filename := pathFolder + uuid.New().String() + fileType
	if err := c.SaveUploadedFile(file, filename); err != nil {
		return "error save file", err
	}

	return "file " + fileType + " saved", err
}

func InArray(a interface{}, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetTypeFile(file string) string {
	buf, _ := ioutil.ReadFile(file)

	kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown {
		return "Unknown"
	}

	return strings.ToLower(kind.Extension)
}
