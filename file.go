package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func download(fileName string, URL string) error {
	//Get the response bytes from the url
	fmt.Println(fileName, "\n", URL)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		response, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("received non 200 response code")
		}
		//Create a empty file
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		//Write the bytes to the fiel
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}
	}
	return nil
}

// copy file from src to dst
func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// copy folder from src to dst
func copyFolder(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		err = copy(src+"/"+file.Name(), dst+"/"+file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}
