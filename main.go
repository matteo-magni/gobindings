package main

import (
	"fmt"
	"github.com/matteo-magni/gobindings/utils"
	"log"
)

func main() {
	rootdir := utils.Getenv("ROOTDIR", ".")

	bindings, err := utils.GetBindingsByType(rootdir, "redis")
	if err != nil {
		log.Panic(err)
	}

	allbindings, err := utils.GetAllBindings(rootdir)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(bindings)

	fmt.Println(bindings["redis"])

	fmt.Println(allbindings)

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
