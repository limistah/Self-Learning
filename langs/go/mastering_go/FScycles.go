package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var visited []int

func main() {

}

func walkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil
	}
	fileInfo, _ = os.Lstat(path)
	mode := fileInfo.Mode()

	// Find regular directories first
	if mode.IsDir() {
		abs, _ := filepath.Abs(path)
		_, ok := visited[abs]
		if ok {
			fmt.Println("Found cycle:", abs)
			return nil
		}
		visited[abs]++
		return nil
	}
	// Find symbolic links to directories
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		temp, err := os.Readlink(path)
		if err != nil {
			fmt.Println("os.Readlink():", err)
			return err
		}
		newPath, err := filepath.EvalSymlinks(temp)
		if err != nil {
			return nil
		}
		linkFileInfo, err := os.Stat(newPath)
		if err != nil {
			return err
		}
		linkMode := linkFileInfo.Mode()
		if linkMode.IsDir() {
			fmt.Println("Following...", path, "-->", newPath)
			_, ok := visited[abs]
			if ok {
				fmt.Println("Found cycle!", abs)
				return nil
			}
			visited[abs]++
			err = filepath.Walk(newPath, walkFunction)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
