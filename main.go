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
		lumber.Fatal(err, "Failed to ask for token")
	}

	_, err = ask.Options()
	if err != nil {
		lumber.Fatal(err, "Failed to ask options")
	}

	api.NewClient(token)
}
