package tcpsender

import (
	"bufio"
	"net"
	"time"
)

func Send(adrr string, data []byte) ([]byte, error) {

	var dialer = net.Dialer{Timeout: time.Second * 3}
	tcpServer, err := net.ResolveTCPAddr("tcp", adrr)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		return nil, err

	}

	conn, err := dialer.Dial("tcp", tcpServer.String())
	if err != nil {
		println("Dial failed:", err.Error())
		return nil, err
	}

	defer conn.Close()

	_, err = conn.Write([]byte(data))
	if err != nil {
		println("Write data failed:", err.Error())
		return nil, err
	}

	received, _ := bufio.NewReader(conn).ReadBytes(('\n'))

	conn.Close()

	return received, nil
}
