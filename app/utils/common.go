package utils

import "strconv"

func AtoI(num string) uint64 {
	atoi, err := strconv.Atoi(num)
	if err != nil {
		panic("所需参数无法转换为uint64类型")
	}
	var u uint64 = uint64(atoi)
	return u
}
