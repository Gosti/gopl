package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var content []byte
	var err error
	if len(os.Args) == 1 {
		content, err = ioutil.ReadAll(os.Stdin)
	} else {
		content, err = ioutil.ReadFile(os.Args[1])
	}

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Gopl: the file does not exist.")
		} else {
			fmt.Println("Gopl: Huehoe it's embarrassing but an error occurs -> ", err)
		}
		return
	}

	url := "http://play.golang.org/share"

	resp, err := http.Post(url, "raw", bytes.NewBuffer(content))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("http://play.golang.org/p/" + string(body))
}
