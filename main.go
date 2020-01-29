package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("file.txt")

	//MOST RECOMMENDED !!!
	if errors.Is(err, os.ErrPermission) {
		err = fmt.Errorf("you don`t have a permission to open file:%w", err)
		log.Println(err)

	} else if err == os.ErrPermission {
		err = fmt.Errorf("you don`t have a permission to open file:%w", err)
		log.Println(err)

	} else if err == os.ErrNotExist {
		err = fmt.Errorf("the file does not exist:%w", err)
		log.Println(err)

	} else if err != nil {

		//old school, version 1
		//panic(err)

		//version 2, more meaningful
		//err = fmt.Errorf("file could not be opened: %w", err)
		//log.Println(err)
	}
	defer f.Close()
}
