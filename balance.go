package balance

type  Balancer interface {
	 init(NodeAddrs)
     DoBalance() (int,error)
}
