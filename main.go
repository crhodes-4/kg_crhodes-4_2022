package main

import (
	"fmt"
	"os"
  //"strconv"
)

func main() {

  var arr []string
  arr = os.Args[1:] //get numbers input and store in array ex: [3, 25, 209]

  //create lookup array
  arr_str := [10]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

  outp := "" //output variable

  for i, num := range arr { //outer loop iterates through arr
    for j, dig := range num { //inner loop iterates through string ex: "209"

      //if chain to check which char is found, concatenate to output str accordingly
      if dig == '0' {
        outp += arr_str[0]
      } else if dig == '1' {
          outp += arr_str[1]
      } else if dig == '2' {
          outp += arr_str[2]
      } else if dig == '3' {
          outp += arr_str[3]
      } else if dig == '4' {
          outp += arr_str[4]
      } else if dig == '5' {
          outp += arr_str[5]
      } else if dig == '6' {
          outp += arr_str[6]
      } else if dig == '7' {
          outp += arr_str[7]
      } else if dig == '8' {
          outp += arr_str[8]
      } else if dig == '9' {
          outp += arr_str[9]
      }
      _ = j
    }

    if i != len(arr)-1 { //this is so if i is at last elem, it won't add "," to string outp
        outp += ","
      }
  }

  fmt.Println(outp) //print outp to stdout

}
