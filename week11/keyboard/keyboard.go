package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func GetInteger() (int, error) {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}
