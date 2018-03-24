package main

import (
	"fmt"
	"os"
	"strconv"
)

//牌序- K, Q, J,10, 9, 8, 7, 6, 5, 4, 3, 2, 1
//黑桃-51,50,49,48,47,46,45,44,43,42,41,40,39
//红桃-38,37,36,35,34,33,32,31,30,29,28,27,26
//梅花-25,24,23,22,21,20,19,18,17,16,15,14,13
//方片-12,11,10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
var pokerMap = map[int]string{1: "A", 11: "J", 12: "Q", 13: "K"}
var pokerColor = []string{"♦", "♣", "♥", "♠"}

const (
	pokerMax = 13
	colorMax = 4
)

func main() {
	params := os.Args
	if len(params) == 1 {
		for _, color := range pokerColor {
			for i := 0; i < pokerMax; i++ {
				value := i%pokerMax + 1
				if _, ok := pokerMap[value]; ok {
					fmt.Print(color, pokerMap[value])
				} else {
					fmt.Print(color, i%pokerMax+1)
				}

			}
			fmt.Println("")
		}
	}
	pokerString := ""
	for i := 1; i < len(params); i++ {
		number, err := strconv.Atoi(params[i])
		if err != nil {
			fmt.Println("Param ", params[i], " is not poker value.")
			pokerString += "X"
			continue
		}
		color := number / pokerMax
		if color >= colorMax {
			fmt.Println("Param ", params[i], " is not poker value.")
			pokerString += "X"
			continue
		}
		value := number%pokerMax + 1
		if _, ok := pokerMap[value]; ok {
			pokerString += fmt.Sprint(pokerColor[color], pokerMap[value])
		} else {
			pokerString += fmt.Sprint(pokerColor[color], value)
		}
	}
	fmt.Println(pokerString)
}
