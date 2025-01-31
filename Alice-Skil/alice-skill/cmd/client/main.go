package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/go-resty/resty/v2"
	"local/alice-skill/config"
)

func main() {
	endpoint := config.InitConfig()
	fmt.Println("Введите длинный URL")
	reader := bufio.NewReader(os.Stdin)
	long, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	long = strings.TrimSpace(long)
	client := resty.New()
    client.SetDebug(true)
	// пишем запрос
	// запрос методом POST должен, помимо заголовков, содержать тело
	// тело должно быть источником потокового чтения io.Reader
	response, err := client.R().
		SetFormData(map[string]string{"url": long}).
		Post(endpoint.BaseURL)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	if response.StatusCode() != 200 && response.StatusCode() != 201 {
		log.Fatalf("Server return error: %s, %v ", response.Status(), response.String())
	}
    result := response.String()

	
	// и печатаем его
    fmt.Println(response.StatusCode())
	fmt.Println(result)
}
