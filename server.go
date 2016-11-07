package whodb

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/brunohenrique/whodb/storage"
)

type Server struct {
	host    string
	port    int
	ln      net.Listener
	storage storage.Storage
}

func NewServer(host string, port int) *Server {
	s := storage.New()
	return &Server{host: host, port: port, storage: s}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		log.Fatalf("Unable to connect. Error: %v", err)
	}

	s.ln = ln
	go func() {
		for {
			conn, err := s.ln.Accept()
			if err != nil {
				continue
			}

			go s.handleClient(conn)
		}
	}()
	return nil
}

func (s *Server) Close() {
	s.ln.Close()
}

func (s *Server) handleClient(conn net.Conn) {
	for {
		msg := make([]byte, 1024)
		_, err := conn.Read(msg)
		if err != nil {
			break
		}

		tokens, _ := parseCommand(msg)
		cmd := tokens[2]
		if cmd == "set" {
			key := tokens[4]
			value := tokens[6]
			s.storage.Set(key, value)
			conn.Write([]byte("+OK\r\n"))
		} else if cmd == "get" {
			key := tokens[4]
			value := s.storage.Get(key)
			conn.Write([]byte("+" + value + "\r\n"))
		} else {
			conn.Write([]byte("+PONG\r\n"))
		}
	}
}

func parseCommand(msg []byte) ([]string, error) {
	tokens, err := tokenizeMsg(string(msg))
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func tokenizeMsg(msg string) ([]string, error) {
	msg = strings.TrimRight(msg, "\r\n")

	return strings.Split(msg, "\r\n"), nil
}
