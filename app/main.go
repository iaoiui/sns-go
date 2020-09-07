package main

import (
	"fmt"

	snsutil "github.com/iaoiui/sns-go/app/snsutil"
	// snsutil "github.com/iaoiui/sns-go/app"
)

func main() {
	fmt.Println("launch")
	snsHandler, err := snsutil.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	message := "# Pod Defect Alert\n"
	message += "\t 2 pod is not running\n"
	snsHandler.PublishAlert(message)
}
