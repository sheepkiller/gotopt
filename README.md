gotopt: A basic package to generate/validate TOPT (RFC6238)
==========
[![Build Status](https://travis-ci.org/sheepkiller/gotopt.png?branch=master)](https://travis-ci.org/sheepkiller/gotopt)
[![GoDoc](https://godoc.org/github.com/sheepkiller/gotopt?status.svg)](https://godoc.org/github.com/sheepkiller/gotopt)

# Usage

```Go
package main

import(
   "fmt"
   "github.com/sheepkiller/gotopt"
   "os"
   "time"
)

func main() {
   token, remain, err := gotopt.GetTOPT("AAAAAAAAAAAAAAAA", 6, "sha1")
   if err != nil {
      fmt.Printf("Error: %s\n", err.Error())
      os.Exit(1)
   }
   fmt.Printf("TOPT: %s will expire in %d seconds\n", token, remain)
   valid, err:= gotopt.ValidateTOPT("AAAAAAAAAAAAAAAA", 6, "sha1", token, 0)
   if valid {
      fmt.Println("token is valid")
   }
   time.Sleep(30 * time.Second)
   valid, err= gotopt.ValidateTOPT("AAAAAAAAAAAAAAAA", 6, "sha1", token, 0)
   if valid {
      fmt.Println("token is valid")
   } else {
      fmt.Println("token is invalid")
   }
   valid, err= gotopt.ValidateTOPT("AAAAAAAAAAAAAAAA", 6, "sha1", token, 1)
   if valid {
      fmt.Println("token is valid")
   } else {
      fmt.Println("token is invalid")
   }
}
```

# References
- [Wikipedia Article](https://en.wikipedia.org/wiki/Time-based_One-time_Password_Algorithm)
- [RFC6238](https://tools.ietf.org/html/rfc6238)

This work is based on:
- [PHP implemention from @jds13](https://github.com/jds13/RFC6328-PHP)
- [Python implementation from @antonio-fr](https://github.com/antonio-fr/OTPpy)

