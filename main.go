package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "https://www.yahoo.co.jp"
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			fmt.Println(res.Status)
		}()
	}
	wg.Wait()
}
