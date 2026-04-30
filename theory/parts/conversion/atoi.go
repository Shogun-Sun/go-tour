package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringNumber := "100"
	// Ascii to integer
	// Возвращает результат первой переменной и в случае ошибки - вторую переменную с описанием ошибки
	i, _ := strconv.Atoi(stringNumber) // Преобразование строки в десятичное целое число

	fmt.Println("Преобразованное число:", i)

	// Integer to ascii
	integerNumber := 100
	s := strconv.Itoa(integerNumber)

	fmt.Println("Полученная строка:", s)
}