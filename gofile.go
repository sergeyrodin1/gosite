package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция для отправки сообщения в канал
func sendMessage(ch chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	ch <- msg
}

// Функция для вычисления суммы чисел
func sum(numbers []int, result chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	result <- sum // Отправляем результат в канал
}

// Функция для демонстрации работы с WaitGroup
func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик после завершения горутины
	fmt.Println(msg)
}

func main() {
	// 1. Работа с каналами и горутинами
	fmt.Println("1. Демонстрация горутин и каналов:")
	ch := make(chan string)

	go sendMessage(ch, "Сообщение из горутины!", 1*time.Second)
	fmt.Println(<-ch) // Получаем сообщение из канала

	// 2. Работа с буферизованными каналами
	fmt.Println("\n2. Буферизованные каналы:")
	bufferedCh := make(chan int, 3)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	close(bufferedCh) // Закрываем канал
	for val := range bufferedCh {
		fmt.Println(val)
	}

	// 3. Использование WaitGroup для синхронизации
	fmt.Println("\n3. WaitGroup для синхронизации:")
	var wg sync.WaitGroup
	wg.Add(2)

	go printMessage("Привет, мир!", &wg)
	go printMessage("Golang — это круто!", &wg)

	wg.Wait() // Ожидаем завершения всех горутин

	// 4. Работа с select и таймаутами
	fmt.Println("\n4. Select и таймауты:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMessage(ch1, "Данные из ch1", 1*time.Second)
	go sendMessage(ch2, "Данные из ch2", 2*time.Second)

	select {
	case msg1 := <-ch1:
		fmt.Println("Получено:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Получено:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Таймаут!")
	}

	// 5. Вычисление суммы чисел параллельно
	fmt.Println("\n5. Параллельное вычисление суммы:")
	numbers := []int{1, 2, 3, 4, 5, 6}
	result := make(chan int)

	go sum(numbers[:len(numbers)/2], result) // Первая половина
	go sum(numbers[len(numbers)/2:], result) // Вторая половина

	sum1 := <-result
	sum2 := <-result

	fmt.Printf("Итоговая сумма: %d\n", sum1+sum2)

	// Завершение
	fmt.Println("\nПрограмма завершена.")
}
