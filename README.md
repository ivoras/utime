# utime - a faster Time library for Go

The goal of this library is to implement many of the Go's time functions on a 32-bit Unix timestamp format (seconds since the Epoch). The primary reason for this is memory savings.

Go's `time.Time` consists of two 64-bit values and a pointer, making it 24 bytes in size on amd64. I happen to have a project which stores a lot of timestamps, so that eats up memory like crazy. Since I don't need the extended precision, this library's Time type is just 4 bytes.

Should you use this library? Probably not, unless you know what you're doing. This library will mostly help in circumstances when there are a lot of in-memory structures holding timestamps. As an example, in my own case, there were three `time.Time` values in a struct, amounting to 72 bytes, which was reduced to 12 bytes with this library.

Some benchmarks:

On Linux:
```
Sizeof time.Time: 20 Sizeof utime.Time: 4
goos: linux
goarch: arm
pkg: github.com/ivoras/utime
BenchmarkUNow-4          1289191               918 ns/op
BenchmarkNow-4           1457226               828 ns/op
BenchmarkUSince-4        1311031               932 ns/op
BenchmarkSince-4         1841196               662 ns/op
BenchmarkUSub-4         236264396                5.01 ns/op
BenchmarkSub-4           6587425               182 ns/op
PASS
ok      github.com/ivoras/utime 11.334s
```

On Windows:
```
Sizeof time.Time: 24 Sizeof utime.Time: 4
goos: windows
goarch: amd64
pkg: github.com/ivoras/utime
BenchmarkUNow-4         169228026                7.16 ns/op
BenchmarkNow-4          230071538                5.14 ns/op
BenchmarkUSince-4       163814563                7.23 ns/op
BenchmarkSince-4        199206692                6.05 ns/op
BenchmarkUSub-4         1000000000               0.262 ns/op
BenchmarkSub-4          410653164                2.91 ns/op
PASS
ok      github.com/ivoras/utime 10.173s
```

The takeaway here is that complex calls such as `Now()` take a bit more time than the regular Go's library implementation, but the simple calls like `Sub()` become virtually free
(in addition to the reduced memory usage).

The initial implementation tried to directly use appropriate syscalls on Linux and Windows to get the system time in the `Now()` call, but recent Go versions (1.13+) do something
much faster in their `time.Now()` implementation, so now we just use that.

This data type is also specifically compatible with [go-pg](https://github.com/go-pg/pg), where it can parse the TIMESTAMP data type.
