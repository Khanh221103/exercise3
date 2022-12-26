package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

var m = make(map[string]int)

// --In ra các message theo thứ tự. (mutex, chan, waitGroup)
// mutex
func chanRoutine1() {
	log.Print("hello 1")
	mutex.Lock()
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		mutex.Unlock()
	}()
	log.Print("hello 2")
	// Time.Sleep(2 * time.Second)
	mutex.Lock()
}

// chan
func chanRoutine2() {
	ch := make(chan string)
	
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		<- ch
	}()
	log.Print("hello 2")
	ch <- "abc"
	// time.Sleep(2 * time.Second)
}

// waitGroup
func chanRoutine3() {
	var wg sync.WaitGroup
	wg.Add(1)
	log.Print("hello 1")
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
	}()
	log.Print("hello 2")
	wg.Wait()
}

// -- In ra message 3 trước message 2. (mutex, chan, waitGroup)
// mutex
func chanRoutine4() {
	log.Print("hello 1")
	mutex.Lock()
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		mutex.Unlock()
	}()
	mutex.Lock()
	log.Print("hello 2")
}

// chan
func chanRoutine5() {
	ch := make(chan string)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		ch <- "abc"
	}()
	<-ch
	log.Print("hello 2")
}

// waitGroup
func chanRoutine6() {
	log.Print("hello 1")
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	wg.Wait()
	log.Print("hello 2")
}

// tạo 1 biến X map[string]string và 3 goroutine cùng thêm dữ liệu vào X. Mỗi goroutine thêm 1000 key khác nhau
func number1(m map[string]int) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		string := strconv.Itoa(i)
		key := "a" + string
		m[key] = i
		mutex.Unlock()
	}
	wg.Done()
}
func number2(m map[string]int) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		string := strconv.Itoa(i)
		key := "b" + string
		m[key] = i
		mutex.Unlock()
	}
	wg.Done()
}
func number3(m map[string]int) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		string := strconv.Itoa(i)
		key := "c" + string
		m[key] = i
		mutex.Unlock()
	}
	wg.Done()
}

func ex2() {
	i := 0
	wg.Add(3)
	go number1(m)
	go number2(m)
	go number3(m)
	wg.Wait()
	for key, value := range m {
		fmt.Printf("key: %v, value: %d\n ", key, value)
		i++

		if i == 15 {
			break
		}
	}
	fmt.Println("Total value: ", i)
}

// Chạy đoạn chương trình dưới đây. Nếu có lỗi hãy thêm logic để nó chạy đúng.
// Lý giải nguyên nhân lỗi.
func errFunc() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		wg.Add(1) //++
		go func() {
			for j := 1; j < 10000; j++ {
				mutex.Lock() //++
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mutex.Unlock() //++
			}
		}()
		wg.Done() //++
	}
	wg.Wait()
	log.Print("done")
}

// đọc từng dòng file này nạp dữ liệu vào 1 buffer channel có size 10
// Chỉ được sử dụng 3 go routine. Kết quả xử lý xong ỉn ra màn hình + từ xong

// func ex4() {
// 	channel := make(chan string, 10)
// 	finish := make(chan bool)
// 	text, err := os.Open("file.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	num := 1
// 	line := []*Line{}
// 	defer text.Close()
// 	scanner := bufio.NewScanner(text)
// 	for scanner.Scan() {
// 		channel <- scanner.Text()
// 		for i := 0; i < 3; i++ {
// 			wg.Add(1)
// 			go work(finish, channel, scanner)
// 			wg.Done()
// 			wg.Wait()
// 		}
// 		time.Sleep(time.Millisecond)
// 		app := append(line, &Line{line_number: num, data: scanner.Text()})
// 		num++
// 		for i := range app {
// 			fmt.Printf("%v giá trị là: %v\n", app[i].line_number, app[i].data)
// 		}
// 	}
// 	fmt.Println("xong")
// }
// func work(finish chan bool, channel chan string, scanner *bufio.Scanner) {
// 	for range scanner.Text() {
// 		fmt.Println(<-channel)
// 	}
// 	finish <-false
// 	close(finish)
// 	close(channel)
// }


//=====================


// type Line struct {
// 	line_number int
// 	data        string
// }

// func worker(finish chan bool, id int, jobs <-chan string, wg *sync.WaitGroup) {
// 	num := 1
// 	li := []*Line{}
	
// 	for j := range jobs {
// 		fmt.Println("worker", id, "print", j)
// 		// time.Sleep(time.Millisecond)

// 		a := Line{line_number: num, data: j}
// 		new := append(li, &a)
// 		for i := range new {
// 			fmt.Printf("%v giá trị là: %v\n",num,new[i].data)
// 		}
// 		num ++
// 	}
// 	// wg.Done()
// 	// finish <- false
// 	// close(finish)
// }
// func ex4() {
// 	jobs := make(chan string, 10)
// 	finish := make(chan bool)

// 	// num:= 1
// 	// li := []*Line{}
// 	text, err := os.Open("file.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer text.Close()
// 	scanner := bufio.NewScanner(text)
// 	for w := 1; w <= 3; w++ {
// 		go worker(finish, w, jobs, &wg)
// 	}
// 	for scanner.Scan() {
// 		jobs <- scanner.Text()
// 		// wg.Add(1)
// 		// a := Line{line_number: num, data: scanner.Text()}
// 		// new := append(li, &a)
// 		// num ++
// 		// for i := range new {
// 		// // 	fmt.Printf("%v giá trị là: %v\n",new[i].line_number,new[i].data)
// 		// }
// 		// time.Sleep(time.Second)
// 	}
// 	close(jobs)
// 	fmt.Println("Xong")
// }

// =====================

type Line struct {
	line_number int
	data        string
}

func worker(finish chan bool, w int, jobs chan *Line, wg *sync.WaitGroup) {
	for {
		
	 	data, ok := <-jobs
		 if !ok {
			println("done")
			wg.Done()
			return
		}

		fmt.Println("worker:", w, "print:", data.data)
		fmt.Printf("%v gia tri la: %v\n", data.line_number, data.data)
		// time.Sleep(5 * time.Millisecond)
		// wg.Done()
		
	}

}

func ex4() {
	jobs := make(chan *Line, 10)
	finish := make([]chan bool, 4)
	// defer close(jobs)      

	// li := []*Line{};
	text, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer text.Close()
	
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		finish[w] = make(chan bool)
		go worker(finish[w], w, jobs, &wg)
	}

	num:= 1
	scanner := bufio.NewScanner(text)

	for scanner.Scan() {
		a := &Line{line_number: num, data: scanner.Text()}
		num ++
		jobs <- a
	}
	close(jobs)
	wg.Wait()
	fmt.Println("Xong")
	// time.Sleep(10 * time.Millisecond)
}

func main() {
	// //EX1
	// // in theo thứ tự
	// //sử dụng mutex
		// chanRoutine1()
		// fmt.Println("==========")
	// //sử dụng chan
		// chanRoutine2()
		// fmt.Println("==========")
	// //sửm dụng waitGroup
		// chanRoutine3()
		// fmt.Println("==========")

	// // in ra message 3 trước message 2
		// sử dụng mutex
	// chanRoutine4()
	// 	fmt.Println("==========")
		//sử dụng chan
	// chanRoutine5()
	// 	fmt.Println("==========")
		// sủ dụng waitGroup
	// chanRoutine6()
	// 	fmt.Println("==========")

	// //EX2
	// ex2()
	// fmt.Println("==========")

	// //EX3
	// errFunc()

	// //EX4
	ex4()


}
