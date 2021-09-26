package main

import (
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/api"
	"github.com/gleich/silvanus/pkg/ask"
)

func main() {
	log := lumber.NewCustomLogger()
	log.Timezone = time.Local

	token, err := ask.Token()
	if err != nil {
		log.Fatal(err, "Failed to ask for token")
	}

	opts, err := ask.Options()
	if err != nil {
		log.Fatal(err, "Failed to ask options")
	}

	client := api.NewClient(token)

	_, err = api.Repos(log, client, opts)
	if err != nil {
		log.Fatal(err, "Failed to get data for repos")
	}
}
