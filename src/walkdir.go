package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := os.Args[1]
	err := filepath.Walk(path ,  func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}

		if f.IsDir() {
			fmt.Printf("目录：%s\n", path)
		} else {
			fmt.Printf("文件：%s\n", path)
		}
		return nil
	})
	if err == nil{
		fmt.Println("filepath.Walk ok!")
	} else {
		fmt.Println("filepath.Walk err!")
	}
}