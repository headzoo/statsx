package main

import (
	"fmt"
	"github.com/headzoo/statsx/storage"
)

func main() {
	s := storage.New(10)
	foo := storage.NewNode(storage.TYPE_DIRECTORY, []byte("foo"), 2)
	bar := storage.NewNode(storage.TYPE_DIRECTORY, []byte("bar"), 2)
	gaz := storage.NewNode(storage.TYPE_DIRECTORY, []byte("gaz"), 2)
	bar.Children.Append(gaz)
	s.Root.Children.Append(foo)
	s.Root.Children.Append(bar)
	
	s.Walk(storage.WalkerFunc(func(node *storage.Node, path []byte) {
		fmt.Printf("%s/%s\n", path, node.Name)
	}));
}
