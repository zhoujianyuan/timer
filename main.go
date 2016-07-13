// timer project main.go
package main

import (
            "fmt"
                "time"
                    "timer/timemanage"
       )

var newTime *timemanage.Timer

func printHello() {
        newTime.NewTimer(time.Second, printHello)
}

func printWorld() {
        fmt.Println("world")
                newTime.NewTimer(time.Second*2, printWorld)
}
func main() {
        fmt.Println("Hello World!")
                newTime = timemanage.New(time.Millisecond * 100)
                    newTime.NewTimer(time.Second, printHello)
                    // timer project main.go
                    package main

                    import (
                                "fmt"
                                    "time"
                                        "timer/timemanage"
                           )

                    var newTime *timemanage.Timer

                    func printHello() {
                            newTime.NewTimer(time.Second, printHello)
                    }

        func printWorld() {
                fmt.Println("world")
                        newTime.NewTimer(time.Second*2, printWorld)
        }
        func main() {
                fmt.Println("Hello World!")
                        newTime = timemanage.New(time.Millisecond * 100)
                            newTime.NewTimer(time.Second, printHello)
                                newTime.NewTimer(time.Second*2, printWorld)
                                    newTime.Start()
        }

