package balance

import (
	"testing"
	"fmt"
)

func Test_GetMaxWeight(t *testing.T){
	addrs := genNodeAddrs(3,10)
	fmt.Println(getMaxWeight(addrs))
    addrs[0].weight=4
    addrs[1].weight=4
    addrs[2].weight=8
    fmt.Println(Gcd(addrs))
}

func TestBalanceMgr_RoundRobin(t *testing.T) {
	addrs := genNodeAddrs(3,10)
	for _,addr :=range addrs {
		fmt.Println(addr)
	}

	mgr := NewBalanceMgr(addrs)
	mgr.RegisterBalancer("roundrobin",&RoundRobinBalance{})
	mgr.RegisterBalancer("roundrobinweight",&RoundRobinWeightBalance{})

	note := map[string]int{}
	for i := 0; i < 100; i++ {
		s,_,_ := mgr.GetAddrStringDebug("roundrobin")
		fmt.Println(s)
		if note[s] != 0 {
			note[s]++
		} else {
			note[s] = 1
		}
	}
	for k, v := range note {
		fmt.Println(k, " ", v)
	}

	note = map[string]int{}
	//mgr.DeleteNodeAddr(0)
	for i := 0; i < 100; i++ {
		s,_,_ := mgr.GetAddrStringDebug("roundrobinweight")
		fmt.Println(s)
		if note[s] != 0 {
			note[s]++
		} else {
			note[s] = 1
		}
	}
	for k, v := range note {
		fmt.Println(k, " ", v)
	}
}

