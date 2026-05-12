package main 

import "fmt"

func main () {
	fmt.Print(RemoveChar("test"))
}

func RemoveChar(word string) string {
  var resultString string
  
  for i, e := range word {
    if i == 0 {
		continue
    } else if i == len(word) - 1 {
		continue
    }
    resultString += string(e)
  }
  
  return resultString
}