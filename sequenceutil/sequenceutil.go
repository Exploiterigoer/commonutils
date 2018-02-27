package sequenceutil

import (
	"sync"
)

var m sync.Mutex
var sequenceId uint32

// SequenceGenerator 流水号发生器
func SequenceGenerator() uint32 {
	m.Lock()
	sequenceId++
	if sequenceId > (^uint32(0)) {
		sequenceId = 0
	}
	m.Unlock()
	return sequenceId
}
