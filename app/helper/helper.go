package helper

import (
	"aoc/app/config"
	"io"
	"net/http"
)

func GetContent(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	// AOC requires cookie to be part of header
	req.Header.Add("Cookie", config.AppConfig.Cookie)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(content), nil
}
