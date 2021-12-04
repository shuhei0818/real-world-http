package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
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
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data: name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	rf, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer rf.Close()
	io.Copy(fileWriter, rf)

	resp, err := http.Post("http://localhost:8888/", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}

	realhttpworld.Write(resp)
}
