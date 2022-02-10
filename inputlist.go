package inputlist

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetListFloat64(format string) map[int]float64 {
	cin := bufio.NewScanner(os.Stdin)
	input := make(map[int]float64)
	fmt.Print(format)
	for i := 0; cin.Scan(); i++ {
		text := strings.TrimSpace(cin.Text())
		if text == "" {
			i--
			continue
		}
		input[i], _ = strconv.ParseFloat(text, 64)
	}
	return input
}
