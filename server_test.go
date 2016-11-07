package whodb_test

import (
	"fmt"
	"testing"

	"github.com/brunohenrique/whodb"

	"gopkg.in/redis.v5"
)

const PORT = 6378

func startServer() *whodb.Server {
	s := whodb.NewServer("localhost", PORT)
	ready := make(chan error)
	go func() {
		ready <- s.Start()
	}()
	<-ready
	return s
}

func TestClientPingingTheServer(t *testing.T) {
	s := startServer()
	defer s.Close()

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("localhost:%v", PORT),
	})
	defer client.Close()

	pong, _ := client.Ping().Result()
	if pong != "PONG" {
		t.Errorf("must return PONG")
	}
}

func TestClientSetAndGetValueFromTheServer(t *testing.T) {
	s := startServer()
	defer s.Close()

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("localhost:%v", PORT),
	})
	defer client.Close()

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		t.Errorf("error to set a value")
	}

	value, _ := client.Get("key").Result()
	if value != "value" {
		t.Errorf("must return 'value'")
	}
}
