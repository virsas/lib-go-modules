package vssutil

import (
	"bufio"
	"net/http"
	"strconv"
	"strings"
)

func GetExchangeRate(cur string) (float64, error) {
	var err error
	var rate string

	client := http.Client{}

	req, err := http.NewRequest("GET", "https://www.cnb.cz/cs/financni-trhy/devizovy-trh/kurzy-devizoveho-trhu/kurzy-devizoveho-trhu/denni_kurz.txt", nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36`)
	resp, err := client.Do(req)
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
