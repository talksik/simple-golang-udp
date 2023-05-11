package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", "localhost:5000")
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	// send json data
	data := map[string]interface{}{
		"name":      "John",
		"audioData": 30,
	}
	marshalledData, err := json.Marshal(data)
	if err != nil {
		logrus.Error("Error while marshalling data", err)
		return
	}

	_, err = c.Write(marshalledData)
	if err != nil {
		logrus.Error("Error while writing data", err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply: %s\n", string(buffer[0:n]))
}
