package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

type options struct {
	letter  string
	idxBits int
	idxMask int64
	idxMax  float64
}

var src = rand.NewSource(time.Now().UnixNano())

func GenerateRandom(n int) string {
	return generate(n, options{
		letter:  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", // 62 letters
		idxBits: 6,                                                                // 00111111 = max 63 letters
		idxMask: (1 << 6) - 1,
		idxMax:  63 / 6,
	})
}

func GenerateUUIDv4() string {
	id := generate(32, options{
		letter:  "0123456789abcdefghijklmnopqrstuvwxyz", // 36 letters
		idxBits: 6,                                      // 00111111 = max 63 letters
		idxMask: (1 << 6) - 1,
		idxMax:  63 / 6,
	})

	return id[0:8] + "-" + id[8:12] + "-" + id[12:16] + "-" + id[16:20] + "-" + id[20:]
}

func GenerateOTP() string {
	return generate(6, options{
		letter:  "0123456789", // 10 letters
		idxBits: 4,            // 00001111 = max 15 letters
		idxMask: (1 << 4) - 1,
		idxMax:  15 / 4,
	})
}

func generate(n int, opt options) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), opt.idxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), opt.idxMax
		}
		if idx := int(cache & opt.idxMask); idx < len(opt.letter) {
			b[i] = opt.letter[idx]
			i--
		}
		cache >>= opt.idxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
