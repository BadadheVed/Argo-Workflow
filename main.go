package main

import (
	"fmt"

	"github.com/BadadheVed/argo-flow/routes"
)

func main() {
	fmt.Println("Hello, World!")

	r := routes.SetupRouter()
	r.Run(":8080") // Start the server
}
