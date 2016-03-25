package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
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
		channel, err := addChannel(recMessage.Data)
		// Added to the database here
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMsg, err := json.Marshal(sendMessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(sendRawMsg))
	}

}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}
	channel.ID = "1"
	fmt.Printf("%#v\n", channel)
	return channel, nil
}
