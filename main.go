package main

import (
	"fmt"
)

func main() {
	searchEl := 25
	array := []int{3, 5, 12, 15, 15, 25, 27, 29, 32, 55, 74, 79, 90, 99, 99, 100}
	fmt.Println("Массив:", array)
	fmt.Println("Применим алгоритм бинарного поиска числа ", searchEl)
	index := len(array)
	for {
		index /= 2
		el := array[index]
		fmt.Println("Рассматриваемый массив = ", array)
		fmt.Println("Индекс текущего элемента в этом массиве = ", index)
		fmt.Println("Текущий элемент = ", el)
		if searchEl > el {
			array = array[index:]
		} else if searchEl < el {
			array = array[0:index]
		} else {
			fmt.Println("Число найдено")
			return
		}

		if index == 0 {
			fmt.Println("Число не найдено")
			return
		}
	}
}
