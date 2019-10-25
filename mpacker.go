package mpacker

import (
  "github.com/ugorji/go/codec"
)

type  Buffer struct {
  data []byte
  msg  map[string]string
}


func FromBytes(buf []byte) *Buffer {
  return &Buffer{buf, make(map[string]string)}
}

func FromDict(d map[string]string) *Buffer {
  return &Buffer{make([]byte, 0), d}
}

func (b *Buffer) Bytes() []byte {
  return b.data
}


func (b *Buffer) Msg() map[string]string {
  return b.msg
}

func (b *Buffer) Encode() error {
  enc := codec.NewEncoderBytes(&b.data, new(codec.MsgpackHandle))
  err := enc.Encode(&b.msg)
  return err
}

func (b *Buffer) Decode() error {
  dec := codec.NewDecoderBytes(b.data, new(codec.MsgpackHandle))
  err := dec.Decode(&b.msg)
  return err
}
