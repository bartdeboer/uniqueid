package uniqueid

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	lastTimestamp uint64
	incr          uint64
	mutex         sync.Mutex
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"

func pow(base, exponent uint64) uint64 {
	result := uint64(1)
	for exponent != 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}
	return result
}

func reverseBytes(b []byte) []byte {
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func generateUniqueID() string {
	mutex.Lock()
	defer mutex.Unlock()
	currentTimestamp := uint64(time.Now().UnixNano())
	if currentTimestamp == lastTimestamp {
		incr++
	} else {
		lastTimestamp = currentTimestamp
		incr = 0
	}
	randomPart := uint64(rand.Intn(1000))
	combined := (lastTimestamp * 100000) + (incr * 1000) + randomPart
	return EncodeBase62(combined)
}

func Encode(number uint64, base int) string {
	if number == 0 {
		return string(chars[0]) // Return the first character if number is 0
	}
	var encoded bytes.Buffer
	uint64Base := uint64(base)
	for number > 0 {
		remainder := number % uint64Base
		number /= uint64Base
		encoded.WriteByte(chars[remainder])
	}
	return string(reverseBytes(encoded.Bytes()))
}

// func encode2(number uint64, base int) string {
// 	if number == 0 {
// 		return string(chars[0]) // Return the first character if number is 0
// 	}
// 	var encoded []byte
// 	uint64Base := uint64(base)
// 	for number > 0 {
// 		remainder := number % uint64Base
// 		number /= uint64Base
// 		encoded = append([]byte{chars[remainder]}, encoded...)
// 	}
// 	return string(encoded)
// }

// func encode3(number uint64, base int) string {
// 	var encoded string
// 	uint64Base := uint64(base)
// 	for number > 0 {
// 		remainder := number % uint64Base
// 		number /= uint64Base
// 		encoded = string(chars[remainder]) + encoded
// 	}
// 	return encoded
// }

func Decode(encoded string, base int) (uint64, error) {
	var number uint64 = 0
	uint64Base := uint64(base)
	for i, char := range encoded {
		index := strings.IndexRune(chars, char)
		if index == -1 {
			return 0, fmt.Errorf("invalid character: %s\n", string(char))
		}
		number += uint64(index) * pow(uint64Base, uint64(len(encoded)-i-1))
	}
	return number, nil
}

func EncodeBase62(number uint64) string {
	return Encode(number, 62)
}

func DecodeBase62(coded string) (uint64, error) {
	return Decode(coded, 62)
}

func UnixTimestampID() string {
	return Encode(uint64(time.Now().Unix()), 62)
}

func UnixMilliTimestampID() string {
	return Encode(uint64(time.Now().UnixMilli()), 62)
}

func UnixMicroTimestampID() string {
	return Encode(uint64(time.Now().UnixMicro()), 62)
}

func UnixNanoTimestampID() string {
	return Encode(uint64(time.Now().UnixNano()), 62)
}

func LowerUnixTimestampID() string {
	return Encode(uint64(time.Now().Unix()), 36)
}

func LowerUnixMilliTimestampID() string {
	return Encode(uint64(time.Now().UnixMilli()), 36)
}

func LowerUnixMicroTimestampID() string {
	return Encode(uint64(time.Now().UnixMicro()), 36)
}

func LowerUnixNanoTimestampID() string {
	return Encode(uint64(time.Now().UnixNano()), 36)
}

func Generate() string {
	return generateUniqueID()
}
