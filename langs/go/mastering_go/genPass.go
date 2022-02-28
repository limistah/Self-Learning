package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

const MIN = 0
const MAX = 95

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)

		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	var length int64 = 10
	if len(os.Args) == 2 {
		arguments := os.Args
		length, err := strconv.ParseInt(arguments[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("unknown argument: %d", length)
			log.Panicf("Gen Pass: %s", err)
		}
	}

	//fmt.Println(getString(length))
	b, _ := generateBytes(length)
	converted := base64.URLEncoding.EncodeToString(b)
	fmt.Println(converted)
}
