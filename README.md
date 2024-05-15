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
BenchmarkSnowflake/goroutines-8-8         	 4920124	       244.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-408-8       	 4902696	       250.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-808-8       	 4733638	       249.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-1208-8      	 4710220	       249.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-1608-8      	 4824832	       246.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-2008-8      	 4826086	       249.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-2408-8      	 4885420	       246.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-2808-8      	 4889487	       248.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-3208-8      	 4929100	       245.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-3608-8      	 4854890	       248.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-4008-8      	 4817530	       246.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-4408-8      	 4907401	       247.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-4808-8      	 4831908	       245.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-5208-8      	 4891472	       247.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-5608-8      	 4798198	       246.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-6008-8      	 4834064	       247.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-6408-8      	 4868370	       248.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-6808-8      	 4897899	       250.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-7208-8      	 4792924	       252.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnowflake/goroutines-7608-8      	 4780502	       247.6 ns/op	       0 B/op	       0 allocs/op
```
