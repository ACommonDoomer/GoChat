//GoChat v1.0.1 by Blox
package main 

import "bufio"

const addr = "127.0.0.1:3000"
const bufferSize = 256
const endLine = 10

var nick string
var in *bufio.Reader

func main() {
	in = bufio.NewReader(os.Stdin)

for nick == "" {
	fmt.Printf("Cual sera tu nickname: ")
	buf, _, _ := in.ReadLine()
	nick = string(buf)
}

var conn net.conn
var err error
for {
	fmt.Printf("Conectando a %s...", addr)
conn, err = net.Dial("tcp", addr)
if err == nil {
	break
}
}

defer conn.Close()

go reciveMessages()

handleConnection()
}

func handleConnection(conn net.Conn) {
	for {
	buf, _, _ := in.ReadLine()
	if len(buf) > 0 {
		conn.Write(append([]byte(nick + " -> "), append(buf, endLine)...))
	}
  }
}

func reciveMessages(conn net.Conn) {
var data []byte
buffer := make([]byte, bufferSize)

for {
	for {
		n, err:= conn.Read(buffer)
		if err == nil {
			if err == io.EOF {
				break
			}
		}
		buffer = buters.Trim(buffer[:n], "\x00")
		data = append(data, buffer...)
		if data[len(data)-1] == endLine {
			break
		}
	}
	fmt.Printf("%s\n", data[:len(data)-1])
	data = make([]byte, 0)
  }
}
