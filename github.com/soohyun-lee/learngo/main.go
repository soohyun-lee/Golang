package main

import (
	"errors"
	"fmt"
	"net/http"
)

// func lenAndUpper(name string) (lenght int, uppercase string) {
// 	defer fmt.Println(("이것은 끝!"))
// 	lenght = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func superAdd(numbers ...int) int {
// 	// for i := 0; i < len(numbers); i++ {
// 	// 	fmt.Println(numbers[i])
// 	// }
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func canIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		return false
// 	}
// 	return true
// }

// // swtich
// func canIDrink(age int) bool {
// 	switch {
// 	case age < 10:
// 		return false
// 	case age == 18:
// 		return true
// 	case age > 50:
// 		return false
// 	}
// 	return false
// }

// func canIDrink(age int) bool {
// 	switch koreanAgeage := age + 2; koreanAgeage {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	fmt.Println(canIDrink(18))
// }

// func retpeatMe(words ...string) string {
// 	fmt.Println(words)
// }

// func main() {
// 	totalLenght, _ := lenAndUpper("soohyun")
// 	fmt.Println(totalLenght)
// }

// func main() {
// 	result := superAdd(1, 2, 3, 4, 5, 6)ㅋ
// 	fmt.Println(result)
// }

// func main() {
// 	a := 2
// 	b := &a
// 	a = 5
// 	fmt.Println(*b)
// }

// func main() {
// 	names := []string{"nico", "soo", "jaw"}
// 	names = append(names, "sooohyun")
// 	fmt.Println(names)
// }

// func main() {
// 	nico := map[string]string{"name": "soohyun", "age": "12"}
// 	for _, value := range nico {
// 		fmt.Println(value)
// 	}
// }
// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }

// func main() {
// 	favFood := []string{"dddd", "ddadfa"}
// 	nico := person{name: "soohyun", age: 16, favFood: favFood}
// 	fmt.Println(nico)
// // }

// func main() {
// 	account := banking.Account{Owner: "nicolas", Balance: 1000}
// 	fmt.Println(account)
// }

// func main() {
// 	account := accounts.NewAccount("nico")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.)
// }

// func main() {
// 	dictionary := mydict.Dictionary{"first": "First word"}
// 	definition, err := dictionary.Search("first")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(definition)
// 	}
// }

// func main() {
// 	dictionary := mydict.Dictionary{}
// 	word := "hello"
// 	definition := "Greeting"
// 	err := dictionary.Add(word, definition)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	hello, _ := dictionary.Search(word)
// 	fmt.Println("found", word, "definition:", hello)
// 	err2 := dictionary.Add(word, definition)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}
// }

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.facebook.com/",
		"https://www.reddit.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}

// func main() {
// 	c := make(chan string)
// 	people := [4]string{"nico", "flynn", "flynn2", "flynn3"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + "is sexy"
// }

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }
