package main

import (
	"context"
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Errorf(err.Error())
	}

	client := godo.NewFromToken(os.Getenv("DO_PAT"))
	acc, _, err := client.Account.Get(context.TODO())

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf(acc.Name)
}
