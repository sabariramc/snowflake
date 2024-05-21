// Package snowflake implements a generator for unique snowflake IDs.
// For more information on snowflake IDs, refer to: https://en.wikipedia.org/wiki/Snowflake_ID
package snowflake

import (
	"fmt"
	"sync"
	"time"
)

// ID represents a snowflake ID.
type ID int64

// Snowflake is a structure that holds the configuration and state of a snowflake ID generator.
type Snowflake struct {
	machineID      int64
	sequenceNumber int64
	lock           sync.Mutex
	lastTs         int64
	maxSequence    int64
	maxTimestamp   int64
	timestampShift int64
	c              Config
}

// calculateMaxForMask calculates the maximum value for a given bit mask.
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

// New creates a new Snowflake instance with the provided options.
func New(options ...Options) (*Snowflake, error) {
	config := defaultConfig
	for _, fu := range options {
		fu(&config)
	}
	if config.MachineIdMask+config.SequenceIdMask+uint8(config.TimestampMask) != 63 || config.SequenceIdMask == 0 || config.MachineIdMask == 0 || config.TimestampMask == 0 {
		return nil, fmt.Errorf("invalid mask config for Snowflake")
	}
	if config.Epoch > time.Now().UnixMilli() {
		return nil, fmt.Errorf("invalid epoch config for Snowflake")
	}
	maxSequence := calculateMaxForMask(int(config.SequenceIdMask))
	if config.MachineId > calculateMaxForMask(int(config.MachineIdMask)) {
		return nil, fmt.Errorf("invalid maxMachineId config for Snowflake")
	}
	if config.SequenceNo > maxSequence {
		return nil, fmt.Errorf("invalid sequenceNo config for Snowflake")
	}
	return &Snowflake{
		machineID:      config.MachineId << int64(config.SequenceIdMask),
		sequenceNumber: config.SequenceNo,
		maxSequence:    maxSequence,
		maxTimestamp:   calculateMaxForMask(int(config.TimestampMask)),
		timestampShift: int64(config.MachineIdMask + config.SequenceIdMask),
		c:              config,
	}, nil
}

// Stats returns the current machine ID and sequence number of the Snowflake instance.
func (s *Snowflake) Stats() (int64, int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.c.MachineId, s.sequenceNumber
}

// GenerateID generates a new snowflake ID based on the current timestamp, machine ID, and sequence number.
func (s *Snowflake) GenerateID() ID {
	var sequenceNumber int64
	s.lock.Lock()
	ts := time.Now().UnixMilli()
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
	ts -= s.c.Epoch
	if ts > s.maxTimestamp {
		panic("timestamp exceeds max limit")
	}
	return ID(ts<<s.timestampShift | s.machineID | sequenceNumber)
}
