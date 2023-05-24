package main

import (
	"log"
	"time"

	HotRemote "github.com/yakiroren/hotremote"
)

func main() {
	client, err := HotRemote.New()
	if err != nil {
		log.Fatal(err)
	}

	session := client.CreateSession("316501071005")

	session.SwitchToChannel(13)

	time.Sleep(5 * time.Second)

	session.SwitchToChannel(12)

	time.Sleep(5 * time.Second)

	defer client.Close()
}
