package main

import (
	"fmt"
	"github.com/Archs/chrome/net"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/tcp"
	"github.com/Archs/chrome/net/sockets/tcpserver"
	_ "github.com/Archs/js/koSecureBindings"
	QUnit "github.com/fabioberger/qunit"
	"github.com/gopherjs/gopherjs/js"
	"github.com/mibitzi/gopherjs-ko"
	"log"
)

var (
	address    = "127.0.0.1"
	port       = 8088
	serverSock = 0
	clientSock = 0

	in  = ko.NewObservable("sending to server")
	out = ko.NewObservable("server response")

	conn net.Conn
)

func appendToOut(msg string) {
	old := out.Get().String()
	out.Set(old + "\n" + msg)
}

func applyBindings() {
	model := js.M{
		"input":  in,
		"output": out,
		"connect": func() {
			tcp.Connect(clientSock, address, port, func(result int) {
				if result < 0 {
					appendToOut("client connect failed")
				} else {
					appendToOut("client connect successfully")
				}
			})
		},
		"send": func() {
			dat := []byte(in.Get().String())
			tcp.Send(clientSock, dat, func(si *tcp.SendInfo) {
				if si.ResultCode < 0 {
					appendToOut("send failed")
				} else {
					appendToOut(fmt.Sprintf("send %d bytes", si.BytesSent))
				}
			})
		},
		"dial": func() {
			go func() {
				var err error
				conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
				if err != nil {
					appendToOut(err.Error())
				} else {
					appendToOut("chrome conn created")
				}
			}()
		},
		"write": func() {
			go func() {
				dat := []byte(in.Get().String())
				n, err := conn.Write(dat)
				if err != nil {
					appendToOut(err.Error())
				} else {
					appendToOut(fmt.Sprintf("%d bytes write to server", n))
				}
			}()
		},
	}
	ko.ApplyBindings(model)
}

func main() {
	applyBindings()
	QUnit.Module("sockets")
	tcpserver.Create(func(ci *sockets.CreateInfo) {
		serverSock = ci.SocketId
		QUnit.Test("tcpserver.Create", func(assert QUnit.QUnitAssert) {
			assert.NotEqual(serverSock, 0, "tcpserver.Create")
		})
		tcpserver.Listen(serverSock, "127.0.0.1", port, func(result int) {
			QUnit.Test("tcpserver.Listen", func(assert QUnit.QUnitAssert) {
				assert.Ok(result >= 0, "listen failed")
			})
		})
		tcpserver.OnAccept(func(ai *tcpserver.AcceptInfo) {
			log.Println("new client:", ai.ClientSocketId)
			tcp.SetPaused(ai.ClientSocketId, false)
		})
	})

	// QUnit.Module("sockets.tcp")
	tcp.CreateEx(func(ci *sockets.CreateInfo) {
		QUnit.Test("tcp.CreateEx", func(assert QUnit.QUnitAssert) {
			clientSock = ci.SocketId
			println(clientSock)
			assert.NotEqual(ci.SocketId, 0, "CreateEx")
		})
	})

	tcp.OnReceive(func(ri *tcp.ReceiveInfo) {
		println("receiving from:", ri.SocketId, string(ri.Data))
		appendToOut("server receive:" + string(ri.Data))
	})

	tcp.OnReceiveError(func(re *tcp.ReceiveError) {
		log.Printf("tcp receive error: %v\n", re)
		appendToOut(fmt.Sprintf("tcp receive error: %v\n", re))
	})
}
