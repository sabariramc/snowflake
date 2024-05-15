// Package snowflake is a implement of [snowflake]
//
// [snowflake]: https://en.wikipedia.org/wiki/Snowflake_ID
package snowflake

import (
	"fmt"
	"sync"
	"time"
)

type ID int64

type Snowflake struct {
	machineID      int64
	sequenceNumber int64
	lock           sync.Mutex
	lastTs         int64
	maxSequence    int64
	c              Config
}

const maxTimestamp int64 = 0x1FFFFFFFFFF

func calculateMaxForMask(v int) int64 {
	if v > 63 || v <= 0 {
		return -1
	}
	res := 1
	for i := 1; i < v; i++ {
		res <<= 1
		res++
	}
	return int64(res)
}

func New(options ...Options) (*Snowflake, error) {
	config := defaultConfig
	for _, fu := range options {
		fu(&config)
	}
	if config.machineIdMask+config.sequenceIdMask != 22 || config.sequenceIdMask == 0 || config.machineIdMask == 0 {
		return nil, fmt.Errorf("invalid mask config for Snowflake")
	}
	if config.epoch > time.Now().UnixMilli() {
		return nil, fmt.Errorf("invalid epoch config for Snowflake")
	}
	maxSequence := calculateMaxForMask(int(config.sequenceIdMask))
	if config.machineId > calculateMaxForMask(int(config.machineIdMask)) {
		return nil, fmt.Errorf("invalid maxMachineId config for Snowflake")
	}
	if config.sequenceNo > maxSequence {
		return nil, fmt.Errorf("invalid sequenceNo config for Snowflake")
	}
	return &Snowflake{
		machineID:      config.machineId << int64(config.sequenceIdMask),
		sequenceNumber: config.sequenceNo,
		maxSequence:    maxSequence,
		c:              config,
	}, nil
}

func (s *Snowflake) Stats() (int64, int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.c.machineId, s.sequenceNumber
}

func (s *Snowflake) GenerateID() ID {
	var sequenceNumber int64
	s.lock.Lock()
	ts := time.Now().UnixMilli() - s.c.epoch
	if s.sequenceNumber > s.maxSequence {
		s.sequenceNumber = 0
		for ts == s.lastTs {
			ts = time.Now().UnixMilli()
		}
	}
	sequenceNumber = s.sequenceNumber
	s.sequenceNumber++
	s.lastTs = ts
	s.lock.Unlock()
	if ts > maxTimestamp {
		panic("timestamp exceed max limit")
	}
	return ID(ts<<22 | s.machineID | sequenceNumber)
}
