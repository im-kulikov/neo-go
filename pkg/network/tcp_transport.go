package network

import (
	"net"
	"regexp"
	"time"

	"go.uber.org/zap"
)

// TCPTransport allows network communication over TCP.
type TCPTransport struct {
	log      *zap.Logger
	server   *Server
	listener net.Listener
	bindAddr string
}

var reClosedNetwork = regexp.MustCompile(".* use of closed network connection")

// NewTCPTransport returns a new TCPTransport that will listen for
// new incoming peer connections.
func NewTCPTransport(s *Server, bindAddr string, log *zap.Logger) *TCPTransport {
	return &TCPTransport{
		log:      log,
		server:   s,
		bindAddr: bindAddr,
	}
}

// Dial implements the Transporter interface.
func (t *TCPTransport) Dial(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return err
	}
	p := NewTCPPeer(conn, t.server)
	go p.handleConn()
	return nil
}

// Accept implements the Transporter interface.
func (t *TCPTransport) Accept() {
	l, err := net.Listen("tcp", t.bindAddr)
	if err != nil {
		t.log.Panic("TCP listen error", zap.Error(err))
		return
	}

	t.listener = l

	for {
		conn, err := l.Accept()
		if err != nil {
			t.log.Warn("TCP accept error", zap.Error(err))
			if t.isCloseError(err) {
				break
			}
			continue
		}
		p := NewTCPPeer(conn, t.server)
		go p.handleConn()
	}
}

func (t *TCPTransport) isCloseError(err error) bool {
	if opErr, ok := err.(*net.OpError); ok {
		if reClosedNetwork.Match([]byte(opErr.Error())) {
			return true
		}
	}

	return false
}

// Close implements the Transporter interface.
func (t *TCPTransport) Close() {
	t.listener.Close()
}

// Proto implements the Transporter interface.
func (t *TCPTransport) Proto() string {
	return "tcp"
}
