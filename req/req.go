package req

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

func GetJSON(url string, target interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return xerrors.Errorf("URL: %v returned status code %v", url, resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		return xerrors.Errorf("could not decode url: %v, error: %w", url, err)
	}

	return nil
}

func Download(url string, folder string) (bool, error) {
	fileName := filepath.Base(url)
	filePath := filepath.Join(folder, fileName)

	if _, err := os.Stat(filePath); err == nil {
		return false, nil
	}

	err := os.MkdirAll(folder, 0755)

	if err != nil {
		return false, xerrors.Errorf("could not create download folder: %w", err)
	}

	resp, err := http.Get(url)

	if err != nil {
		return false, xerrors.Errorf("could not download file: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return false, xerrors.Errorf("URL: %v returned status code %v", url, resp.StatusCode)
	}

	file, err := os.Create(filePath)

	if err != nil {
		return false, xerrors.Errorf("could not create file for download: %w", err)
	}

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return false, xerrors.Errorf("could not copy from from request body to file: %w", err)
	}

	return true, nil
}
