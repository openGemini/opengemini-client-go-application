package main

import (
	"fmt"
	_ "github.com/libgox/buffer"
	_ "github.com/libgox/flyway"
	_ "github.com/libgox/gocollections/listx"
	_ "github.com/libgox/spring-cloud-go"
	_ "github.com/protocol-laboratory/opcua-go/opcua"
)

func main() {
	fmt.Println("Hello, World!")
}
