package main

import (
	"fmt"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/startup"
)

const watcher = `
   __
  /  \
 |    |
  \__/
Watcher is active
`

func main() {
	server := startup.Server{}
	server.StartUp()
	server.HandleSignal()
	fmt.Print(watcher)

	select {}
}
