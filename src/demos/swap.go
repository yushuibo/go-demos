/*
 * @Since: 2020-02-23 16:28:50
 * @Author: shy
 * @Email: yushuibo@ebupt.com / hengchen2005@gmail.com
 * @Version: v1.0
 * @Description: -
 */
package main

import "fmt"

func swappedByArgs(a, b int) {
	a, b = b, a
	fmt.Printf("SwappedByArgs: a=%d, b=%d\n", a, b)
}

func swappedByAddr(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("SwappedByAddr: p1=%d, p2=%d\n", *p1, *p2)
}

func main() {
	a, b := 10, 20
	swappedByArgs(a, b)
	fmt.Printf("Main: a=%d, b=%d\n", a, b)
	swappedByAddr(&a, &b)
	fmt.Printf("Main: a=%d, b=%d\n", a, b)
}
