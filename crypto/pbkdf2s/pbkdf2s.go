// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package pbkdf2

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	block_size = 1 << 3
	block_iter = 1 << 12
)

func GenerateFromPassword(password []byte, klen int) ([]byte, error) {

	if len(password) < block_size {
		return []byte{}, fmt.Errorf("Too short")
	}

	ilen := block_size + len(password)%block_size
	salt := make([]byte, ilen)

	n, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return []byte{}, err
	}

	if n != ilen {
		return []byte{}, fmt.Errorf("Length Not Match")
	}

	return append(salt, pbkdf2.Key(password, salt, block_iter, klen, sha256.New)...), nil
}

func CompareHashAndPassword(hashed []byte, password []byte) bool {

	if len(password) < block_size {
		return false
	}

	ilen := block_size + len(password)%block_size
	if len(hashed) <= ilen {
		return false
	}

	salt := hashed[:ilen]
	phash := append(salt, pbkdf2.Key(password, salt, block_iter, len(hashed)-ilen, sha256.New)...)
	if !bytes.Equal(hashed, phash) {
		return false
	}

	return true
}
