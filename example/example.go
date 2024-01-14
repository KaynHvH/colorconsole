package example

import (
	"fmt"
	"library/colorconsole"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func PrintExample() {
	colorconsole.SetTextColor(colorconsole.Blue)
	colorconsole.SetBackgroundColor(colorconsole.Yellow)
	fmt.Println("Hello, Colorful Console!")

	colorconsole.ResetColor()

	fmt.Println(colorconsole.ColoredText("Colored Text", colorconsole.Green, colorconsole.Black))

	fmt.Println("Example number 1")
	example1()
	fmt.Println("Example number 2")
	example2()
}

func example1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(2)

		go func(index int) {
			defer wg.Done()
			randomBackground := colorconsole.TextColor(rand.Intn(8))
			colorconsole.SetBackgroundColor(randomBackground)
			time.Sleep(100 * time.Millisecond)
		}(i)

		go func() {
			defer wg.Done()
			randomColor := colorconsole.TextColor(rand.Intn(8))
			colorconsole.SetTextColor(randomColor)
			time.Sleep(100 * time.Millisecond)
		}()
		fmt.Println(i)
	}
	wg.Wait()
	time.Sleep(time.Second)
	colorconsole.ResetColor()
}

func example2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			randomBackground := colorconsole.ColoredText(strconv.Itoa(i), colorconsole.TextColor(rand.Intn(8)), colorconsole.TextColor(rand.Intn(8)))
			time.Sleep(100 * time.Millisecond)
			fmt.Println(randomBackground)
		}()
	}
	wg.Wait()
	time.Sleep(time.Second)
	colorconsole.ResetColor()
}
