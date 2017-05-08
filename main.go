package main

import (
	"flag"
	"log"
	"os"

	"github.com/convyr/http/que"
)

var (
	channel string
)

func init() {
	que.Flags()
	flag.StringVar(&channel, "channel", os.Getenv("CONVYR_CHANNEL"), "channel for convyr communication")

}

func main() {
	flag.Parse()
	q, err := que.New(channel)
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()
	log.Fatal(q.Watch())
}
