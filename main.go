//name
//date
//description of program

package main

import "fmt"

func main() {
  //declares input var as string
  var input string
  //prints whats your name
  fmt.Println("what is your name")
  //stores response as input
  fmt.Scanln(&input)
  //prints your names is input
  fmt.Println("your name is" , input)
  //prints whats your favorite food
  fmt.Println("whats your favorite food")
  //stores response as input
  fmt.Scanln(&input)
  //prints favorite food is input
  fmt.Println("favorite food is " , input)
  //print //prints whats your age
  fmt.Println("whats your age")
  //stores response as input
  fmt.Scanln(&input)
  //prints your age is input
  fmt.Println("your age is" , input)
}