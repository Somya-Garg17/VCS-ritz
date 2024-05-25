package utility

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// Remove comments when using in your code
// comments are for you to understand
// You can blindly copy but having some idea of what you are doing gives some confidence
// The same goes for gpt, use any code snippet(from anywhere) which works but atleast try to understand for your own good 
func Sha1(content string) string {
  // returns a type implementing Hash interface
  // golang standard library implements hashing in very general way
  // So any hash pkg (sha256 etc) will have method to return a type implementing Hash interface and start a hash procees
  // The Hash interface has io.Writer interface embedded and so write method can be invoked ontype satisfying Hash interface, it will add/write data to running hash method
  // So in this case sha1.new() starts a hash proccess and return Hash interface
  // now you can use write method on this interface to add your data which will get hashed with the running hash process
	h := sha1.New()

  // Writing our data to running Hash by using io.WriteString
  // io.WriteString expects type implementing io.Writer
  // We are passing type implementing Hash interface
  // but Hash interface has io.Writer embedded
  // io.WriteString will write data to the running hash procees
	io.WriteString(h, content)

  // The Sum method on Hash interface returns the hashed data
  // You can also pass a offset to the hash(a cont value to be added to hash)
  // We are passing nil as we want 0 offset
  // Finally h.Sum(nil) returns []byte
  // Remember each entry in []byte is ASCII encoding
  // But we want 40 hex characters
  // So we use hex.EncodeString to Encode ASCII Characters to Hexadecimal
  // h.Sum(nil) will return []byte of length 20
  // hashedContent is 40 hex characters long
  hashedContent := hex.EncodeToString(h.Sum(nil))
  return hashedContent
}

// []byte is very similar to string
// []byte stores ASCII encoding of each character whereas string consist of chars themselves
// b := []byte{97, 98, 99, 100}
// s := "ABCD"
// b and s can be typecasted into each other
// h := hex.EncodeToString(b)
// fmt.Println(h) => 61626364
// 61 = 6 * 16 + 1 = 97
