package arraysgfg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Create_Biggest_Number(arr []string) string {
	sort.Slice(arr, func(i, j int) bool {
		x, _ := strconv.Atoi(fmt.Sprintf("%s%s", arr[i],arr[j]))
		y, _ := strconv.Atoi(fmt.Sprintf("%s%s", arr[j],arr[i]))
		return x > y
	})
	return strings.Join(arr, "")
}