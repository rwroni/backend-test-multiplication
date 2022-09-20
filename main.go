package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masuk kan bilangan bulat : ")
	fmt.Scan(&n)
	i := 1

	for {
		if i > 99 {
			break
		}
		fmt.Println(n, " X ", i, " = ", n*i)
		i++
	}
}
