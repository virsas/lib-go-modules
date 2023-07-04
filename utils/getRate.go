package utils

import (
	"bufio"
	"net/http"
	"strconv"
	"strings"
)

func GetRate(cur string) (float64, error) {
	var err error
	var rate string

	resp, err := http.Get("http://www.cnb.cz/cs/financni_trhy/devizovy_trh/kurzy_devizoveho_trhu/denni_kurz.txt")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		lineArray := strings.Split(scanner.Text(), "|")
		if len(lineArray) == 5 {
			if lineArray[3] == cur {
				rate = strings.Replace(lineArray[4], ",", ".", 1)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}
