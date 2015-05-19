package net

import (
	"errors"
	"fmt"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/udp"
	"net"
	"strings"
	"time"
)

var (
	// scoket map
	udpMap = make(map[int]*UDPConn)
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

type UDPConn struct {
	socketId int
	// data buffer
	ch           chan *udpPacket
	readDeadline time.Time
	// addr
	laddr *net.UDPAddr
	raddr *net.UDPAddr
}

func newUDPConn(socketId int) *UDPConn {
	// println("creating conn, id:", socketId)
	conn := &UDPConn{
		socketId:     socketId,
		ch:           make(chan *udpPacket, 10),
		readDeadline: time.Time{},
	}
	udpMap[socketId] = conn
	// println("c udpMap length:", len(udpMap))
	return conn
}

// DialUDP connects to the remote address raddr on the network net,
// which must be "udp", "udp4", or "udp6".
// If laddr is not nil, it is used as the local address for the connection.
func DialUDP(network string, laddr, raddr *net.UDPAddr) (*UDPConn, error) {
	if !strings.HasPrefix(network, "udp") || len(network) > 4 {
		return nil, errors.New("network not supported")
	}
	var err error
	var conn *UDPConn
	sig := make(chan struct{})
	socketId := 0
	udp.CreateEx(func(ci *sockets.CreateInfo) {
		socketId = ci.SocketId
		// default bind to local free port, by using 0
		if laddr == nil {
			laddr, err = net.ResolveUDPAddr(network, ":0")
		}
		if err != nil {
			close(sig)
			return
		}
		// bind to local address
		udp.Bind(socketId, laddr.IP.String(), laddr.Port, func(result int) {
			if result < 0 {
				err = fmt.Errorf("udp bind local address failed:%d", result)
			} else {
				conn = newUDPConn(socketId)
				conn.laddr = laddr
				conn.raddr = raddr
			}
			close(sig)
		})
	})
	<-sig
	return conn, err
}

// ListenMulticastUDP listens for incoming multicast UDP packets addressed to the group address gaddr on ifi, which specifies the interface to join. ListenMulticastUDP uses default multicast interface if ifi is nil.
// func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error) {
// 	reutrn nil, nil
// }

// ListenUDP listens for incoming UDP packets addressed to the local address laddr.
// Net must be "udp", "udp4", or "udp6".
// If laddr has a port of 0, ListenUDP will choose an available port.
// The LocalAddr method of the returned UDPConn can be used to discover the port.
// The returned connection's ReadFrom and WriteTo methods can be used to receive and send UDP packets with per-packet addressing.
func ListenUDP(net string, laddr *net.UDPAddr) (*UDPConn, error) {
	return DialUDP(net, laddr, nil)
}

func (u *UDPConn) ReadFrom(b []byte) (int, net.Addr, error) {
	for {
		select {
		case p := <-u.ch:
			n := len(b)
			if len(p.data) < n {
				n = len(p.data)
			}
			copy(b, p.data[:n])
			return n, p.remoteAddr, nil
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

// TODO add the ablity to read full of b, maybe using a double list
func (c *UDPConn) Read(b []byte) (int, error) {
	n, _, err := c.ReadFrom(b)
	return n, err
}

func (c *UDPConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	ch := make(chan struct{})
	n := 0
	raddr, err := net.ResolveUDPAddr(addr.Network(), addr.String())
	if err != nil {
		return 0, nil
	}
	udp.Send(c.socketId, b, raddr.IP.String(), raddr.Port, func(si *udp.SendInfo) {
		if si.ResultCode < 0 {
			err = fmt.Errorf("socket write failed: %d", si.ResultCode)
		} else {
			n = si.BytesSent
		}
		close(ch)
	})
	<-ch
	return n, nil
}

// Write writes data to the connection.
// Write can be made to time out and return a Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *UDPConn) Write(b []byte) (n int, err error) {
	return c.WriteTo(b, c.raddr)
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c *UDPConn) Close() error {
	// println("conn.Close called")
	udp.Close(c.socketId)
	delete(udpMap, c.socketId)
	return nil
}

// LocalAddr returns the local network address.
func (c *UDPConn) LocalAddr() net.Addr {
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
func (c *UDPConn) RemoteAddr() net.Addr {
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
func (c *UDPConn) SetDeadline(t time.Time) error {
	c.SetReadDeadline(t)
	c.SetWriteDeadline(t)
	return nil
}

// SetReadDeadline sets the deadline for future Read calls.
// A zero value for t means Read will not time out.
func (c *UDPConn) SetReadDeadline(t time.Time) error {
	c.readDeadline = t
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *UDPConn) SetWriteDeadline(t time.Time) error {
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
