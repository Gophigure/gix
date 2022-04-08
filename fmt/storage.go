package fmt

import (
	"math"
	"strconv"
)

const suffixes = "b B    KbKiBKBMbMiBMBGbGiBGBTbTiBTBPbPiBPBEbEiBEB"

var sizesSI = []uint64{
	999,
	1_000,
	1_000_000,
	1_000_000_000,
	1_000_000_000_000,
	1_000_000_000_000_000,
	1_000_000_000_000_000_000,
}

// Bytes formats using SI sizes. Only supports base 10.
func Bytes(b uint64) string {
	for i, size := range sizesSI {
		if i > 0 && b/1000 < size {
			val, idx := float64(b)/math.Pow(1000, float64(i))-0.005, 7*i+5
			return strconv.FormatFloat(val, 'f', 2, 64) + suffixes[idx:idx+2]
		} else if i == 0 && b < size {
			return strconv.FormatUint(b, 10) + suffixes[2:3]
		}
	}
	return "didn't match a size"
}
