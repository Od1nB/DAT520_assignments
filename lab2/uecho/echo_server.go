// Leave an empty line above this comment.
package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// UDPServer implements the UDP Echo Server specification found at
// https://github.com/COURSE_TAG/assignments/tree/master/lab2/README.md#udp-echo-server
type UDPServer struct {
	conn *net.UDPConn
}

// NewUDPServer returns a new UDPServer listening on addr. It should return an
// error if there was any problem resolving or listening on the provided addr.
func NewUDPServer(addr string) (*UDPServer, error) {
	
	conn, err1 := net.ResolveUDPAddr("udp",addr)
	if err1 != nil {
		return nil, err1
	}

	connUDP, err2 := net.ListenUDP("udp",conn)
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
		fmt.Println(string(buff[:n]))
		inpArray := strings.Split(string(buff[:n]),"|:|")
		if len(inpArray) != 2{u.conn.WriteToUDP([]byte("Unknown command"),addr)} 
		cmd := inpArray[0]
		inptxt := inpArray[1]


		switch cmd {
		case "UPPER":
			outtxt := strings.ToUpper(inptxt)
			u.conn.WriteToUDP([]byte(outtxt),addr)
		case "LOWER":
			outtxt := strings.ToLower(inptxt)
			u.conn.WriteToUDP([]byte(outtxt),addr)

		case "CAMEL":
			temptxt := strings.Split(inptxt," ")
			var outtxt string
			for i := 0; i < len(temptxt); i++ {
				if len(temptxt[i]) == 1 {
					outtxt += strings.ToUpper(temptxt[i])+" "

				}else{
					outtxt += strings.ToUpper(temptxt[i][:1])+ strings.ToLower(temptxt[i][1:])+" "
				}
			}
			outtxt = outtxt[:len(outtxt)-1]
			u.conn.WriteToUDP([]byte(outtxt),addr)
		
		case "ROT13":
			b1 := make([]byte,1024)
			r := rot13Reader{strings.NewReader(inptxt)}
			outb,err := r.Read(b1)
			if err == nil {
				outtxt := string(b1[:outb])
				u.conn.WriteToUDP([]byte(outtxt),addr)
			}
		
		case "SWAP":
			outtxt := strings.Map(swapC,inptxt)
			u.conn.WriteToUDP([]byte(outtxt),addr)
		default:
			outtxt := "Unknown command"
			u.conn.WriteToUDP([]byte(outtxt),addr)
		}
		
	}
}


//Implemented a function from go playground: https://play.golang.org/p/6kzKnWG7AK to swap cases
func swapC(r rune) rune  {
	switch{
	case 'a' <= r && r <= 'z':
		return r - 'a' + 'A'
    case 'A' <= r && r <= 'Z':
        return r - 'A' + 'a'
    default:
        return r
	}
}


//Copied from lab1/gointro
type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	ind, errs := r.r.Read(p) 
	if errs != nil {
		return 0, errs
	}

	for i := 0; i < ind; i++ {
		val := p[i]
		//ascii value of alphabetical signs 65-90 for capital and 97-122 for lower case
		if val >= 65 && val <= 90 {
			val += 13
			if val > 90 {
				val -= 26
			}
		}else if val >= 97  && val <= 122{
			val += 13
			if val > 122 {
				val -= 26
			}
		}
		p[i] = val
	}
	return ind,nil
}

// socketIsClosed is a helper method to check if a listening socket has been
// closed.
func socketIsClosed(err error) bool {
	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}
	return false
}
