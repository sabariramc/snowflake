# Snowflake
Implementation of [snowflake_id](https://en.wikipedia.org/wiki/Snowflake_ID)

In default configuration generates upto `4098 unique ids per millisecond`

## Usage

```go
s, _ := snowflake.New()
id := s.GenerateID()
```
For advance usage check examples in [test file]

[test file]: snowflake_test.go

## Benchmarks Results

```
goos: linux
goarch: amd64
pkg: github.com/sabariramc/snowflake
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkSnowflake/goroutines-8-8         	 4922703	       244.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-408-8       	 4790684	       249.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-808-8       	 4914920	       249.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-1208-8      	 4732038	       247.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-1608-8      	 4807579	       247.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-2008-8      	 4708533	       245.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-2408-8      	 4868290	       246.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-2808-8      	 4865926	       248.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-3208-8      	 4645274	       244.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-3608-8      	 4848453	       246.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-4008-8      	 4858170	       247.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-4408-8      	 4742110	       245.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-4808-8      	 4854056	       252.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-5208-8      	 4896246	       245.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-5608-8      	 4857505	       247.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-6008-8      	 4786718	       246.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-6408-8      	 4841553	       247.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-6808-8      	 4857250	       246.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-7208-8      	 4730719	       249.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-7608-8      	 4885704	       249.4 ns/op	       0 B/op	       0 allocs/op
```
