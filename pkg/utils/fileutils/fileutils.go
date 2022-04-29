package fileutils

import (
	"path"
	"strings"

	"github.com/duffywang/entrytask/pkg/utils/hashutils"
)

type FileType int

const (
	TypeImage FileType = iota + 1
)

func GetFileName(fileName string) string {
	ext := path.Ext(fileName)
	return hashutils.Hash(strings.TrimSuffix(fileName, ext)) + ext
}
