package net

import (
	"errors"
	"fmt"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/udp"
	"net"
	"time"
)

var (
	// scoket map
	udpMap = make(map[int]*udpConn)
)

type udpPacket struct {
	data       []byte
	remoteAddr net.Addr
	err        error
}

func newPacket(data []byte, addr net.Addr) *udpPacket {
	return &udpPacket{
		data:       data,
		remoteAddr: addr,
		err:        nil,
	}
}

func newPacketErr(err error) *udpPacket {
	return &udpPacket{
		err: err,
	}
}

type udpConn struct {
	socketId int
	// data buffer
	ch           chan *udpPacket
	readDeadline time.Time
	// addr
	laddr *net.UDPAddr
	raddr *net.UDPAddr
}

func newUdpConn(socketId int) *udpConn {
	// println("creating conn, id:", socketId)
	conn := &udpConn{
		socketId: socketId,
		ch:       make(chan *udpPacket, 10),
	}
	udpMap[socketId] = conn
	// println("c udpMap length:", len(udpMap))
	return conn
}

func (u *udpConn) ReadFrom(b []byte) (int, net.Addr, error) {
	for {
		select {
		case p := <-u.ch:
			copy(b, p.data)
			return len(p.data), p.remoteAddr, nil
		case <-time.Tick(10 * time.Microsecond):
			if !u.readDeadline.IsZero() && time.Now().After(u.readDeadline) {
				return 0, nil, &timeoutError{
					OpError: net.OpError{
						Op:  "udp readfrom",
						Err: errors.New("timeout"),
					},
					isTimeOut: true,
				}
			}
		}
	}
}

func (c *udpConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	ch := make(chan struct{})
	udp.Send(c.socketId, b, addr.IP.String(), addr.Port, func(si *udp.SendInfo) {
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

// Write writes data to the connection.
// Write can be made to time out and return a Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *udpConn) Write(b []byte) (n int, err error) {
	return c.WriteTo(b, c.raddr)
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c *udpConn) Close() error {
	// println("conn.Close called")
	udp.Close(c.socketId)
	delete(udpMap, c.socketId)
	return nil
}

// LocalAddr returns the local network address.
func (c *udpConn) LocalAddr() net.Addr {
	var addr net.Addr
	ch := make(chan struct{})
	udp.GetInfo(c.socketId, func(si *sockets.SocketInfo) {
		addr = &net.UDPAddr{
			IP:   net.ParseIP(si.LocalAddress),
			Port: si.LocalPort,
		}
		close(ch)
	})
	<-ch
	return addr
}

// RemoteAddr returns the remote network address.
func (c *udpConn) RemoteAddr() net.Addr {
	return nil
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
func (c *udpConn) SetDeadline(t time.Time) error {
	c.SetReadDeadline(t)
	c.SetWriteDeadline(t)
	return nil
}

// SetReadDeadline sets the deadline for future Read calls.
// A zero value for t means Read will not time out.
func (c *udpConn) SetReadDeadline(t time.Time) error {
	c.readDeadline = t
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *udpConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func init() {
	udp.OnReceive(func(ri *udp.ReceiveInfo) {
		go func() {
			// println("udp receive on socket:", ri.SocketId)
			// println("udpMap length:", len(udpMap))
			conn, ok := udpMap[ri.SocketId]
			if ok {
				addr, err := net.ResolveUDPAddr("udp", ri.RemoteAddress)
				if err != nil {
					conn.ch <- newPacket(ri.Data, addr)
				} else {
					conn.ch <- newPacketErr(err)
				}
			} else {
				println("no conn found", udpMap)
			}
		}()
	})
	udp.OnReceiveError(func(re *udp.ReceiveError) {
		conn, ok := udpMap[re.SocketId]
		if ok {
			conn.ch <- newPacketErr(fmt.Errorf("recv error code %d", re.ResultCode))
		}
	})
}
