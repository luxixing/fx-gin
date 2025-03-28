package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Run swag command to generate swagger documentation
	cmd := exec.Command("swag", "init",
		"--generalInfo", "swagger.go",
		"--dir", "./cmd/swagger,./internal/transport/http/handler", // Include handler directory
		"--output", "./docs/swagger",
		"--parseDependency",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to generate Swagger documentation: %s\n", err)
		fmt.Printf("Output: %s\n", output)
		return
	}

	fmt.Println("Swagger documentation generated successfully!")
}
