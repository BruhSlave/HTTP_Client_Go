package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "srv.msk01.gigacorp.local/_stats"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Ошибка при создании запроса ", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса ", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Ошибка чтения ответа ", err)
	}

	fmt.Println(string(body))
}
