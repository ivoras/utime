# utime - a faster Time library for Go

The goal of this library is to implement many of the Go's time functions on a 32-bit Unix timestamp format (seconds since the Epoch). The primary reasons for this are memory savings and performance.

Go's `time.Time` consists of two 64-bit values and a pointer, making it 24 bytes in size on amd64. This library's Time type is just 4 bytes.

Should you use this library? Probably not, unless you know what you're doing. This library will mostly help in circumstances when there are a lot of in-memory structures holding timestamps. In my own case, there were three `time.Time` values in a struct, amounting to 72 bytes, which was reduced to 12 bytes with this library.

Some benchmarks:

On Linux:
```
goos: linux
goarch: amd64
BenchmarkUNow-2         20000000               155 ns/op
BenchmarkNow-2          10000000               187 ns/op
```

On Windows:
```
goos: windows
goarch: amd64
BenchmarkUNow-4         100000000               15.6 ns/op
BenchmarkNow-4          200000000                8.54 ns/op
```

Go's runtime does some kind of magic on Windows which makes it around 25x faster than a regular `GetSystemTimeAsFileTime` system call, so here we just call `time.Now().Unix` instead. On linux, we're just calling the `time()` call. Obviously, these benchmarks were done on systems with differ in performance themselves, using go 1.10.3.