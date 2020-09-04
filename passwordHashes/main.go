package main

import(
  "fmt"
  "io"
  "crypto/md5"
  "encoding/hex"
)

func PassHash(str string) string {
  // Create a hash interface   
	hi := md5.New()
  // Write str to interface for hashing 
  io.WriteString(hi, str)
  // Return hash and convert from hex to string 
  h := hex.EncodeToString(hi.Sum(nil))
  
  fmt.Printf(h)
  return h
}