package main

import "fmt"

func main() {
  var number int

  fmt.Println("enter a 4-digit number")
  fmt.Scanln(&number)

for number>999{
  fmt.Println((number/1000) + ((number%1000)/100) + ((number%100)/10) +((number%100)%10))
 
  number= (number/1000) + ((number%1000)/100) + ((number%100)/10) +((number%100)%10)
  
  if number<10{
  fmt.Println("your answer is",number)
  
 }else if number>10{
    fmt.Println("Your answer is",number/10 + number%10)
  }
  
}
}
