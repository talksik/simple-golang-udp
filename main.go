package main

import (
	"math/rand"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

// port global variable
const (
	PORT = ":5000"
)

func udpServer() {
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		logrus.Error("Error while resolving udp address", err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		logrus.Error("Error while listening to udp address", err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	logrus.Infof("Starting listening for udp packets on port %s", PORT)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			logrus.Error("Error while reading from udp address", err)
			return
		}

		// go routine to process the request
		go handleRequest(buffer[:n], connection, addr)
	}
}

func handleRequest(data []byte, connection *net.UDPConn, addr *net.UDPAddr) {
	logrus.Info("Processing request")

  logrus.Info("===============")
	logrus.Infof("raw byte array -> ", data)
	logrus.Info("string -> ", string(data))
  logrus.Infof("addr -> %s\n", addr.String())
  logrus.Info("===============")

	respondUdp(connection, addr)
}

func respondUdp(conn *net.UDPConn, addr *net.UDPAddr) {
	message := "received message from you! " + addr.String()

	data := []byte(message)
	_, err := conn.WriteToUDP(data, addr)
	if err != nil {
		logrus.Error("Error while responding back to client", err)
		return
	}
}

func main() {
	udpServer()
}
