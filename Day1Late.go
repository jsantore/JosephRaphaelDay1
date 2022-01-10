package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	//var response, err = http.Get("https://www.google.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer response.Body.Close()
	//var result, err1 = ioutil.ReadAll(response.Body)
	//if err1 != nil {
	//	log.Fatalln("hackers we went bad")
	//}
	//if len(result) > 500 {
	//	fmt.Println(string(result))
	//} else if len(result) > 100 {
	//	fmt.Println("thats a pretty sus impostor")
	//} else {
	//	fmt.Println("wow that was short")
	//}
	//fmt.Print(response.Body)
	countDown(25)

}

func countDown(start int) { //in honor of the falcon heavy launch
	counter := start
	for {
		fmt.Println(counter)
		time.Sleep(1 * time.Millisecond)
		counter--
		if counter <= 0 {
			break
		}
	}
	//fmt.Println("blastoff")
	//time.Sleep(3 * time.Second)
	//fmt.Println("Success!")
	//time.Sleep(3 * time.Second)
	//fmt.Println("Sucsuss!")
	//time.Sleep(1 * time.Second)
	fmt.Println("with rune count", utf8.RuneCountInString("oh no ඞ"))

	fmt.Println("with len", len("oh no 😀"))
	fmt.Println("oh no the impostor ඞ")
}

// 1 error
