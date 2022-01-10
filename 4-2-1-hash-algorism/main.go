package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md5 := md5.Sum([]byte("sample\n"))

	fmt.Println(hex.EncodeToString(md5[:]))
}
