package main

import (
	"fmt"
	"github.com/matteo-magni/gobindings/utils"
	"log"
)

func main() {
	rootdir := utils.Getenv("ROOTDIR", ".")
	// maxdepth, err := strconv.Atoi(utils.Getenv("WALKDEPTH", "-1"))
	// if err != nil {
	// 	maxdepth = -1
	// }
	// fmt.Printf("WALKDEPTH=%v\n", maxdepth)

	binding, err := utils.GetBinding(rootdir, "redis")

	// filesmap, err := utils.ReadFiles(rootdir, maxdepth)
	if err != nil {
		log.Panic(err)
	}

	for k, v := range binding {
		fmt.Println()
		fmt.Println("=============")
		fmt.Printf("File: %s\n", k)
		fmt.Println(string(v))
		fmt.Println("=============")
	}

}
