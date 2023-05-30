package main
import (
    "fmt"
    "math/rand"
    "time"
    "sort"
)

func main() {
    var arr[10] int
    var Searchfor int
    rand.Seed(time.Now().UnixNano())
    for i:=0; i < 5; i++ {
        arr[i] = rand.Intn(1000)
    }
    fmt.Println(arr)
    fmt.Println("Number to search for: ")
    fmt.Scanln(&Searchfor)
    arrSlice := arr[:]
    sort.Ints(arrSlice)
    fmt.Println(arrSlice)
    low := 0
    high := len(arrSlice) - 1
    for {
        mid := (low + high) / 2
        guess := arrSlice[mid]
        if low <= high {
            if Searchfor < guess  {
                high = mid - 1
            } else if Searchfor == guess {
                fmt.Println("Found it at: ", mid)
                break
            } else {
                low = mid + 1
            }
        } else {
            fmt.Println("Nothing found")
            break
        }
    }
}

