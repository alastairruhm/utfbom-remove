package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dimchansky/utfbom"
	"github.com/urfave/cli"
)

func main() {
	var path string
	var checkOnly bool

	app := cli.NewApp()
	app.Name = "utfbom-remove"
	app.Version = "v1.0.0"
	app.Usage = "detect and remove BOM in utf-8 encoding files"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "check-only",
			Usage:       "dry-run mode",
			Destination: &checkOnly,
		},
		cli.StringFlag{
			Name:        "path",
			Value:       ".",
			Usage:       "the path to scan",
			Destination: &path,
		},
	}

	app.Action = func(c *cli.Context) error {
		absDir, err := filepath.Abs(path)
		if err != nil {
			fmt.Fprintf(c.App.Writer, "Error: %#v\n", err)
			return cli.NewExitError("error in path "+path, 1)
		}
		if checkOnly {
			files, err := ListFilesWithBOM(absDir)
			if err != nil {
				fmt.Println(err)
				return err
			}

			for _, file := range files {
				fmt.Println(file)
			}
		}
		err = RemoveBomForFiles(path)
		if err != nil {
			return err
		}
		return nil
	}

	app.Run(os.Args)
}

func RemoveUtfBom(byteData []byte) ([]byte, error) {

	// just skip BOM
	output, err := ioutil.ReadAll(utfbom.SkipOnly(bytes.NewReader(byteData)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return output, nil
}

func IsDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if info.IsDir() {
		return true, nil
	}
	return false, nil
}

func IsRugular(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if info.Mode().IsRegular() {
		return true, nil
	}
	return false, nil
}

// ListFilesWithBOM ...
func ListFilesWithBOM(path string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.Name() == ".git" { // filter .git subdirectory
			return filepath.SkipDir
		}
		isRegularFile, err := IsRugular(path)
		if err != nil {
			return err
		}
		if isRegularFile { // regular file
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				return err
			}
			output, err := RemoveUtfBom(data)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if bytes.Compare(output, data) != 0 {
				fileList = append(fileList, path)
			}

		}
		return nil
	})
	return fileList, err
}

func RemoveBomForFiles(path string) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.Name() == ".git" { // filter .git subdirectory
			return filepath.SkipDir
		}
		isRegularFile, err := IsRugular(path)
		if err != nil {
			return err
		}
		if isRegularFile { // regular file
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				return err
			}
			output, err := RemoveUtfBom(data)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if bytes.Compare(output, data) != 0 {
				err = ioutil.WriteFile(path, output, 0644)
				return err
			}

		}
		return nil
	})
	return err
}
