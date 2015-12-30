package main

import(
  "fmt"
  "os"
  "github.com/sheepkiller/gotopt"
)

func main() { 
        fmt.Println(gotopt.GetTOPT(os.Getenv("TOPT_SECRET"), 6, "sha1"))
}


