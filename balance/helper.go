package balance

import (
	"fmt"
	"math/rand"
	"time"
)

func GenNodeAddrs(num int,weight int) NodeAddrs{
	var addrs NodeAddrs
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := NewNodeAddr(host, 8080,rand.Intn(weight)+1)
		addrs = append(addrs, one)
	}
	return addrs
}