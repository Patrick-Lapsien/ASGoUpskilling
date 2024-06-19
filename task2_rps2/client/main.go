package main

import ("fmt"
        "bufio"
        "strings"
        "os"
        "math/rand"

        pb "github.com/Patrick-Lapsien/ASGoUpskilling/task2_rps2/client"
        "google.golang.org/grpc"
        )

func main(){
    choiceString:= []string{"rock", "paper", "scissor"}
    fmt.Println("Welcome to Rock, Paper, Scissors!. Let's start.")
    fmt.Println("Which do you pick? Rock, Paper or Scissor?")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    input = strings.ToLower(input)
    var playerChoice string
    switch{
        case strings.Contains(input, choiceString[0]): playerChoice= choiceString[0]
        case strings.Contains(input, choiceString[1]): playerChoice= choiceString[1]
        case strings.Contains(input, choiceString[2]): playerChoice= choiceString[2]
        default: fmt.Println("you didn't pick any of the available choices")
    }
    fmt.Println("you choose "+ playerChoice)
    randomNumber := rand.Intn(len(choiceString) )
    computerChoice := choiceString[randomNumber]
    fmt.Println("Computer picked " + computerChoice)
    var ergebnis int
    switch{
        case computerChoice== playerChoice: ergebnis=0
        case computerChoice==choiceString[0] && playerChoice==choiceString[1]|| computerChoice==choiceString[1] && playerChoice==choiceString[2] || computerChoice==choiceString[2] && playerChoice==choiceString[0]: ergebnis=1
        case computerChoice==choiceString[1] && playerChoice==choiceString[0]|| computerChoice==choiceString[2] && playerChoice==choiceString[1] || computerChoice==choiceString[0] && playerChoice==choiceString[2]: ergebnis=-1
        default: fmt.Println("result not yet implemented")
                ergebnis=999
    }
    switch ergebnis{
        case 0 : fmt.Println("draw")
        case 1 : fmt.Println("player wins")
        case -1 : fmt.Println("computer  wins")
        default: fmt.Println("result not yet implemented 2")
    }

}