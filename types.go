package balance

import "strconv"

type NodeAddrs []*NodeAddr

type NodeAddr struct{
	host string
	port int
	weight int
}

func NewNodeAddr(host string,port int,weight int) *NodeAddr{
	return &NodeAddr{
		host:host,
		port:port,
		weight:weight,
	}
}

func (p *NodeAddr) GetHost() string {
	return p.host
}

func (p *NodeAddr) GetPort() int {
	return p.port
}

func (p *NodeAddr) GetWeight() int {
	return p.weight
}

func (p *NodeAddr) UpdateWeight( weight int) {
	p.weight = weight
}

func (p *NodeAddr) String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}

func (p *NodeAddr) StringDebug() string {
	return p.host + ":" + strconv.Itoa(p.port)+":" + strconv.Itoa(p.weight)
}
