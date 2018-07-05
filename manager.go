package balance

import "fmt"

type BalanceMgr struct{
	addrs NodeAddrs
	balancers map[string]Balancer
	change map[string]bool
}

func NewBalanceMgr(addrs NodeAddrs) *BalanceMgr {
	return &BalanceMgr{
		balancers:make(map[string]Balancer),
		addrs:addrs,
		change:make(map[string]bool),
	}
}

func (p *BalanceMgr) RegisterBalancer(name string,balancer Balancer){
	    balancer.init(p.addrs)
    	p.balancers[name]=balancer
    	p.change[name]=false
}
func (p *BalanceMgr) updateBalancer(name string){
	 fmt.Println("update ",name)
     balancer := p.balancers[name]
     balancer.init(p.addrs)
	 p.balancers[name]=balancer
	 p.change[name]=false
}

func (p *BalanceMgr) GetAddr(name string)(*NodeAddr,int,error) {
	if p.change[name]{
		p.updateBalancer(name)
	}
	balancer := p.balancers[name]
	addr,index,err  := balancer.DoBalance(p.addrs)
	return addr,index,err
}

func (p *BalanceMgr) GetAddrString(name string)(string,int,error) {
	if p.change[name]{
		p.updateBalancer(name)
	}

	addr,index,err := p.GetAddr(name)
	return addr.String(),index,err
}

func (p *BalanceMgr) GetAddrStringDebug(name string)(string,int,error) {
	if p.change[name]{
		p.updateBalancer(name)
	}

	addr,index,err := p.GetAddr(name)
	return addr.StringDebug(),index,err
}

func (p *BalanceMgr) AddNodeAddr(addr *NodeAddr){
	p.addrs = append(p.addrs,addr)
	for name,_ :=range p.change{
		p.change[name]=true
	}
}

func (p *BalanceMgr) DeleteNodeAddr(i int){
	p.addrs = append(p.addrs[:i], p.addrs[i+1:]...)
	for name,_ :=range p.change{
		p.change[name]=true
	}
}


func (p *BalanceMgr) UpdateWeightNodeAddr(i int,weight int){
	p.addrs[i].UpdateWeight(weight)
	for name,_ :=range p.change{
		p.change[name]=true
	}
}