package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	size int
}

func (p *RandomBalance) init(addrs NodeAddrs){
	p.size = len(addrs)
}

func (p *RandomBalance) DoBalance()(int,error){
	if p.size == 0 {
		err := errors.New("No addr")
		return 0,err
	}
	curindex := rand.Intn(p.size)
	return curindex,nil
}

type RandomWeightBalance struct {
	size int
	distribution     []int
}

func (p *RandomWeightBalance) init(addrs NodeAddrs){
	p.distribution =[]int{}
	for index,addr := range addrs{
		for i:=0;i<addr.GetWeight();i++{
			p.distribution = append(p.distribution,index)
		}
	}
	p.size = len(p.distribution)
}

func (p *RandomWeightBalance) DoBalance()(int,error){
	if p.size == 0 {
		err := errors.New("No addr")
		return 0,err
	}
	curindex :=p.distribution[rand.Intn(p.size)]
	return curindex,nil
}
