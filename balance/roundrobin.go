package balance

import (
	"errors"
)

type RoundRobinBalance struct {
	size int
	curindex int
}

func (p *RoundRobinBalance) init(addrs NodeAddrs){
	p.size = len(addrs)
	p.curindex = -1
}

func (p *RoundRobinBalance) DoBalance(addrs NodeAddrs)(*NodeAddr,int,error){
	if p.size == 0 {
		err := errors.New("No addr")
		return &NodeAddr{},0,err
	}


	p.curindex = (p.curindex + 1) % p.size

	if p.curindex >= p.size {
		p.curindex = 0
	}

	return addrs[p.curindex],p.curindex,nil
}



type RoundRobinWeightBalance struct {
	size int
	curindex int
	maxweight int
	gcd int
	cw int
}

func (p *RoundRobinWeightBalance) init(addrs NodeAddrs){
	p.size = len(addrs)
	p.curindex = -1
	p.maxweight = getMaxWeight(addrs)
	p.gcd = Gcd(addrs)
	p.cw = 0
}

func (p *RoundRobinWeightBalance) DoBalance(addrs NodeAddrs)(*NodeAddr,int,error){
	if p.size == 0 {
		err := errors.New("No addr")
		return &NodeAddr{},0,err
	}

	for {
		p.curindex = (p.curindex + 1) % p.size
		if p.curindex == 0 {
			p.cw = p.cw - p.gcd
			if p.cw <= 0 {
				p.cw = p.maxweight
				if p.cw == 0 {
					err := errors.New("the max weight is 0")
					return &NodeAddr{},0,err
				}
			}
		}

		if weight:= addrs[p.curindex].GetWeight(); weight >= p.cw {
	        return addrs[p.curindex],p.curindex,nil
		}
	}


}



func getMaxWeight(addrs NodeAddrs) int{
	max:=0
	for _,addr := range addrs{
		if addr.GetWeight()>max{
			max = addr.GetWeight()
		}
	}
	return max
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}
func Gcd(addrs NodeAddrs) int {
	var c int
	w := addrs[0].GetWeight()
	for _,addr := range addrs{
		c = addr.GetWeight()
		if c!=0{
			w = gcd(w,c)
		}
	}
	return  w
}
