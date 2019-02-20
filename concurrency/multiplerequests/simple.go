package multiplerequests

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func makeRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(content))
}

// MaxInMultipleRequestsSimple Fan-out algorithm
func MaxInMultipleRequestsSimple(urls []string) (int, error) {
	ch := make(chan int)

	for _, url := range urls {
		go func(url string) {
			code, _ := makeRequest(url)
			ch <- code
		}(url)
	}

	var max int
	for i := 0; i < len(urls); i++ {
		v := <-ch
		if v > max {
			max = v
		}
	}

	return max, nil
}
