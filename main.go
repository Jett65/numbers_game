// Game Roles:
/* An array of random numbers is generated for each player.
Next an additional number is generated that is between 0 and half of the largest
number in the two arrays. Each players array is unknown to the other player. 
Play begins with player one(P1) inputs 5 numbers, each number corresponds to 
the index of one of the numbers in the openest array, P2 will then do the same. 
The chosen numbers must be from 0 to 4 and cannot be reputed. Next, both players
arrays will be reviled. Two signs are randomly generated for each player, each round 
that will determine whether to add or subtract the shift, Next, the first number in 
P1's array will be be added or subtracted from the shift number depending on the 
sign of the shift. The new number will be compare to the number in the openest 
array with the index of the first number P1 inputted. The openest's number will also be added
or subtracted depending on there generated sign. The player with the highest number in the 
comparison, after shifting them, will win. If P1 wins the first number in the round score
will be a 1 if P1 looses the number will be a 0. Play will continua this way for each number
in P1 array and a new sign will be generated after each comparison. When all five numbers will
have bean compare the binomial round score will be converted to a decimal number and that will
be P1 score. After the decimal score is calculated P2's array will be compare to P1's array using 
the numbers P2 inputted at the end the player with the highest decimal number wins the game.*/

package main

import (
    "fmt"
    "math/rand"
)

// Generates a random number in the range of min and max
func randint(min int, max int) int {
   var randint int = rand.Intn(max-min+1)+min  
   return randint
}

func main() {
    var player1_nums [5]int
    var player2_nums [5]int

    // Populates 2 arrays of each players numbers
    for i := 0; i < 5; i++ {
        p1_num := randint(0, 9)
        p2_num := randint(0, 9)
        
        player1_nums[i] = p1_num
        player2_nums[i] = p2_num
    }
    fmt.Printf("Player1: %v, Player2: %v\n", player1_nums, player2_nums)
}
