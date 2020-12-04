package context

import (
	"sync"
)

type BalancerStrategy interface {
	SelectAmf() *AmfData
	AddAmf(*AmfData)
	RemoveAmf(string)
}

type RoundinRobin struct {
	Amfs    []*AmfData
	Current int
	Mutex   sync.Mutex
}

func (rr *RoundinRobin) SelectAmf() *AmfData {
	var index int
	rr.Mutex.Lock()
	index = rr.Current
	rr.Current++
	if rr.Current >= len(rr.Amfs) {
		rr.Current = 0
	}
	rr.Mutex.Unlock()
	return rr.Amfs[index]
}

func (rr *RoundinRobin) AddAmf(amfData *AmfData) {
	rr.Mutex.Lock()
	rr.Amfs = append(rr.Amfs, amfData)
	rr.Mutex.Unlock()
}

func (rr *RoundinRobin) RemoveAmf(id string) {
	rr.Mutex.Lock()
	var i int
	for i = 0; i < len(rr.Amfs); i++ {
		if rr.Amfs[i].ID == id {
			rr.Amfs = append(rr.Amfs[:i], rr.Amfs[i+1:]...)
			break
		}
	}
	if rr.Current >= len(rr.Amfs) {
		rr.Current = 0
	}
	rr.Mutex.Unlock()
}
