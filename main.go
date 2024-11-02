package main

import (
	"fmt"
	_ "github.com/libgox/addr"
	_ "github.com/libgox/asciitable"
	_ "github.com/libgox/buffer"
	_ "github.com/libgox/envx"
	_ "github.com/libgox/flyway"
	_ "github.com/libgox/gocollections/listx"
	_ "github.com/libgox/must"
	_ "github.com/libgox/properties"
	_ "github.com/libgox/retry"
	_ "github.com/libgox/slogsimple"
	_ "github.com/libgox/spring-cloud-go"
	_ "github.com/libgox/unicodex/letter"
	_ "github.com/protocol-laboratory/opcua-go/opcua"
	_ "github.com/protocol-laboratory/zookeeper-client-go/zk"
)

func main() {
	fmt.Println("Hello, World!")
}
