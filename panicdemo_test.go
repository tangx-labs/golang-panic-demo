package panicdemo

import (
	"fmt"
	"testing"
	"time"
)

func Test_Panic(t *testing.T) {
	panicDemo()
}

func panicDemo() {
	// 1. panic 会终止服务
	// 2. panic 会沿着函数调用链， 逆向传染，直到退出或被捕获。
	// 3. 但是， panic 退出不会终止 defer
	// 4. 因此可以在 defer 中使用 recover 捕获 panic， 中断传染。
	// 5. 只能在相同 G 内的函数调用链中使用 defer ， 才能在任意一个环节捕获 panic。
	// 6. 根据 5. , 如果函数 panicfunc_2 被 go 出去， 则父函数将无法捕获 panicfunc_2 中的 panic

	defer func() {
		fmt.Printf("defer of panicDemo: ")
		if err := recover(); err != nil {
			fmt.Println("catch panic in panicDemo")
		}
		fmt.Println("")
		println()
	}()

	panicfunc()
	time.Sleep(10 * time.Second)
}

func panicfunc() {

	defer func() {
		fmt.Printf("defer of panicfunc: ")
		if err := recover(); err != nil {
			fmt.Println("catch panic in panicfunc")
		}
		fmt.Println("")
	}()

	go panicfunc_2()
	time.Sleep(3 * time.Second)
	fmt.Println("end panicfunc")
}

func panicfunc_2() {
	defer func() {
		fmt.Println("header of panicfunc")
	}()
	defer func() {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now())
	}()

	panicfunc_3()
	println("end panicfunc_2")
}

func panicfunc_3() {
	panic("panic in panicfunc_3")
}
