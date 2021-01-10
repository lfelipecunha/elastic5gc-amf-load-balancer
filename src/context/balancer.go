package context

import (
	"amfLoadBalancer/src/logger"
	"sync"
)

type BalancerStrategy interface {
	SelectAmf(*RanUe) *AmfData
	AddAmf(*AmfData)
	RemoveAmf(string)
	Lock()
	Unlock()
	GetAmfs() []*AmfData
	BlockAmf(string)
}

type RoundinRobin struct {
	Amfs    []*AmfData
	Current int
	Mutex   sync.Mutex
	Ue2Amf  sync.Map // map[RanUE.RanUeNgapId] *AmfData
}

func (rr *RoundinRobin) SelectAmf(ranUe *RanUe) *AmfData {
	rr.Lock()
	amfData, ok := rr.Ue2Amf.Load(ranUe.RanUeNgapId)
	if ok && amfData != nil {
		data := amfData.(*AmfData)
		for _, amf := range rr.Amfs {
			if amf.ID == data.ID {
				rr.Unlock()
				return data
			}
		}

		logger.BalancerLog.Debugf("Changing AMF of UE[%s]", ranUe.RanUeNgapId)
	}

	var result *AmfData
	for {
		rr.Current++
		if rr.Current >= len(rr.Amfs) {
			rr.Current = 0
		}
		result = rr.Amfs[rr.Current]
		if !result.Blocked {
			break
		}
	}
	rr.Unlock()

	rr.Ue2Amf.Store(ranUe.RanUeNgapId, result)
	logger.BalancerLog.Debugf("Select AMF[%s] to UE[%s]", result.ID, ranUe.RanUeNgapId)
	return result
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

func (rr *RoundinRobin) BlockAmf(url string) {
	rr.Lock()
	for i := 0; i < len(rr.Amfs); i++ {
		if rr.Amfs[i].IP == url {
			logger.BalancerLog.Debugf("Blocking AMF[%s]", rr.Amfs[i].IP)
			rr.Amfs[i].Blocked = true
		}
	}
}
