package main

import (
	"io"
	"log"
	"os"
)

const (
	BackupExt  = ".backup"
	MsgfFromTo = "from: %d, to: %d"
	//
	ErrInvalidArgs = "invalid args"
	ErrfFileRead   = "file read error: %s"
	ErrfFileOpen   = "src file open error: %s, %s"
)

func main() {

	if len(os.Args) < 2 {
		log.Println(ErrInvalidArgs)
		return
	}
	srcFile := os.Args[1]
	b, err := os.ReadFile(srcFile)
	if err != nil {
		log.Printf(ErrfFileRead, srcFile)
		return
	}

	old, err := Extract(b)
	if err != nil {
		log.Println(err.Error())
		return
	}
	new := old + 1
	newB := Replace(b, old, new)

	src, err := os.Open(srcFile)
	if err != nil {
		log.Printf(ErrfFileOpen, srcFile, err.Error())
		return
	}
	defer src.Close()

	dstFile := srcFile + BackupExt

	dst, err := os.Create(dstFile)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if err := os.WriteFile(srcFile, newB, 0644); err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf(MsgfFromTo, old, new)
}
