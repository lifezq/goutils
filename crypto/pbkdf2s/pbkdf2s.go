// Copyright 2016 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package pbkdf2s

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

const (
	block_size = 1 << 2
	block_iter = 1 << 12
)

func GenerateFromPassword(password []byte, klen int) ([]byte, error) {

	if len(password) < block_size {
		return nil, fmt.Errorf("Too short")
	}

	salt, err := generateSaltFromPassword(password)
	if err != nil {
		return nil, err
	}

	return append(salt, pbkdf2.Key(password, salt, block_iter, klen-len(salt), sha256.New)...), nil
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

func ScryptGenerateFromPassword(password []byte, klen int) ([]byte, error) {

	if len(password) < block_size {
		return nil, fmt.Errorf("Too short")
	}

	salt, err := generateSaltFromPassword(password)
	if err != nil {
		return nil, err
	}

	dk, err := scrypt.Key(password, salt, block_iter, block_size<<1, block_size>>2, klen-len(salt))
	if err != nil {
		return nil, err
	}

	return append(salt, dk...), nil
}

func ScryptCompareHashAndPassword(hashed []byte, password []byte) bool {

	if len(password) < block_size {
		return false
	}

	ilen := block_size + len(password)%block_size
	if len(hashed) <= ilen {
		return false
	}

	salt := hashed[:ilen]
	dk, err := scrypt.Key(password, salt, block_iter, block_size<<1, block_size>>2, len(hashed)-ilen)
	if err != nil {
		return false
	}

	phash := append(salt, dk...)
	if !bytes.Equal(hashed, phash) {
		return false
	}

	return true
}

func generateSaltFromPassword(password []byte) ([]byte, error) {

	ilen := block_size + len(password)%block_size
	salt := make([]byte, ilen)

	n, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	if n != ilen {
		return nil, fmt.Errorf("Length Not Match")
	}

	return salt, nil
}
