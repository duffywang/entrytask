package fileutils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/duffywang/entrytask/pkg/utils/hashutils"
)

func GetFileName(fileName string) string {
	ext := path.Ext(fileName)
	return hashutils.Hash(strings.TrimSuffix(fileName, ext)) + ext
}

func GetSavePath() string {
	return "upload/picfiles"
}

func CheckSavePathValid(dest string) bool {
	_, err := os.Stat(dest)
	return os.IsNotExist(err)
}

func CreateSavePath(dest string) error {
	err := os.MkdirAll(dest, os.ModePerm)
	return err
}

func SaveFileByte(file *[]byte, dest string) error {
	err := ioutil.WriteFile(dest, *file, 0777)
	return err
}

func CheckPermisson(dest string) bool {
	_, err := os.Stat(dest)
	return os.IsPermission(err)
}
