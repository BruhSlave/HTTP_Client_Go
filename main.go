package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getData(url string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании запроса: %w ", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса: %w ", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %w ", err)
	}

	return (string(body)), nil
}

func main() {
	url := "http://srv.msk01.gigacorp.local/_stats"
	serverData, err := getData(url)
	if err != nil {
		log.Fatal(err)
	}

	partData := strings.Split(serverData, ",")

	loadAverage, _ := strconv.ParseFloat(partData[0], 64)
	amountRAM, _ := strconv.ParseFloat(partData[1], 64)
	usageRAM, _ := strconv.ParseFloat(partData[2], 64)
	amountDiscSpace, _ := strconv.ParseFloat(partData[3], 64)
	usageDiscSpace, _ := strconv.ParseFloat(partData[4], 64)
	networkBandwidth, _ := strconv.ParseFloat(partData[5], 64)
	networkUsage, _ := strconv.ParseFloat(partData[6], 64)

	if loadAverage > 30 {
		fmt.Printf("Load Average is too high: %d\n", int(loadAverage))
	}
	if percent := (usageRAM * 100) / amountRAM; percent > 80 {
		fmt.Printf("Memory usage too high: %d%%\n", int(percent))
	}
	if persent := (usageDiscSpace * 100) / amountDiscSpace; persent > 90 {
		freeDiscSpace := (amountDiscSpace - usageDiscSpace) / (1024 * 1024)
		fmt.Printf("Free disc space is too low: %d Mb left\n", int(freeDiscSpace))
	}
	if percent := (networkUsage * 100) / networkBandwidth; percent > 90 {
		freeBandwidth := (networkBandwidth - networkUsage) / 1000_000
		fmt.Printf("Network bandwidth usage high: %d Mbits/s available", int(freeBandwidth))
	}
}
