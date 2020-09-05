package main

import (
	"fmt"
)

func main() {
	fmt.Println("launch")
	snsHandler, err := Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	message := "# Pod Defect Alert\n"
	message += "\t 2 pod is not running\n"
	snsHandler.publishAlert(message)
}
