package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"realhttpworld"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// ファイル以外は、WriteFieldを使用する
	writer.WriteField("name", "Michael Jackson")

	// ファイルPart
	// 読み込み
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	rf, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer rf.Close()
	io.Copy(fileWriter, rf)
	writer.Close()

	resp, err := http.Post("http://localhost:8888/", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}

	realhttpworld.Write(resp)
}
