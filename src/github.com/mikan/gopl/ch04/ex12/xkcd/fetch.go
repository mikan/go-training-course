package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const idToReplace = "{ID}"

func FetchLatestNum() (int, error) {
	// Fetch latest
	resp, err := http.Get("http://xkcd.com/info.0.json")
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return -1, fmt.Errorf("fetch failed: %s", resp.Status)
	}

	// Parse
	var latest Comic
	if err := json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		resp.Body.Close()
		return -1, err
	}
	resp.Body.Close()
	return latest.Num, nil
}

func FetchAll(latest int) ([]string, error) {
	dir := os.Getenv("GOPATH") + "/out/xkcd"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Target dir: " + dir)

	url := "https://xkcd.com/" + idToReplace + "/info.0.json"
	for i := 1; i <= latest; i++ {
		path := dir + "/" + strconv.Itoa(i) + ".json"
		_, err := os.Stat(path)
		if err == nil {
			continue
		}
		resp, err := http.Get(strings.Replace(url, idToReplace, strconv.Itoa(i), 1))
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == http.StatusNotFound {
			fmt.Print("!") // Oops! #404 is not found! Is it a joke?
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Println(" ERROR!")
			return nil, fmt.Errorf("fetch failed: %s", resp.Status)
		}
		fmt.Print(".")

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(path, body, 0644)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
	}
	fmt.Println(" done.")

	files, _ := filepath.Glob(dir + "/*")
	return files, nil
}

func GetCachedAll() ([]string, error) {
	dir := os.Getenv("GOPATH") + "/out/xkcd"
	files, err := filepath.Glob(dir + "/*")
	if err != nil {
		return nil, err
	}
	return files, nil
}
