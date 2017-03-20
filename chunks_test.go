package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestChunks(t *testing.T) {
	for mult := 0; mult < 15; mult++ {
		coeff := int(1 << uint(mult))

		for _, offset := range []int{-13, 0, 13} {
			n := coeff * (chunkSize - offset)
			t.Run(strconv.Itoa(n), func(t *testing.T) {
				b := nats(n)
				chks := chunks(b, chunkSize)

				want := n / chunkSize
				if n%chunkSize > 0 {
					want++
				}

				if got := len(chks); want != got {
					t.Errorf("want %v; got %v", want, got)
				}

				if want, got := b, join(chks); !reflect.DeepEqual(want, got) {
					t.Fatalf("want %v; got %v", want, got)
				}

			})
		}
	}
}

func nats(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(i)
	}
	return b
}

func join(chunks [][]byte) []byte {
	var b []byte
	for _, c := range chunks {
		b = append(b, c...)
	}
	return b
}
