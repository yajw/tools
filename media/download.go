package media

import (
	"io"
	"net/http"
	"os"
)

func Download(url string, targetFile string) (int64, error) {
	file, err := os.Create(targetFile)
	if err != nil {
		return 0, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return 0, err
	}

	return size, err
}
