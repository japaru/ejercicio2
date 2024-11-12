package main

import (
	"math/rand"
	"time"
)

func RandRunes(n int, cadena string) string {

	cadenaByte := []byte(cadena)

	rand.Seed(time.Now().UnixNano())

	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = cadenaByte[rand.Intn(len(cadenaByte))]
	}

	return string(result)
}

func main() {
	println(RandRunes(10, "abcdefghijklmnopqrstuvwxyz"))
}
