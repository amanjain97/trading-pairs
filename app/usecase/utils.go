package usecase

import (
	"bufio"
	"fmt"
	"os"
	"tradingpairs/domain/models"
)

func writeToFile(pairs []models.TradingPairs, filename string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var f *os.File
	defer f.Close()

	filename = currentDir + "/" + filename

	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		err := os.MkdirAll(currentDir, 0700)
		if err != nil {
			return "", err
		}
	}
	f, err = os.Create(filename)
	if err != nil {
		return "", err
	}

	w := bufio.NewWriter(f)
	for _, p := range pairs {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%s/%s", p.BaseCurrency, p.QuoteCurrency))
	}
	if err = w.Flush(); err != nil {
		return "", err
	}

	return filename, nil
}
