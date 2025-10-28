#实现了生成质数的功能。它从 2 开始，依次找出前 6 个质数，并将它们打印出来。

*/代码相较于普通写法是否有性能上的提升
普通的单线程实现需要依次检查每个数是否为质数，时间复杂度较高。
而这段代码利用了并发和通道的特性，多个 goroutine 可以同时进行整数的生成和过滤，提高了求解质数的速度。

m 个线程打印 n 个数，保证打印的有序性
可以使用 Go 语言的通道和互斥锁来实现。*/


package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i// 把自然数i发给通道ch
	}

}

func filter(in chan int, out chan int, prime int) {
	for {
		num := <-in
		if num%prime != 0 {
			out <- num// 如果num不能被prime整除，就把num发给输出通道out
		}
	}
}

func main() {
    ch := make(chan int) 
    go generate(ch) // 启动generate函数生成自然数序列

    for i := 0; i < 6; i++ {
        prime := <-ch // 从通道ch接收一个素数
        fmt.Printf("prime:%d\n", prime) // 打印素数

        out := make(chan int) 
        go filter(ch, out, prime) // 启动一个协程运行filter函数筛选数字
        ch = out 
    }
}
