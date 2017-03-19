package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//处理连接
func handleConnection(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {

		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

		// 		conn.Write([]byte(`
		// 		HTTP/1.1 200 OK
		// Server: nginx
		// Date: Fri, 17 Mar 2017 14:17:32 GMT
		// Content-Type: application/json;charset=UTF-8
		// Transfer-Encoding: chunked
		// Connection: close
		// \n

		// {"code":404,"msg":"欢迎使用23mofang API服务,请参见相关文档了解本API相关信息!"}
		// 		`))
		bytes := []byte(`{"code":404,"msg":"欢迎使用23mofang API服务,请参见相关文档了解本API相关信息!"}
`)

		conn.Write([]byte(`HTTP/1.1 200 OK
Server: nginx
Date: Fri, 17 Mar 2017 14:17:32 GMT
Content-Type: application/json;charset=UTF-8
Transfer-Encoding: chunked
Connection: close
`))
		//conn.Write([]byte(fmt.Sprintf("Content-Length:%d", len(bytes))))

		conn.Write([]byte(`
`))

		conn.Write(bytes)
		conn.Write([]byte{30, 13, 10})

		Log("write success ", string([]byte{13, 10, 13, 10}), "--", string([]byte{30}))

		conn.Close()
		break
	}

}
func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {

	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}
}
