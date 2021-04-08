package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func IsProbExist(num int) bool {
	rangeFolderList, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, rangeFolder := range rangeFolderList {
		if rangeFolder.Name() == GetRangeOfProb(num) {
			files, err := ioutil.ReadDir(GetRangeOfProb(num))
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				if strings.Contains(file.Name(), strconv.Itoa(num)+"번") {
					if filerc, _ := os.Open(GetRangeOfProb(num) + "/" + file.Name() + "/" + strconv.Itoa(num) + ReadFileExtension()); filerc != nil {
						return true
					}
				}
			}

		}

	}

	return false
}
