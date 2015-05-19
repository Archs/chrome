package net

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/tcp"
	"github.com/Archs/chrome/net/sockets/tcpserver"
	"net"
	"time"
)

type Conn net.Conn

// chrome conn
type buffer struct {
	readBuf      *bytes.Buffer
	readError    error
	readDeadLine time.Time
}

func newBuffer() *buffer {
	return &buffer{
		readBuf:      bytes.NewBuffer(nil),
		readError:    nil,
		readDeadLine: time.Time{},
	}
}

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

// Read reads data from the connection.
// Read can be made to time out and return a Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *buffer) Read(b []byte) (n int, err error) {
	if c.readError != nil {
		return 0, c.readError
	}
	for {
		m.Lock()
		n, err = c.readBuf.Read(b)
		m.Unlock()
		// data read
		if n > 0 {
			// in case of err == EOF
			err = nil
			return
		}
		// timeout
		if c.hasTimeout() {
			return 0, &timeoutError{
				OpError: net.OpError{
					Op:  "read",
					Err: errors.New("timeout"),
				},
				isTimeOut: true,
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future I/O, not just
// the immediately following call to Read or Write.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (c *buffer) SetDeadline(t time.Time) error {
	c.SetReadDeadline(t)
	c.SetWriteDeadline(t)
	return nil
}

// SetReadDeadline sets the deadline for future Read calls.
// A zero value for t means Read will not time out.
func (c *buffer) SetReadDeadline(t time.Time) error {
	c.readDeadLine = t
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *buffer) SetWriteDeadline(t time.Time) error {
	return nil
}
