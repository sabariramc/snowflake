package snowflake_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/sabariramc/snowflake"
	"gotest.tools/assert"
)

func Example_epochreset() {
	//Epoch is reset to the current time, with default machine id as `1` and default starting sequence as `0` the first id will be always Ox1000 == 4096
	s, _ := snowflake.New(snowflake.WithEpoch(time.Now()))
	id := s.GenerateID()
	fmt.Printf("%X\n", id)
	fmt.Println(id)
	//Epoch is reset to the current time, machine id as `2` and starting sequence as `10` the first id will be always Ox200A == 8202
	s, _ = snowflake.New(snowflake.WithEpoch(time.Now()), snowflake.WithMachineId(2), snowflake.WithSequenceNo(10))
	id = s.GenerateID()
	fmt.Printf("%X\n", id)
	fmt.Println(id)
	//Output:
	//1000
	//4096
	//200A
	//8202
}

func Example_maskchange() {
	s, _ := snowflake.New(snowflake.WithEpoch(time.Now()), snowflake.WithMachineIdMask(5), snowflake.WithSequenceIdMask(17))
	id := s.GenerateID()
	fmt.Printf("%X\n", id)
	fmt.Println(id)
	s, _ = snowflake.New(snowflake.WithEpoch(time.Now()), snowflake.WithMachineIdMask(5), snowflake.WithSequenceIdMask(17), snowflake.WithMachineId(10), snowflake.WithSequenceNo(100))
	id = s.GenerateID()
	fmt.Printf("%X\n", id)
	fmt.Println(id)
	//Output:
	//20000
	//131072
	//140064
	//1310820
}

func Example() {
	s, _ := snowflake.New()
	id := s.GenerateID()
	fmt.Println(id)
}

func TestSnowflake(t *testing.T) {
	s, err := snowflake.New()
	assert.NilError(t, err)
	fmt.Printf("%b\n", time.Now().UnixMilli())
	for i := 0; i < 10000; i++ {
		id := s.GenerateID()
		fmt.Printf("%b\n", id)
	}
}

func TestTimestampMask(t *testing.T) {
	cnt := 4097
	idList := make([]snowflake.ID, cnt)
	s, _ := snowflake.New(snowflake.WithTimestampMask(43), snowflake.WithMachineIdMask(8), snowflake.WithEpoch(time.Now()))
	for i := 0; i < cnt; i++ {
		idList[i] = s.GenerateID()
	}
	assert.Equal(t, idList[0], snowflake.ID(0x1000))
	assert.Equal(t, idList[cnt-1], snowflake.ID(0x101000))
}

func TestSnowflakeTimer(t *testing.T) {
	s, err := snowflake.New()
	assert.NilError(t, err)
	st := time.Now()
	for i := 0; i < 10000; i++ {
		s.GenerateID()
	}
	assert.Assert(t, 6 >= time.Since(st).Milliseconds())
}

func TestSnowflakeDuplicate(t *testing.T) {
	s, err := snowflake.New()
	assert.NilError(t, err)
	totalN := 1000000
	ch := make(chan snowflake.ID, totalN)
	var wg sync.WaitGroup
	concurrencyFactor := 100
	for i := 0; i < concurrencyFactor; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < totalN/concurrencyFactor; j++ {
				id := s.GenerateID()
				ch <- id
			}
			wg.Done()
		}()
	}
	wg.Add(1)
	duplicateCount := 0
	go func() {
		idSet := make(map[snowflake.ID]int, totalN)
		total := 0
		for id := range ch {
			if _, ok := idSet[id]; ok {
				duplicateCount++
			}
			idSet[id]++
			total++
			if total == totalN {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	assert.Equal(t, duplicateCount, 0)
}

func BenchmarkSnowflake(b *testing.B) {
	s, err := snowflake.New()
	assert.NilError(b, err)
	var goprocs = runtime.GOMAXPROCS(0)
	for i := 1; i < 1000; i += 50 {
		b.Run(fmt.Sprintf("goroutines-%d", i*goprocs), func(b *testing.B) {
			b.SetParallelism(i)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					s.GenerateID()
				}
			})
		})
	}
}
