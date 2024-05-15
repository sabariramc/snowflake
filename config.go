package snowflake

import "time"

type Config struct {
	epoch          int64
	machineIdMask  uint8
	sequenceIdMask uint8
	machineId      int64
	sequenceNo     int64
}

var defaultConfig = Config{
	epoch:          0,
	machineIdMask:  10,
	sequenceIdMask: 12,
	machineId:      1,
	sequenceNo:     0,
}

type Options func(*Config)

// WithEpoch option is to rest the offset of epoch for snowflake id generation
func WithEpoch(t time.Time) Options {
	return func(c *Config) {
		c.epoch = t.UnixMilli()
	}
}

func WithMachineIdMask(m uint8) Options {
	return func(c *Config) {
		c.machineIdMask = m
	}
}

func WithSequenceIdMask(m uint8) Options {
	return func(c *Config) {
		c.sequenceIdMask = m
	}
}

func WithMachineId(s uint32) Options {
	return func(c *Config) {
		c.machineId = int64(s)
	}
}

func WithSequenceNo(s uint32) Options {
	return func(c *Config) {
		c.sequenceNo = int64(s)
	}
}
