package sequenceutil

import (
	"sync"
)

var m sync.Mutex
var sequenceID uint32

// SequenceGenerator serial number generator
func SequenceGenerator() uint32 {
	m.Lock()
	sequenceID++
	if sequenceID > (^uint32(0)) {
		sequenceID = 0
	}
	m.Unlock()
	return sequenceID
}
