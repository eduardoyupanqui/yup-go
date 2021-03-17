package filemanager

import (
	"io"
	"os"
)

func IsFileExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func IsDirectoryExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.Mode().IsDir()
}

func CreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func Copy(sourcePath string, destinationPath string) {
	srcFile, err := os.Open(sourcePath)
	if err != nil {
		//	fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
	defer srcFile.Close()

	destFile, err := os.Create(destinationPath) // creates if file doesn't exist
	if err != nil {
		//	fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		//	fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}

	err = destFile.Sync()
	if err != nil {
		//	fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}
