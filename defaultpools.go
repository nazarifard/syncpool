package syncpool

var IntPool = NewPool[int]()
var BytePool = NewPool[byte]()
var ErrorPool = NewPool[error]()
var StringPool = NewPool[string]()

var IntSlicePool SlicePool[int]
var ByteSlicePool SlicePool[byte]
var ErrorSlicePool SlicePool[error]
var StringSlicePool SlicePool[string]

//var BytesPool SlicePool[byte] //default []byte pool
