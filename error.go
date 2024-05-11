// in the code i implement panic,recovers and defer
package main
import (
"fmt"
)

func badCall() {
panic("bad end")
}
func test() {
defer func() {
if e :=recover(); e !=nil {
fmt.Printf("panicking%\r\n",e);
}
}()
badCall()
fmt.Printf("After badc call\r\n");
}

func main() {
fmt.Printf("calling test\r\n");
test()
fmt.Println("Test completed\r\n");
}
