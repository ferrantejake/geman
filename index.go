package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	GnomeExtDir string // name of the object
}

func main() {
	fmt.Println("hello world")

	// config := (
	// 	GnomeExtDir = "~/.local/share/gnome-shell/"
	// )

	info, err := ioutil.ReadDir(config.GnomeExtDir)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for index := 0; index < len(info); index++ {
		fmt.Println(info[index])
	}
}
