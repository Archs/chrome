package net

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/tcp"
	"net"
	"sync"
	"time"
)

var (
	// scoket map
	smap = make(map[int]*crConn)
	m    = new(sync.Mutex)
)

type Conn net.Conn

// chrome conn
type crConn struct {
	socketId     int
	readBuf      *bytes.Buffer
	readError    error
	readDeadLine time.Time
}

func Dial(network, address string) (Conn, error) {
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
	return &crConn{
		socketId: socketId,
	}, nil
}

type timeoutError struct {
	net.OpError
	isTimeOut bool
}

func (t *timeoutError) Timeout() bool {
	return t.isTimeOut
}

// Read reads data from the connection.
// Read can be made to time out and return a Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *crConn) Read(b []byte) (n int, err error) {
	m.Lock()
	defer m.Unlock()
	if c.readError != nil {
		return 0, c.readError
	}
	n, err = c.readBuf.Read(b)
	if n == 0 && !c.readDeadLine.IsZero() && time.Now().Before(c.readDeadLine) {
		time.Sleep(time.Millisecond * 10)
		return c.Read(b)
	}
	if n == 0 && !c.readDeadLine.IsZero() && time.Now().After(c.readDeadLine) {
		return 0, &timeoutError{
			OpError: net.OpError{
				Op:   "read",
				Net:  "tcp",
				Addr: c.RemoteAddr(),
				Err:  errors.New("timeout"),
			},
			isTimeOut: true,
		}
	}
	return n, err
}

// Write writes data to the connection.
// Write can be made to time out and return a Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *crConn) Write(b []byte) (n int, err error) {
	ch := make(chan struct{})
	tcp.Send(c.socketId, b, func(si *tcp.SendInfo) {
		if si.ResultCode < 0 {
			err = fmt.Errorf("socket write failed: %d", si.ResultCode)
		} else {
			n = si.BytesSent
		}
		close(ch)
	})
	<-ch
	return
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c *crConn) Close() error {
	tcp.Close(c.socketId)
	return nil
}

// LocalAddr returns the local network address.
func (c *crConn) LocalAddr() net.Addr {
	var addr net.Addr
	ch := make(chan struct{})
	tcp.GetInfo(c.socketId, func(si *sockets.SocketInfo) {
		addr = &net.TCPAddr{
			IP:   net.ParseIP(si.LocalAddress),
			Port: si.LocalPort,
		}
		close(ch)
	})
	<-ch
	return addr
}

// RemoteAddr returns the remote network address.
func (c *crConn) RemoteAddr() net.Addr {
	var addr net.Addr
	ch := make(chan struct{})
	tcp.GetInfo(c.socketId, func(si *sockets.SocketInfo) {
		addr = &net.TCPAddr{
			IP:   net.ParseIP(si.PeerAddress),
			Port: si.PeerPort,
		}
		close(ch)
	})
	<-ch
	return addr
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
func (c *crConn) SetDeadline(t time.Time) error {
	c.SetReadDeadline(t)
	c.SetWriteDeadline(t)
	return nil
}

// SetReadDeadline sets the deadline for future Read calls.
// A zero value for t means Read will not time out.
func (c *crConn) SetReadDeadline(t time.Time) error {
	c.readDeadLine = t
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *crConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func init() {
	tcp.OnReceive(func(ri *tcp.ReceiveInfo) {
		m.Lock()
		defer m.Unlock()
		smap[ri.SocketId].readBuf.Write(ri.Data)
	})
	tcp.OnReceiveError(func(re *tcp.ReceiveError) {
		m.Lock()
		defer m.Unlock()
		smap[re.SocketId].readError = fmt.Errorf("recv error code %d", re.ResultCode)
	})
}
