package balance

type  Balancer interface {
	 init(NodeAddrs)
     DoBalance(NodeAddrs) (*NodeAddr,int,error)
}
