package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dimchansky/utfbom"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "utfbom-remove"
	app.Usage = "detect and remove BOM in utf-8 encoding files"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	app.Run(os.Args)
}

func mainBack() {
	checkOnly := flag.Bool("check", false, "only do files with BOM check")

	flag.Parse()
	fmt.Println()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	dir := flag.Args()[1]
	absDir, err := filepath.Abs(dir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(*checkOnly)
	if *checkOnly {
		files, err := ListFilesWithBOM(absDir)
		fmt.Println(files)
		if err != nil {
			fmt.Println(err)
		}

		for _, file := range files {
			fmt.Println(file)
		}
	}

}

func RemoveUTFBOM(byteData []byte) ([]byte, error) {

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
		isRegularFile, err := IsRugular(path)
		if err != nil {
			return err
		}
		if isRegularFile { // Is not regular file
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}
