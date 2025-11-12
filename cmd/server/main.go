package main

import (
	"fmt"

	cfg "github.com/sariya23/manage_pr_service/internal/config"
)

func main() {
	config := cfg.MustLoad()
	fmt.Println(config)
}
