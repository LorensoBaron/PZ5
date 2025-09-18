package main

import (
 "fmt"
 "time"
)

func main() {
 resultChan := make(chan string)

 go func() { 
  time.Sleep(3 * time.Second)
  resultChan <- "Результат!"
 }()

 select { 
 case result := <-resultChan:
  fmt.Println(result)
 case <-time.After(2 * time.Second):
  fmt.Println("Таймаут!")
 }

 fmt.Println("Готово.")
}

