package net

import (
	"errors"
	"fmt"
	"github.com/Archs/chrome/api/sockets"
	"github.com/Archs/chrome/api/sockets/tcp"
	"github.com/Archs/chrome/api/sockets/tcpserver"
	"net"
	"strings"
	"time"
)

type Conn net.Conn

// Dial connects to the address on the named network.
//
// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only).
//
// For TCP and UDP networks, addresses have the form host:port. If host is a literal IPv6 address it must be enclosed in square brackets as in "[::1]:80" or "[ipv6-host%zone]:80". The functions JoinHostPort and SplitHostPort manipulate addresses in this form.
func Dial(network, address string) (net.Conn, error) {
	if len(network) > 4 || len(network) == 0 {
		return nil, errors.New("bad network")
	}
	if strings.HasPrefix(network, "udp") {
		raddr, err := net.ResolveUDPAddr(network, address)
		if err != nil {
			return nil, err
		}
		return DialUDP(network, nil, raddr)
	}
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
