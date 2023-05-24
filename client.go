package HotRemote

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

//go:embed certificates/*
var certificatesDir embed.FS

type Client struct {
	conn    *websocket.Conn
	ID      string
	results chan []byte
}

func New() (*Client, error) {
	targetURL := url.URL{Scheme: "wss", Host: "192.168.1.3:7682"}

	dialer, err := createWSDialer()
	if err != nil {
		return nil, fmt.Errorf("websocket creation failed: %w", err)
	}

	conn, _, err := dialer.Dial(targetURL.String(), getHeaders())
	if err != nil {
		return nil, fmt.Errorf("dial failed: %w", err)
	}
	hexString, err := generateHexString(16)
	if err != nil {
		return nil, fmt.Errorf("failed to generate client ID: %w", err)
	}

	client := &Client{
		conn:    conn,
		results: make(chan []byte),
		ID:      hexString,
	}

	go client.readLoop()

	return client, nil
}

func (c *Client) readLoop() {
	for {
		message, err := c.readMsg()
		if err != nil {
			log.Fatal(err)
		}

		if msgIsKA(message) {
			if err := c.handleKeepAliveMessage(message); err != nil {
				log.Fatal(err)
			}

			continue
		}

		output, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}

		c.results <- output

	}
}

func (c *Client) readMsg() (map[string]interface{}, error) {
	response := map[string]interface{}{}

	err := c.conn.ReadJSON(&response)
	if err != nil {
		return nil, err
	}

	var message map[string]interface{}

	// special case for listDevices devices, where the response doesn't have the params field
	if response["MethodName"] == ListDevicesAction {
		message = response
	} else {
		message = response["Params"].(map[string]interface{})
	}
	return message, err
}

func msgIsKA(message map[string]interface{}) bool {
	return message["Action"] == "UtilsSDK"
}

func (c *Client) handleKeepAliveMessage(message map[string]interface{}) error {
	output, err := json.Marshal(message)
	if err != nil {
		return err
	}

	b := &KeepAliveResponse{}
	if err := json.Unmarshal(output, b); err != nil {
		return err
	}

	return nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
