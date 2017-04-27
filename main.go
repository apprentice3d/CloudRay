package main

import (
	"fmt"
	"github.com/apprentice3d/CloudRay/service"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	fmt.Printf("Started server on port %s", port)

	server := service.NewServer()
	server.Run(":" + port)

}
