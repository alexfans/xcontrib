package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"flag"
	"github.com/alexfans/xcontrib"
	"log"
)

func main() {
	var (
		schemaPath   = flag.String("path", "", "path to schema directory")
		schemaAssign = flag.String("assign", "", "schema name")
	)
	flag.Parse()
	if *schemaPath == "" {
		log.Fatal("xcontrib: must specify schema path. use xcontrib -path ./ent/schema [-assign name]")
	}
	graph, err := entc.LoadGraph(*schemaPath, &gen.Config{})
	if err != nil {
		log.Fatalf("xcontrib: failed loading ent graph: %v", err)
	}
	if err := xcontrib.Generate(graph, *schemaAssign); err != nil {
		log.Fatalf("xcontrib: failed generating protos: %s", err)
	}
}
