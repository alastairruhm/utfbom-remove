package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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

// RemoveUtfBom remove the bom header with given bytes
func RemoveUtfBom(byteData []byte) ([]byte, error) {
	// just skip BOM
	output, err := ioutil.ReadAll(utfbom.SkipOnly(bytes.NewReader(byteData)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return output, nil
}

// IsDirectory ...
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

// IsRugular returns true if the path given is a regular file
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
		if strings.Contains(path, ".git") { // filter .git subdirectory
			return filepath.SkipDir
		}
		b, _, err := DetectBom(path)
		if err != nil {
			return err
		}
		if b {
			fileList = append(fileList, path)
			return err
		}
		return nil
	})
	return fileList, err
}

// RemoveBomForFiles ...
func RemoveBomForFiles(path string) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".git") { // filter .git subdirectory
			return filepath.SkipDir
		}
		b, output, err := DetectBom(path)
		if err != nil {
			return err
		}
		if b {
			err = ioutil.WriteFile(path, output, 0644)
			return err
		}
		return nil
	})
	return err
}

// DetectBom detect bom of file as the path
// returns true and content of byte array as file after remove the bom
func DetectBom(path string) (bool, []byte, error) {
	isRegularFile, err := IsRugular(path)
	if err != nil {
		return false, nil, err
	}
	if isRegularFile { // regular file
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return false, nil, err
		}
		output, err := RemoveUtfBom(data)
		if err != nil {
			return false, nil, err
		}
		if bytes.Compare(output, data) != 0 {
			return true, output, nil
		}
	}
	return false, nil, nil
}
