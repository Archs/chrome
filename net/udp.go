package net

import (
	"fmt"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/udp"
	"net"
	"time"
)

var (
	// scoket map
	udpMap = make(map[int]*udpConn)
	// listen chan
	listenerMap = make(map[int]*udpListener)
)

type udpPacket struct {
	data          []byte
	remoteAddress net.Addr
}

type udpConn struct {
	socketId     int
	ch           chan *udpPacket
	readDeadline time.Time
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
	var p *udpPacket
	select {
	case p = <-u.ch:
		copy(b, p.data)
		return len(p.data), p.remoteAddress, nil
	default:
		if !u.readDeadline.IsZero() && time.Now().After(u.readDeadline) {
			return 0, nil, &timeoutError{
				OpError: net.OpError{
					Op:  "udp readfrom",
					Err: errors.New("timeout"),
				},
				isTimeOut: true,
			}
		}
		time.Sleep(time.Microsecond * 10)
	}

}

// Write writes data to the connection.
// Write can be made to time out and return a Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *udpConn) Write(b []byte) (n int, err error) {
	ch := make(chan struct{})
	udp.Send(c.socketId, b, func(si *udp.SendInfo) {
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
		addr = &net.udpAddr{
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
	var addr net.Addr
	ch := make(chan struct{})
	udp.GetInfo(c.socketId, func(si *sockets.SocketInfo) {
		addr = &net.udpAddr{
			IP:   net.ParseIP(si.PeerAddress),
			Port: si.PeerPort,
		}
		close(ch)
	})
	<-ch
	return addr
}

func init() {
	udp.OnReceive(func(ri *udp.ReceiveInfo) {
		go func() {
			// println("udp receive on socket:", ri.SocketId)
			// println("udpMap length:", len(udpMap))
			conn, ok := udpMap[ri.SocketId]
			if ok {
				// println("ri.Data", ri.Data)
				conn.readBuf.Write(ri.Data)
			} else {
				println("no conn found", udpMap)
			}
		}()
	})
	udp.OnReceiveError(func(re *udp.ReceiveError) {
		conn, ok := udpMap[re.SocketId]
		if ok {
			conn.readError = fmt.Errorf("recv error code %d", re.ResultCode)
		}
	})
	udpserver.OnAccept(func(ai *udpserver.AcceptInfo) {
		// println("udpserver accept:", ai.ClientSocketId)
		cl, ok := listenerMap[ai.SocketId]
		if !ok {
			return
		}
		udp.SetPaused(ai.ClientSocketId, false)
		cl.ch <- ai.ClientSocketId
	})
	udpserver.OnAcceptError(func(ae *udpserver.AcceptError) {
		cl, ok := listenerMap[ae.SocketId]
		if !ok {
			return
		}
		cl.err = fmt.Errorf("accept error: %d", ae.ResultCode)
	})
}
