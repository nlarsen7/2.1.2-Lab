//Nicholas Larsen
//February 20,2020
//This takes a 4-digit number and adds each digit together until it is a single digit number
package main

import "fmt"

func main() {
  var number int

  fmt.Println("enter a 4-digit number")
  fmt.Scanln(&number)
  //gets the user's number

for number>999{
  fmt.Println((number/1000) + ((number%1000)/100) + ((number%100)/10) +((number%100)%10))
  //prints out the number's digits added together
 
  number= (number/1000) + ((number%1000)/100) + ((number%100)/10) +((number%100)%10)
  //makes the variable that new number
  
  if number<10{
  fmt.Println("your answer is",number)
  //if the sum is less than 10 it ends here
  
 }else if number>10{
    fmt.Println("Your answer is",number/10 + number%10)
    //adds the digits again and gives you your answer
  }
  
}
}
