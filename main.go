package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("0~2の数字で手を表します'\n'0 = パー '\n'1 = チョキ'\n'2 = グー")
	fmt.Println("ジャンケン開始！！！！！")
	RockPaperScissors()
}

func RockPaperScissors() bool {

	flags := true
	for flags {
		IntInput, _ := IntStdin()

		rand.Seed(time.Now().UnixNano())
		var randomAction int = rand.Intn(3)
		fmt.Println(randomAction)

		if IntInput == randomAction {
			fmt.Println("draw")
			fmt.Println("もう一度入力してください")
			flags = true

		} else if IntInput == 2 {
			if IntInput > randomAction {
				fmt.Println("lose")
				flags = false
			} else {
				fmt.Println("win")
				flags = false
			}
		} else if IntInput > randomAction {
			fmt.Println("win")
			flags = false
		} else {
			fmt.Println("lose")
			flags = false
		}
	}
	return true
}

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}

func IntStdin() (int, error) {
	stringInput := StrStdin()
	return strconv.Atoi(strings.TrimSpace(stringInput))
}
