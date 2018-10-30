package main

import "fmt"

//interface{} 能夠存放任何事物

func main() {
	var mixed interface{}

	mixed = 123
	fmt.Println(mixed) // 輸出：123

	mixed = "Moon, Dalan!"
	fmt.Println(mixed) // 輸出：Moon, Dalan!

	mixed = []string{"A", "B", "C"}
	fmt.Println(mixed) // 輸出：["A", "B", "C"]
}
