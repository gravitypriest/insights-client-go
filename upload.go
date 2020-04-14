package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"log"
	"os"
	"mime/multipart"
	"bytes"
	"path/filepath"
	"fmt"
	"net/textproto"
)

const uploadURL = "https://cloud.redhat.com/api/ingress/v1/upload"

// upload submits archivePath to the Insights service for analysis.
func upload(cfg *config, archivePath string) error {
	var err error

	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return err
	}

	client, err := newClient(cfg)
	if err != nil {
		return err
	}

	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	postbody := &bytes.Buffer{}
	writer := multipart.NewWriter(postbody)

	// file
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="file"; filename="%s"`, (filepath.Base(archivePath))))
	h.Set("Content-Type", "application/vnd.redhat.advisor.collection+tgz")
	filePart, err := writer.CreatePart(h)
	_, err = io.Copy(filePart, f)

	// metadata
	h = make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			"metadata", "metadata"))
	metadataPart, err := writer.CreatePart(h)
	metadataPart.Write([]byte("{}"))
	err = writer.Close()

	if err != nil {
		return err
	}
	
	req, err := http.NewRequest(http.MethodPost, uploadURL, postbody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if cfg.AuthMethod == "BASIC" {
		req.SetBasicAuth(cfg.Username, cfg.Password)
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusOK:
		break
	case http.StatusAccepted:
		log.Println("Uploaded successfully.")
		break
	default:
		return &unexpectedResponseErr{statusCode: res.StatusCode, body: string(data)}
	}

	return nil
}
