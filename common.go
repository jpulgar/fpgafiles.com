package main

import (
	"bufio"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func findAllFiles(root, ext string, skip string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			if d.Name() == skip {
				return filepath.SkipDir
			}
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func CopyFile(source string, destination string) error {
	bytesRead, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	//Copy all the contents to the desitination file
	err = ioutil.WriteFile(destination, bytesRead, 0644)
	if err != nil {
		return err
	}

	return nil
}

func FileSize(filename string) (int64, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
