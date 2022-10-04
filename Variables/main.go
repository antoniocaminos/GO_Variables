package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "y no tipee el nemero solo presiona ENTER"

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction

	playTheGame(firstNumber, secondNumber, substraction, answer)
}
func playTheGame(firstNumber, secondNumber, substraction, answer int) {
	reader := bufio.NewReader(os.Stdin)
	//display welcome
	fmt.Println("adivina el numero")
	fmt.Println("------------------")
	fmt.Println("")

	fmt.Println("piensa un numero del 1 al 10", prompt)
	reader.ReadString('\n')
	//mostrar como jugar
	fmt.Println("multiplica tu numero por", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("ahora multillica el resultado por", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("divide el resultado por el numero que iniciaste", prompt)
	reader.ReadString('\n')

	fmt.Println("ahora resta", substraction, prompt)
	reader.ReadString('\n')
	//dar una respuesta+
	fmt.Println("el numero es", answer)

}
