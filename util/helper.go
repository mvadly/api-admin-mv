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

func UploadFile(c *gin.Context, fileName, pathFolder string) (interface{}, string, error) {
	exist, err := Exists(pathFolder)
	if !exist {
		err := os.Mkdir(pathFolder, 0777)
		if err != nil {
			return nil, "can't create folder " + pathFolder, err
		}
	}

	file, err := c.FormFile(fileName)
	if err != nil {
		return nil, "form file undefined", err
	}

	fmt.Println("file: ", fileName)
	fmt.Println("file: ", file.Filename)
	fmt.Println("file: ", file.Size)
	fmt.Println("file: ", file.Header)

	fileSize := file.Size / 1024

	if fileSize > 500 {
		return nil, fmt.Sprintf("file size is %vkb and its too large, max 500kb", fileSize), fmt.Errorf("file size is too large")
	}

	fileType := strings.ToLower(filepath.Ext(file.Filename))
	imgAllow := InArray(fileType, typeExt)
	fmt.Println("file type: ", fileType)
	if imgAllow == false {
		return nil, "file type not support", fmt.Errorf("file type unknown")
	}

	filename := uuid.New().String() + fileType
	pathSave := pathFolder + filename

	if err := c.SaveUploadedFile(file, pathSave); err != nil {
		return nil, "error save file", err
	}

	fmt.Println("file saved")

	return "/" + pathSave, "file " + fileType + " saved", err
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
