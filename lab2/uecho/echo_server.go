// Leave an empty line above this comment.
package main

import (
	"fmt"
	"net"
	"strings"
)

// UDPServer implements the UDP Echo Server specification found at
// https://github.com/COURSE_TAG/assignments/tree/master/lab2/README.md#udp-echo-server
type UDPServer struct {
	conn *net.UDPConn
	// TODO(student): Add fields if needed
}

// NewUDPServer returns a new UDPServer listening on addr. It should return an
// error if there was any problem resolving or listening on the provided addr.
func NewUDPServer(addr string) (*UDPServer, error) {
	
	conn, err1 := net.ResolveUDPAddr("udp",addr)
	if err1 != nil {
		return nil, err1
	}

	connUDP, err2 := net.ListenUDP("udp",conn)
	fmt.Println(connUDP)
	if err2 != nil {
		return nil, err2
	}
	udpserv := UDPServer{connUDP}

	return &udpserv,err2
}

// ServeUDP starts the UDP server's read loop. The server should read from its
// listening socket and handle incoming client requests as according to the
// the specification.
func (u *UDPServer) ServeUDP() {
		buff:= make([]byte,1024)
		defer u.conn.Close()
	for {
		n, addr, err := u.conn.ReadFromUDP(buff[0:])
		_= err
		fmt.Printf(string(rune(n)))
		inpArray := strings.Split(string(buff[:n]),"|:|")
		cmd := inpArray[0]
		fmt.Println(cmd)
		inptxt := inpArray[1]
		fmt.Println(inptxt)


		switch cmd {
		case "UPPER":
			outtxt := strings.ToUpper(inptxt)
			//u.conn.WriteToUDP([]byte(outtxt),&u.conn.LocalAddr())
			//fmt.Println(outtxt)
			u.conn.WriteToUDP([]byte(outtxt),addr)
		case "LOWER":
			//outtxt := strings.ToLower(inptxt)
			//u.conn.WriteToUDP([]byte(outtxt),addr)

		case "CAMEL":
			//temptxt := strings.Split(inptxt," ")
			//var outtxt string
			u.conn.WriteToUDP([]byte(inptxt),addr)
		
		case "ROT13":
			//temptxt := strings.Split(inptxt," ")
			//var outtxt string
			u.conn.WriteToUDP([]byte(inptxt),addr)
		
		case "SWAP":
			//temptxt := strings.Split(inptxt," ")
			//var outtxt string
			u.conn.WriteToUDP([]byte(inptxt),addr)
		}
	}
}

// socketIsClosed is a helper method to check if a listening socket has been
// closed.
func socketIsClosed(err error) bool {
	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}
	return false
}
