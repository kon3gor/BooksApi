package utils

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func GenerateThreeRandoms(bottom int, top int) []int {
	result :=  make([]int, 3)
	rand.Seed(time.Now().UnixNano())	
	
	v := rand.Intn(top - bottom) + bottom
	result[0] = v

	for result[0] == v {
		v = rand.Intn(top - bottom) + bottom
	}
	result[1] = v

	for result[1] == v || result[0] == v {
		v = rand.Intn(top - bottom) + bottom
	}
	result[2] = v

	return result
}

func MakeLikedQuery(bookIds []string) string {
	var sb strings.Builder
	fmt.Fprint(&sb, "key:")
	for _, item := range bookIds {
		fmt.Fprintf(&sb, "/works/%s+OR+", item)
	}
	return strings.TrimSuffix(sb.String(), "+OR+")
}
