package main

import (
	"fmt"
	"math/rand"
)

// main является функцией, с которой все начинается
// убрать комментарии с нужной сортировки, а активную закомментировать
func main() {
	size := 10000
	array := generateArray(size)
	fmt.Println("Исходный массив:", array)
	//fmt.Println("Сортировка выборкой:", selection(array, size))
	//fmt.Println("Сортировка вставками:", insertion(array))
	//fmt.Println("Сортировка методом пузырька: ", bubble(array))
	fmt.Println("Быстрая сортировка", quickSort(array))

}

// возвращает true, если элемент есть в массиве
func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// массив для генерации массива k длины
// заполняется случайными неповторторяющимся числами
func generateArray(k int) []int {
	var array []int
	success := false
	// добавляю в массив первый элемент, чтобы просто так не гонять цикл
	array = append(array, rand.Intn(k*10))
	for i := 0; i <= k-1; i++ {
		// цикл будет генерировать случайное числа, пока не будет
		// сгенерировань число, которого нет в массиве, флаг станет success будет true
		for success == false {
			num := rand.Intn(k * 10)
			containsSuccess := contains(array, num)
			if containsSuccess == false {
				array = append(array, num)
				success = true
			}
		}
		success = false
	}
	return array

}

// сортировка выбором
// в неотсортированном подмассиве ищем минимум
// меняем с первым элементом подмассива местами
// O (n^2)
func selection(array []int, size int) []int {
	for i := range array {
		// создаем заведо максимальное числа
		min := size * 100
		//создаем переменную, которая хранит в себе индекс минимального элемента
		var minIndex int
		for j := i; j <= len(array)-1; j++ {
			if array[j] <= min {
				min = array[j]
				minIndex = j
			}
		}
		array[i], array[minIndex] = min, array[i]
	}
	return array
}

// Сортировка вставками
// на каждом шаге берем элемент и находим ему место для вставки
// O (n^2)
// если массив почти отсортирован, то O (n)
func insertion(array []int) []int {
	for i := 1; i <= len(array)-1; i++ {
		j := i - 1
		key := array[i]
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}

// Сортировка методом пузырька
// если следующий элемент массива больше предыдущего, меняем местами
// O (n^2)
// если массив почти отсортирован, то O (n)
func bubble(array []int) []int {
	for i := range array {
		for j := 0; j <= len(array)-i-2; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// Сортировка методом Хоаром или быстрая сортировка
// берем случайный элемент k последовательности (опорный элемент)
// разбиваем массив на две части
// в одной части будем хранить элементы, которые больше случайного элемента K
// в другой, те, которые меньше k
// в среднем  O (n log n)
// но если массив почти отсортирован или k - это найменьший элемент последовательности
// или найбольший, то тут уже O (n^2)

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	left, right := 0, len(array)-1

	k := rand.Int() % len(array)

	array[k], array[right] = array[right], array[k]

	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	quickSort(array[:left])
	quickSort(array[left+1:])

	return array

}
