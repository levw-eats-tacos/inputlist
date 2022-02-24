package inputlist

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TODO: Figure out templates/generics/interfaces
/**
 *  Description: gets a list of strings and then prompts for input based on them
 *  @Param format:  list of prompt strings
 *  @return: list of floats
 */
func GetListFloat64(format ...string) map[int]float64 {
	cin := bufio.NewScanner(os.Stdin)

	input := make(map[int]float64)
	fmt.Print(format[0])
	for i := 1; cin.Scan(); i++ {
		if i < len(format) {
			fmt.Print(format[i])
		}
		text := strings.TrimSpace(cin.Text())
		if text == "" {
			i--
			continue
		}
		input[i-1], _ = strconv.ParseFloat(text, 64)
	}
	return input
}

/**
 *  Description: gets a list of strings and then prompts for input based on them
 *  @Param format:  list of prompt strings
 *  @return: list of ints
 */
func GetListInt64(format ...string) map[int]int64 {
	cin := bufio.NewScanner(os.Stdin)

	input := make(map[int]int64)
	fmt.Print(format[0])
	for i := 1; cin.Scan(); i++ {
		if i < len(format) {
			fmt.Print(format[i])
		}
		text := strings.TrimSpace(cin.Text())
		if text == "" {
			i--
			continue
		}
		input[i-1], _ = strconv.ParseInt(text, 10, 64)
	}
	return input
}

/**
 *  Description: gets a list of strings and then prompts for input based on them
 *  @Param format:  list of prompt strings
 *  @return: list of ints
 */
func GetListString(format ...string) map[int]string {
	cin := bufio.NewScanner(os.Stdin)

	input := make(map[int]string)
	fmt.Print(format[0])
	for i := 1; cin.Scan(); i++ {
		if i < len(format) {
			fmt.Print(format[i])
		}
		text := strings.TrimSpace(cin.Text())
		if text == "" {
			i--
			continue
		}
		input[i-1] = text
	}
	return input
}
