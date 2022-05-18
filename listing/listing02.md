 
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return 
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
   	x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}

Ответ: программа выводит 2 1, так как в первом случае мы работаем с именнованым возвращаемым значением, и функция defer срабатывает после того, как x увеличивается. Во втором случае - функция defer работает когда вычисляется x.