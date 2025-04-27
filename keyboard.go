package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloatInput() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, inputError := reader.ReadString('\n')
	if inputError != nil {
		return 0, inputError
	}

	processedInput := strings.TrimSpace(input)
	number, parseError := strconv.ParseFloat(processedInput, 64)
	if parseError != nil {
		return 0, parseError
	}

	return number, nil
}