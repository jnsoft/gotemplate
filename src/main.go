package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	TEXT = "Hello, world!"
 )

func main() {
  // pat := os.Getenv("PAT")
  verbose := flag.Bool("v", false, "Enable verbose logging")

  flag.Parse()
  // args := flag.Args()

  if *verbose {
		log.Print("Verbose mode enabled")
  }

  fmt.Println(TEXT)
}
