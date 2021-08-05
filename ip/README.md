
```go
package main

import (
	"github.com/dianjiu/gotool/ip"
	"log"
)

func main() {
	ip := ip.GetLocalIPv4()
	log.Printf("本机IP: %v\n", ip)
}
```