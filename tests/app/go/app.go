package main

import (
	"fmt"
	"github.com/Archs/chrome/api/sockets"
	"github.com/Archs/chrome/api/sockets/tcp"
	"github.com/Archs/chrome/api/sockets/tcpserver"
	"github.com/Archs/chrome/sim/net"
	"github.com/Archs/gopherjs-ko"
	QUnit "github.com/fabioberger/qunit"
	"github.com/gopherjs/gopherjs/js"
	"log"
)

var (
	serverSock = 0
	clientSock = 0

	ip   = ko.NewObservable("127.0.0.1")
	port = ko.NewObservable(8088)
	in   = ko.NewObservable("sending to server")
	out  = ko.NewObservable("server response")

	currentConn net.Conn
)

func appendToOut(msg string) {
	old := out.Get().String()
	out.Set(old + "\n" + msg)
	println(msg)
}

func applyBindings() {
	model := js.M{
		"ip":     ip,
		"port":   port,
		"input":  in,
		"output": out,
		"connect": func() {
			tcp.Connect(clientSock, ip.Get().String(), port.Get().Int(), func(result int) {
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
		"listen": func() {
			go func() {
				port.Set(port.Get().Int() + 1)
				l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip.Get().String(), port.Get().Int()))
				if err != nil {
					appendToOut("listen failed:" + err.Error())
					return
				}
				appendToOut(fmt.Sprintf("listen on port: %d", port.Get().Int()))
				for {
					c, err := l.Accept()
					if err != nil {
						appendToOut("accetp error:" + err.Error())
						continue
					}
					appendToOut("new conn comming:" + c.RemoteAddr().String())
					go func(c net.Conn) {
						for {
							buf := make([]byte, 1024)
							_, err := c.Read(buf)
							if err != nil {
								appendToOut("conn receiv err:" + err.Error())
							} else {
								appendToOut("conn receiv:" + string(buf))
							}
						}
					}(c)
				}
			}()
		},
		"dial": func() {
			go func() {
				appendToOut(fmt.Sprintf("dialing %s:%d", ip.Get().String(), port.Get().Int()))
				conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip.Get().String(), port.Get().Int()))
				if err != nil {
					appendToOut(err.Error())
				} else {
					currentConn = conn
					appendToOut("chrome conn created:" + conn.LocalAddr().String())
				}
			}()
		},
		"write": func() {
			go func() {
				if currentConn == nil {
					return
				}
				dat := []byte(in.Get().String())
				appendToOut("writing:" + currentConn.LocalAddr().String())
				n, err := currentConn.Write(dat)
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
	ko.EnableSecureBinding()
	applyBindings()
	QUnit.Module("sockets")
	tcpserver.Create(func(ci *sockets.CreateInfo) {
		serverSock = ci.SocketId
		QUnit.Test("tcpserver.Create", func(assert QUnit.QUnitAssert) {
			assert.NotEqual(serverSock, 0, "tcpserver.Create")
		})
		tcpserver.Listen(serverSock, ip.Get().String(), port.Get().Int(), func(result int) {
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
}
