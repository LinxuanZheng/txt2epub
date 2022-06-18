package src

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

func GenBook(fileName string) {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	book := NewBook(file)
	err = book.LoadBook()
	if err != nil {
		fmt.Println(err)
	}

	currTime := time.Now().Unix()
	tmpDir := book.Config.Title + "-" + strconv.FormatInt(currTime, 16)
	workDir = workDir + "/" + tmpDir

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err = os.Mkdir(workDir, 0777)
	if err != nil {
		panic(err)
	}

	e := Epub{
		Book:       book,
		OutputPath: workDir,
	}

	e.GenContainerXML()

	e.GenMimetype()

	e.GenOpfContent()

	e.GenChapters()

	e.GenNCX()

	e.CopyCover()

	e.GenZip()
}
