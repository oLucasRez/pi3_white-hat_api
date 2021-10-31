package repositories

import (
	"context"
	"log"
)

var ctx = context.TODO()

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
