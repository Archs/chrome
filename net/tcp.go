package net

import (
	"fmt"
	"github.com/Archs/chrome/net/sockets"
	"github.com/Archs/chrome/net/sockets/tcp"
	"github.com/Archs/chrome/net/sockets/tcpserver"
	"net"
)

var (
	// scoket map
	tcpMap = make(map[int]*tcpConn)
	// listen chan
	listenerMap = make(map[int]*tcpListener)
)

type tcpConn struct {
	*buffer
	socketId int
}

func newTcpConn(socketId int) *tcpConn {
	// println("creating conn, id:", socketId)
	conn := &tcpConn{
		socketId: socketId,
		buffer:   newBuffer(),
	}
	tcpMap[socketId] = conn
	// println("c tcpMap length:", len(tcpMap))
	return conn
}

type tcpListener struct {
	socketId int
	ch       chan int
	err      error
}

// Accept waits for and returns the next connection to the listener.
func (cl *tcpListener) Accept() (c net.Conn, err error) {
	if cl.err != nil {
		c = nil
		err = cl.err
		cl.err = nil
		return
	}
	id := <-cl.ch
	// println("listener.Accept socket id", id)
	c = newTcpConn(id)
	return
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (c *tcpListener) Close() error {
	tcpserver.Close(c.socketId)
	delete(listenerMap, c.socketId)
	return nil
}

// Addr returns the listener's network address.
func (c *tcpListener) Addr() net.Addr {
	var addr net.Addr
	ch := make(chan struct{})
	tcpserver.GetInfo(c.socketId, func(si *sockets.SocketInfo) {
		addr = &net.TCPAddr{
			IP:   net.ParseIP(si.LocalAddress),
			Port: si.LocalPort,
		}
		close(ch)
	})
	<-ch
	return addr
}

// Write writes data to the connection.
// Write can be made to time out and return a Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *tcpConn) Write(b []byte) (n int, err error) {
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
func (c *tcpConn) Close() error {
	// println("conn.Close called")
	tcp.Close(c.socketId)
	delete(tcpMap, c.socketId)
	return nil
}

// LocalAddr returns the local network address.
func (c *tcpConn) LocalAddr() net.Addr {
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
func (c *tcpConn) RemoteAddr() net.Addr {
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

func init() {
	tcp.OnReceive(func(ri *tcp.ReceiveInfo) {
		go func() {
			// println("tcp receive on socket:", ri.SocketId)
			// println("tcpMap length:", len(tcpMap))
			conn, ok := tcpMap[ri.SocketId]
			if ok {
				// println("ri.Data", ri.Data)
				conn.readBuf.Write(ri.Data)
			} else {
				println("no conn found", tcpMap)
			}
		}()
	})
	tcp.OnReceiveError(func(re *tcp.ReceiveError) {
		conn, ok := tcpMap[re.SocketId]
		if ok {
			conn.readError = fmt.Errorf("recv error code %d", re.ResultCode)
		}
	})
	tcpserver.OnAccept(func(ai *tcpserver.AcceptInfo) {
		// println("tcpserver accept:", ai.ClientSocketId)
		cl, ok := listenerMap[ai.SocketId]
		if !ok {
			return
		}
		tcp.SetPaused(ai.ClientSocketId, false)
		cl.ch <- ai.ClientSocketId
	})
	tcpserver.OnAcceptError(func(ae *tcpserver.AcceptError) {
		cl, ok := listenerMap[ae.SocketId]
		if !ok {
			return
		}
		cl.err = fmt.Errorf("accept error: %d", ae.ResultCode)
	})
}
