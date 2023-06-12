package gen

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

type options struct {
	letter  string
	idxBits int
	idxMask int64
	idxMax  float64
}

// Random generates a random string of length n using alphanumeric characters.
func Random(n int) string {
	const (
		letter  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 62 letters
		idxBits = 6                                                                // 00111111 = max 63 letters
		idxMask = (1 << idxBits) - 1
		idxMax  = 63 / idxBits
	)

	return generate(n, options{letter, idxBits, idxMask, idxMax})
}

// UUIDv4 generates a random string following the UUIDv4 format.
func UUIDv4() string {
	const (
		letter  = "0123456789abcdef" // 16 letters (hexadecimal)
		idxBits = 4                  // 00001111 = max 15 letters
		idxMask = (1 << idxBits) - 1
		idxMax  = 15 / idxBits
	)

	id := generate(32, options{letter, idxBits, idxMask, idxMax})
	ninth := string(letter[rand.Intn(4)+8])

	// Set the fourth character of the third group to '4'
	// Set the ninth character of the fourth group to '8', '9', 'a', or 'b'
	return id[0:8] + "-" + id[8:12] + "-4" + id[13:16] + "-" + ninth + id[17:20] + "-" + id[20:]
}

// OTP generates a random string of length 6 following the OTP format using digits.
func OTP() string {
	const (
		letter  = "0123456789" // 10 letters (decimal)
		idxBits = 4            // 00001111 = max 15 letters
		idxMask = (1 << idxBits) - 1
		idxMax  = 15 / idxBits
	)

	return generate(6, options{letter, idxBits, idxMask, idxMax})
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

	return string(b)
	// return *(*string)(unsafe.Pointer(&b))
}
