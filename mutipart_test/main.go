package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

const target_url = "http://127.0.0.1:5001/api/v0/add"

func main() {
	// postFile("text2.txt","TEXT.txt")
	//postFolder("test")
	PostPath("test", target_url)
}

func PostPath(path string, urlString string) {
	fp, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fi, _ := fp.Stat()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if fi.IsDir() {
		err = postDir(writer,path)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err = postFile(writer, path)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = writer.Close()
	if err != nil {
		fmt.Println (err)
		return
	}
	resp, err := http.Post(target_url, writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println("resp err : ", err)
		return
	}
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp read err :", err)
		return
	}
	fmt.Println(string(resp_body))
}

func postDir(writer *multipart.Writer, filename string) (err error) {
	uploadwriter, _ := CreateFormDirectory(writer, "file", filename)
	uploadfile, _ := os.Open(filename)
	defer func() { _ = uploadfile.Close() }()
	_, _ = io.Copy(uploadwriter, uploadfile)

	fis, err := ioutil.ReadDir(filename)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if fi.IsDir() {
			err = postDir(writer, filename+"/"+fi.Name())
			if err != nil {
				return err
			}
		} else {
			err = postFile(writer, filename+"/"+fi.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func postFolder(path string) {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	upload1writer, _ := CreateFormDirectory(writer, "file", path)
	uploadfile1, _ := os.Open(path)
	defer func() { _ = uploadfile1.Close() }()
	_, _ = io.Copy(upload1writer, uploadfile1)

	upload2writer, _ := writer.CreateFormFile("file", path+"/main.go")
	uploadfile2, _ := os.Open(path + "/main.go")
	defer func() { _ = uploadfile2.Close() }()
	_, _ = io.Copy(upload2writer, uploadfile2)

	upload3writer, _ := writer.CreateFormFile("file", path+"/loadFolder.go")
	uploadfile3, _ := os.Open(path + "/loadFolder.go")
	defer func() { _ = uploadfile3.Close() }()
	_, _ = io.Copy(upload3writer, uploadfile3)

	_ = writer.Close()
	//fmt.Println(body.String())

	resp, err := http.Post("http://127.0.0.1:5001/api/v0/add", writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err :", err)
		return
	}
	fmt.Println(string(resp_body))
}

func postFile2(fileName1 string, fileName2 string) {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	upload1writer, _ := writer.CreateFormFile("file", fileName1)
	uploadfile1, _ := os.Open(fileName1)
	defer func() { _ = uploadfile1.Close() }()
	_, _ = io.Copy(upload1writer, uploadfile1)

	upload2writer, _ := writer.CreateFormFile("file", fileName2)
	uploadfile2, _ := os.Open(fileName2)
	defer func() { _ = uploadfile2.Close() }()
	_, _ = io.Copy(upload2writer, uploadfile2)

	_ = writer.Close()
	fmt.Println(body.String())

	resp, err := http.Post(target_url, writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err :", err)
		return
	}
	fmt.Println(string(resp_body))
}

func postFile(writer *multipart.Writer, filename string) (err error) {
	uploadwriter, _ := writer.CreateFormFile("file", filename)
	uploadfile, _ := os.Open(filename)
	defer func() { _ = uploadfile.Close() }()
	_, err = io.Copy(uploadwriter, uploadfile)
	if err != nil {
		return err
	}
	return nil
}

func CreateFormDirectory(w *multipart.Writer, fieldname, filename string) (io.Writer, error) {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			quoteEscaper.Replace(fieldname), quoteEscaper.Replace(filename)))
	h.Set("Content-Type", "application/x-directory")
	return w.CreatePart(h)
}
