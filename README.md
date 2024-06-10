## syncpool
In Golang if a massive amount of objects are allocated repeatedly, It causes a huge workload of GC.
Golang provides sync.Pool to decrease the allocations and GC workload. But sync.Pool is low level package and its not easy to use.
syncpool package is a wrapper on sync.pool and some other related packages that provides a high level easy solution to solve the problem without engaging with low level codes.
Briefly syncpool can provide a temp pool of any arbitrary objects easily. As well as it provides a slice of any objects.
In a especially case it provides SlicePool that can be used in many applications. for example we almost need BufferPool always in most of applications.

## Usage
syncpool provides a generic pool of any entity easily.

```go
var MyString string
var myStringPool = syncpool.NewPool[MyString]()
s:=myStringPool.Get()
s="something"
print(s)
myStringPool.Put(s)
```
## default pools
for more conviniency syncpool defines some traditional global pool. the following shows defaults pools.
```go 
var IntPool = NewPool[int]()
var BytePool = NewPool[byte]()
var ErrorPool = NewPool[error]()
var StringPool = NewPool[string]()

var IntSlicePool SlicePool[int]
var ByteSlicePool SlicePool[byte]
var ErrorSlicePool SlicePool[error]
var StringSlicePool SlicePool[string]
```

## SlicePool
Whenever we want to handle a slice objects instead a singular object we should use of SlicePool.
For example if we have to hold a group of []byte, BufferPool can be best candidate.

## Bechmarks
Benchmarks shows the main reason that why we should use of syncpool and how it can prevent memeory managements challenge in Golang.
```shell
BenchmarkJsonEncoderWithPool-4          42861     27023 ns/op    0    B/op    0    allocs/op
BenchmarkJsonMarshalWithoutPool-4       36492     29455 ns/op    4800 B/op    100  allocs/op
BenchmarkPool-4                         51396350  23.10 ns/op    0    B/op    0    allocs/op
BenchmarkNoPool-4                       19882788  59.78 ns/op    24   B/op    1    allocs/op
```

## License
# MIT 



