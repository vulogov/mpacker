package mpacker

import "fmt"
import "testing"

func TestEncode(t *testing.T) {
  b := FromDict(map[string]string{"message": "Hi!"})
  err := b.Encode()
  if err != nil {
    t.Errorf("Error during msgpack encoding")
  }
  for _, v := range b.Bytes() {
    fmt.Printf("%X", v)
  }
  fmt.Printf("\n")
}

func TestDecode(t *testing.T) {
  b := FromDict(map[string]string{"message": "Hi!"})
  err := b.Encode()
  if err != nil {
    t.Errorf("Error during msgpack encoding")
  }
  b2 := FromBytes(b.Bytes())
  err = b2.Decode()
  if err != nil {
    t.Errorf("Error during msgpack decoding")
  }
  fmt.Println(b2.Msg())
}
