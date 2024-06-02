package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/go-market/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
