package main

import (
	"flag"
	"go_dev/day14/cache-server/server/cache"
	"go_dev/day14/cache-server/server/cluster"
	"go_dev/day14/cache-server/server/http"
	"go_dev/day14/cache-server/server/tcp"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	ttl := flag.Int("ttl", 30, "cache time to live")
	node := flag.String("node", "127.0.0.1", "node address")
	clus := flag.String("cluster", "", "cluster address")
	flag.Parse()
	log.Println("type is", *typ)
	log.Println("ttl is", *ttl)
	log.Println("node is", *node)
	log.Println("cluster is", *clus)
	c := cache.New(*typ, *ttl)
	n, e := cluster.New(*node, *clus)
	if e != nil {
		panic(e)
	}
	go tcp.New(c, n).Listen()
	http.New(c, n).Listen()

}
