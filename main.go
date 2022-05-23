package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func problem(err error) {
	if err != nil {
		panic(err)
	}
}

func bruteforce(val int64, randomsKey big.Int, lenght big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i int64
	start := time.Now()
	for i = 0; big.NewInt(i).Cmp(&lenght) == -1; i++ {
		if big.NewInt(i).Cmp(&randomsKey) == 0 {
			finish := time.Since(start)
			fmt.Println("*За", finish, "для ключа довжиною", val, "підібране значення:", i)
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	listLen := []int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	for _, val := range listLen {
		wg.Add(1)
		var a = big.NewInt(2)
		lenght := new(big.Int).Exp(a, big.NewInt(val), nil)
		randomsKey, err := rand.Int(rand.Reader, lenght)
		go problem(err)
		fmt.Println("Для ключа довжиною в ", val, ",простір якого рівний:", lenght, "біт.\n Вам згенеровано ключ:", randomsKey)
		bruteforce(val, *randomsKey, *lenght, &wg)
	}
	wg.Wait()
}
