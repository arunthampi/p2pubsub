package main

import (
	"context"
	"log"
	"net"

	"github.com/ipfs/go-ipfs/commands"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/corehttp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd, err := core.NewNode(ctx, &core.BuildCfg{})
	if err != nil {
		log.Fatal(err)
	}

	cctx := commands.Context{
		Online: true,
		ConstructNode: func() (*core.IpfsNode, error) {
			return nd, nil
		},
	}

	list, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on: ", list.Addr())

	if err := corehttp.Serve(nd, list, corehttp.CommandsOption(cctx)); err != nil {
		log.Fatal(err)
	}
}
