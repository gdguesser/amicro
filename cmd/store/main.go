package main

import (
	"flag"
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
	"github.com/gdguesser/micro/types"
)

type store struct{}

func newStore() actor.Receiver {
	return &store{}
}

func (s *store) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case *types.CatFact:
		fmt.Println("stored fact into the db: ", msg.Fact)
	case actor.Started:
		fmt.Println("store is started")
	}
}

func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:4000", "todo")
	flag.Parse()

	e := actor.NewEngine()
	r := remote.New(e, remote.Config{
		ListenAddr: *listenAddr,
	})
	e.WithRemote(r)

	// pid => 127.0.0.1:4000/store
	e.Spawn(newStore, "store")

	select {}
}
