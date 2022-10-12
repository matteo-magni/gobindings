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

	// fmt.Println(binding)

	fmt.Println(binding["redis"])
	// for k, v := range binding {
	// 	fmt.Printf("Binding found: %s\n", k)
	// 	for k1,v1 := range v {
	// 		fmt.Println()
	// 		fmt.Println("=============")
	// 		fmt.Printf("File: %s\n", k1)
	// 		fmt.Println(string(v1))
	// 		fmt.Println("=============")	
	// 	}
	// }

}
