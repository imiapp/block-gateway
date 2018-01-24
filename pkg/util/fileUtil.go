package util

import (
	"os"
        "log"
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

func CreateFilePath(dirName string) bool {

     	checkPathFlag,err := PathExists(dirName)
        if(err != nil){
          return false;
        }
	if !checkPathFlag {
		createPath(dirName)
	}
	recheckPathFlag,err := PathExists(dirName)
	return recheckPathFlag
}

func createPath(dirName string) {
	//	dir, _ := os.Getwd()                        //当前的目录
	err := os.Mkdir(dirName, os.ModePerm) //在当前目录下生成md目录
	if err != nil {
		log.Println(err)
	}
	log.Println("创建目录" + dirName + "成功")
}

