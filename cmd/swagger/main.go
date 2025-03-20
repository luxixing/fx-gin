package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 运行swag命令生成swagger文档
	cmd := exec.Command("swag", "init",
		"--generalInfo", "swagger.go",
		"--dir", "./cmd/swagger,./internal/transport/http/handler", // Include handler directory
		"--output", "./docs/swagger",
		"--parseDependency",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("生成Swagger文档失败: %s\n", err)
		fmt.Printf("输出: %s\n", output)
		return
	}

	fmt.Println("Swagger文档生成成功!")
}
