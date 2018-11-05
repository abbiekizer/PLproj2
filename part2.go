package main

import "fmt"
//import "time"
//import "io/ioutil"
import "bufio"
import "os"
//import "log"
import "strconv"

func main() {
  c1 := make(chan int)
  c2 := make(chan string)
}

func factorNums(c1 chan int) {
  file, err := os.Open("numers.txt")
  outfile, err := os.Create("out.txt")
  if err != nil { //check for error reading file
        fmt.Println("File reading error", err)
        return
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    //printFactors(scanner, "out.txt")
    var num int = strconv.Atoi(scanner)
    fmt.Fprintf(outfile, " (")
    for i := 2; i <= num; i++ {
      for num%i == 0 {
      fmt.Fprintf(outfile, "%d", i)
      if num == i {
       num = num/i
      } else {
        fmt.Fprintf(outfile,"*")
        num = num/i
      }
  }
    }
    fmt.Fprintf(outfile,") ")
  }

}

/*func printFactors(num int, outfile string) {
     fmt.Fprintf(outfile, " (")
     for i := 2; i <= num; i++ {
     	 for num%i == 0 {
	     fmt.Fprintf(outfile, "%d", i)
	     if num == i {
	     	num = num/i
	     } else {
	       fmt.Fprintf(outfile,"*")
	       num = num/i
	     }
	 }
     }
     fmt.Fprintf(outfile,") ")
}*/
