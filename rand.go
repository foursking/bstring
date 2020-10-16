package ffstring

import (
	"sync"
	"time"
)

var rngPool sync.Pool

func (r *ru) Uint32() uint32 {
	for r.x == 0 {
		x := time.Now().UnixNano()
		r.x = uint32((x >> 32) ^ x)
	}
	x := r.x
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.x = x
	return x
}

// RandUint32 returns pseudorandom uint32
func RandUint32() uint32 {
	v := rngPool.Get()
	if v == nil {
		v = &ru{}
	}
	r := v.(*ru)
	x := r.Uint32()
	rngPool.Put(r)
	return x
}

// RandUint32Max returns pseudorandom uint32 in the range [0..max)
func RandUint32Max(maxN uint32) uint32 {
	x := RandUint32()
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

// RandInt random numbers in the specified range
func RandInt(min int, max int) int {
	if max < min {
		max = min
	}
	return min + int(RandUint32Max(uint32(max+1-min)))
}

// RandString random string of specified length, the second parameter limit can only appear the specified character
func Rand(n int, tpl ...string) string {
	var s string
	b := make([]byte, n)
	if len(tpl) > 0 {
		s = tpl[0]
	} else {
		s = letterBytes
	}
	l := len(s) - 1
	for i := n - 1; i >= 0; i-- {
		idx := RandInt(0, l)
		b[i] = s[idx]
	}
	return Bytes2String(b)
}

var idWorkers, _ = NewIDWorker(0)

func UUID() int64 {
	id, _ := idWorkers.ID()
	return id
}
