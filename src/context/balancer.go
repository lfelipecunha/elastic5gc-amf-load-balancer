package context

import (
	"sync"
)

type BalancerStrategy interface {
	SelectAmf() *AmfData
	AddAmf(*AmfData)
	RemoveAmf(string)
	Lock()
	Unlock()
	GetAmfs() []*AmfData
}

type RoundinRobin struct {
	Amfs    []*AmfData
	Current int
	Mutex   sync.Mutex
}

func (rr *RoundinRobin) SelectAmf() *AmfData {
	var index int
	rr.Lock()
	index = rr.Current
	rr.Current++
	if rr.Current >= len(rr.Amfs) {
		rr.Current = 0
	}
	rr.Unlock()
	return rr.Amfs[index]
}

func (rr *RoundinRobin) Lock() {
	rr.Mutex.Lock()
}

func (rr *RoundinRobin) Unlock() {
	rr.Mutex.Unlock()
}

func (rr *RoundinRobin) AddAmf(amfData *AmfData) {
	rr.Lock()
	exists := false
	for _, r := range rr.Amfs {
		if r.ID == amfData.ID {
			// do not add existent AMF
			exists = true
			break
		}
	}
	if !exists {
		rr.Amfs = append(rr.Amfs, amfData)
	}
	rr.Unlock()
}

func (rr *RoundinRobin) RemoveAmf(id string) {
	rr.Lock()
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
	rr.Unlock()
}

func (rr *RoundinRobin) GetAmfs() []*AmfData {
	return rr.Amfs
}
