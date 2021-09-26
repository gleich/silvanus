package main

import (
	"fmt"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/silvanus/pkg/ask"
)

func main() {
	log := lumber.NewCustomLogger()
	log.Timezone = time.Local

	_, err := ask.Token()
	if err != nil {
		lumber.Fatal(err, "Failed to ask for token")
	}

	opts, err := ask.Options()
	if err != nil {
		lumber.Fatal(err, "Failed to ask options")
	}

	fmt.Println(opts)
}
