package application

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetData() string {
	resp, err := http.Get("https://mindlevelled.pythonanywhere.com/123/test")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Body)
	data, err := io.ReadAll(resp.Body)
	return string(data)
}
