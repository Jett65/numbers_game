package main

import (
    "fmt"
    "math/rand"
)

func randint(min int, max int) int {
   var randint int = rand.Intn(max-min+1)+min  
   return randint
}

func main() {
    var player1_nums [5]int
    var player2_nums [5]int
    for i := 0; i < 5; i++ {
        p1_num := randint(0, 9)
        p2_num := randint(0, 9)
        
        player1_nums[i] = p1_num
        player2_nums[i] = p2_num
    }
    fmt.Printf("Player1: %v, Player2: %v\n", player1_nums, player2_nums)
}
