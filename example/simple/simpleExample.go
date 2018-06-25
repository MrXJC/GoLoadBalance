package simple

import(
	"fmt"
	. "github.com/MrXJC/GoLoadBalance/balance"
)

func main(){
	addrs := GenNodeAddrs(3,10)
	for _,addr :=range addrs {
		fmt.Println(addr)
	}

	mgr := NewBalanceMgr(addrs)
	mgr.RegisterBalancer("random",&RandomBalance{})
	mgr.RegisterBalancer("randomweight",&RandomWeightBalance{})

	note := map[string]int{}
	for i := 0; i < 10000; i++ {
		s,_,_ := mgr.GetAddrStringDebug("randomweight")
		//fmt.Println(s)
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
	mgr.DeleteNodeAddr(0)
	for i := 0; i < 10000; i++ {
		s,_,_ := mgr.GetAddrStringDebug("randomweight")
		//fmt.Println(s)
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