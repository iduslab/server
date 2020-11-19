package utils

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func DownloadImageViaURL(url, filePath string) error {
	response, e := http.Get(url)
	if e != nil {
		return e
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func Hex2RGB(hex string) (uint8, uint8, uint8, error) {
	values, err := strconv.ParseUint(hex, 16, 32)

	if err != nil {
		return 0, 0, 0, err
	}

	return uint8(values >> 16), uint8((values >> 8) & 0xFF), uint8(values & 0xFF), nil
}
