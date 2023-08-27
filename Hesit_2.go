package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  round.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludeGuards := rand.Intn(100)

  if eludeGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")

  } else {
    isHeistOn := false
    fmt.Println("Better luck next time!!")
  }

  openedVault = rand.Intn(100)
  
  if openedVault && eludeGuards >= 70{
    // isHeistOn := true
    fmt.Println("Grab and Go!")

  } else if isHeistOn {
    isHesitOn := false
    fmt.Println("It can't be done right now!!!")
  }

  fmt.Println("isHeistOn is currently: ", isHesitOn)

  leftSafely = rand.Intn(5)

  if isHesitOn {
    switch leftSafely 
    	fmt.Println("Its time to Move")
    case 0:
      isHeistOn := false
      fmt.Print("Looks like you tripped an alarm... run?")
    case 1:
      isHeistOn := false
      fmt.Print("When did they started adding Dogs in Vault?")
    case 2:
      isHeistOn := false
      fmt.Print("Looks like this fingerprint scanner won’t accept any fingerprint…")
    case 3:
      isHeistOn := false
      fmt.Print("Did I even pack the brulap bags?")
    case 4:
      isHeistOn := false
      fmt.Print("it ain't opening")

      amtStolen = 1000 + rand.Intn(1000000)
      fmt.Println("$", amtStolen)
	}
}
