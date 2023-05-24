package main

import (
	"log"

	"github.com/yakiroren/hotremote"
)

func main() {
	client, err := HotRemote.New()
	if err != nil {
		log.Fatal(err)
	}

	devices, err := client.ListDevices()
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		log.Println(device.ID, device.Name)
	}

	defer client.Close()

	// time.Sleep(100 * time.Second)
}
