package frame_test

import (
	"github.com/windflowed/tcp-server-practicedemo/frame"
	"bytes"
	"encoding/binary"
	"testing"
)

func TestEncode(t *testing.T) {
	codec := frame.NewMyFrameCodec()
	buf := make([]byte, 0, 128)
	rw := bytes.NewBuffer(buf)

	err := codec.Encode(rw, []byte("hello world"))
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	var totalLen int32
	err = binary.Read(rw, binary.BigEndian, &totalLen)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	if totalLen != 15 {
		t.Errorf("want 15, actual %d", totalLen)
	}

	left := rw.Bytes()
	if string(left) != "hello world" {
		t.Errorf("want hello world, actual %s", string(left))
	}
}

func TestDecode(t *testing.T) {
    codec := frame.NewMyFrameCodec()
	data := []byte{0x0, 0x0, 0x0, 0xf, 'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}

	payload, err := codec.Decode(bytes.NewReader(data))
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	if string(payload) != "hello world" {
		t.Errorf("want hello world, actual %s", string(payload))
	}
}

