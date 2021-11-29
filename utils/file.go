package util

import (
	"fmt"
	"log"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CheckPath(path string) {
	exist, err := PathExists(path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if !exist {
		fmt.Printf("no dir![%v]\n", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(fmt.Sprintf("mkdir failed![%v]\n", err))
		} else {
			log.Println(fmt.Sprintf("mkdir [%v] success!\n", path))
		}
	}
}
