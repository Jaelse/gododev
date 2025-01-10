package do

import (
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"github.com/joho/godotenv"
)

type DoClient struct {
	client *godo.Client
}

func NewClient() DoClient {
	err := godotenv.Load()

	if err != nil {
		_ = fmt.Errorf(err.Error())
	}

	client := godo.NewFromToken(os.Getenv("DO_PAT"))

	return DoClient{
		client: client,
	}
}
