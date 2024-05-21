package snowflake

import "time"

// Config represents the configuration for the snowflake ID generator.
type Config struct {
	Epoch          int64 // Epoch represents the starting time from which the timestamp will be calculated.
	MachineIdMask  uint8 // MachineIdMask represents the number of bits reserved for the machine ID in the snowflake ID.
	SequenceIdMask uint8 // SequenceIdMask represents the number of bits reserved for the sequence ID in the snowflake ID.
	TimestampMask  int64 // TimestampMask represents the number of bits reserved for the timestamp in the snowflake ID.
	MachineId      int64 // MachineId represents the unique ID assigned to the machine generating the snowflake IDs.
	SequenceNo     int64 // SequenceNo represents the initial sequence number for the snowflake ID.
}

// defaultConfig represents the default configuration for the snowflake ID generator.
var defaultConfig = Config{
	Epoch:          0,
	MachineIdMask:  10,
	SequenceIdMask: 12,
	TimestampMask:  41,
	MachineId:      1,
	SequenceNo:     0,
}

// Options represents functional options for configuring the snowflake generator.
type Options func(*Config)

// WithEpoch sets the offset of epoch for snowflake ID generation.
func WithEpoch(t time.Time) Options {
	return func(c *Config) {
		c.Epoch = t.UnixMilli()
	}
}

// WithMachineIdMask sets the bit mask for the machine ID in the snowflake ID.
func WithMachineIdMask(m uint8) Options {
	return func(c *Config) {
		c.MachineIdMask = m
	}
}

// WithSequenceIdMask sets the bit mask for the sequence ID in the snowflake ID.
func WithSequenceIdMask(m uint8) Options {
	return func(c *Config) {
		c.SequenceIdMask = m
	}
}

// WithTimestampMask sets the bit mask for the timestamp in the snowflake ID.
func WithTimestampMask(m uint8) Options {
	return func(c *Config) {
		c.TimestampMask = int64(m)
	}
}

// WithMachineId sets the machine ID for the snowflake ID.
func WithMachineId(s uint32) Options {
	return func(c *Config) {
		c.MachineId = int64(s)
	}
}

// WithSequenceNo sets the initial sequence number for the snowflake ID.
func WithSequenceNo(s uint32) Options {
	return func(c *Config) {
		c.SequenceNo = int64(s)
	}
}
