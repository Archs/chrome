package net

import (
	"errors"
	"fmt"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/tcp"
	"github.com/Archs/chrome/net/sockets/tcpserver"
	"net"
	"time"
)

type Conn net.Conn

func Dial(network, address string) (net.Conn, error) {
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, err
	}
	socketId := 0
	ch := make(chan struct{})
	tcp.CreateEx(func(ci *sockets.CreateInfo) {
		socketId = ci.SocketId
		tcp.Connect(socketId, addr.IP.String(), addr.Port, func(result int) {
			if result < 0 {
				err = errors.New("connecting " + addr.String() + " error")
			}
			close(ch)
		})
	})
	<-ch
	if err != nil {
		return nil, err
	}
	conn := newTcpConn(socketId)
	return conn, nil
}

func Listen(network, laddr string) (net.Listener, error) {
	var err error
	var addr *net.TCPAddr
	socketId := 0
	ch := make(chan struct{})
	tcpserver.Create(func(ci *sockets.CreateInfo) {
		socketId = ci.SocketId
		addr, err = net.ResolveTCPAddr(network, laddr)
		if err == nil {
			tcpserver.Listen(socketId, addr.IP.String(), addr.Port, func(result int) {
				if result < 0 {
					err = fmt.Errorf("tcp server listen error: %d", result)
				}
			})
		}
		close(ch)
	})
	<-ch
	if err != nil {
		return nil, err
	}
	l := &tcpListener{
		socketId: socketId,
		ch:       make(chan int),
		err:      nil,
	}
	listenerMap[socketId] = l
	return l, nil
}

type timeoutError struct {
	net.OpError
	isTimeOut bool
}

func (t *timeoutError) Timeout() bool {
	return t.isTimeOut
}

func (c *buffer) hasTimeout() bool {
	if !c.readDeadLine.IsZero() && time.Now().After(c.readDeadLine) {
		return true
	}
	return false
}
