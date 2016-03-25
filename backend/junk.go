package main

import (
	"encoding/json"
	"fmt"
)

// Message is an exported type that
// contains the json instance
type Message struct {
	Name string
	Data interface{}
}

// Channel is exported type that
// contains the json instance for a Channel
type Channel struct {
	ID   string
	Name string
}

func main() {
	recRawMsg := []byte(`{"name":"channel add",` +
		`"data":{"name":"Hardware Support"}}`)
	var recMessage Message
	err := json.Unmarshal(recRawMsg, &recMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMessage)
	if recMessage.Name == "channel add" {
		addChannel(recMessage.Data)
	}

}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	channelMap := data.(map[string]interface{})
	channel.Name = channelMap["name"].(string)
	channel.ID = "1"
	fmt.Printf("%#v\n", channel)
	return channel, nil
}
