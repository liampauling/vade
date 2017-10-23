package vade

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"log"
	"path/filepath"
	"io"
	"archive/zip"
	"fmt"
)

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func DownloadZip(bucket string, eventTypeId string, marketId string) string {
	var key = fmt.Sprintf("marketdata/streaming/%v/%v.zip", eventTypeId, marketId)
	var downloadFileZip = fmt.Sprintf("/tmp/%v.zip", marketId)
	var destinationFolder = fmt.Sprintf("/tmp/%v", marketId)

	log.Println("Downloading:", key)

	file, err := os.Create(downloadFileZip)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	awsSession := session.Must(session.NewSession())

	downloader := s3manager.NewDownloader(awsSession)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		})
	if err != nil {
		panic(err)
	}

	log.Println("Downloaded:", numBytes)

	unzip(downloadFileZip, destinationFolder)

	return destinationFolder
}

func FileList(downloadFolder string) ([]string) {
	fileList := []string{}
	err := filepath.Walk(downloadFolder, func(path string, f os.FileInfo, err error) error {
		fi, err := os.Stat(path)
		if err != nil {
			panic(err)
		}
		switch mode := fi.Mode(); {
		case mode.IsRegular():
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return fileList
}
