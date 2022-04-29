package hashutils

import (
	"crypto/md5"
	"encoding/hex"
)


func Hash(str string)string {
	hash := md5.Sum([]byte("salt"+str))
	//数组转切片 hash[:]
	return hex.EncodeToString(hash[:])
}