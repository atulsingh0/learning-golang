package main

import "fmt"

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
	TB
	PB
)

type byteSize float64

func (b byteSize) string() string {
	switch {
	case b >= PB:
		return fmt.Sprintf("%.2f PB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2f KB", b/KB)
	default:
		return fmt.Sprintf("%.2f B", b)
	}
}

func main() {
	//fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(byteSize(104800).string())
}
