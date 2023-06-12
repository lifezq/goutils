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
	blockSize = 1 << 2
	blockIter = 1 << 12
)

func GenerateFromPassword(password []byte, klen int) ([]byte, error) {
	if len(password) < blockSize {
		return nil, fmt.Errorf("Too short")
	}

	salt, err := generateSaltFromPassword(password)
	if err != nil {
		return nil, err
	}

	return append(salt, pbkdf2.Key(password, salt, blockIter, klen-len(salt), sha256.New)...), nil
}

func CompareHashAndPassword(hashed, password []byte) bool {
	if len(password) < blockSize {
		return false
	}

	ilen := blockSize + len(password)%blockSize
	if len(hashed) <= ilen {
		return false
	}

	salt := hashed[:ilen]
	phash := append(salt, pbkdf2.Key(password, salt, blockIter, //nolint:gocritic // ignore
		len(hashed)-ilen, sha256.New)...)
	return bytes.Equal(hashed, phash)
}

func ScryptGenerateFromPassword(password []byte, klen int) ([]byte, error) {
	if len(password) < blockSize {
		return nil, fmt.Errorf("Too short")
	}

	salt, err := generateSaltFromPassword(password)
	if err != nil {
		return nil, err
	}

	dk, err := scrypt.Key(password, salt, blockIter, blockSize<<1, blockSize>>2, klen-len(salt))
	if err != nil {
		return nil, err
	}

	return append(salt, dk...), nil
}

func ScryptCompareHashAndPassword(hashed, password []byte) bool {
	if len(password) < blockSize {
		return false
	}

	ilen := blockSize + len(password)%blockSize
	if len(hashed) <= ilen {
		return false
	}

	salt := hashed[:ilen]
	dk, err := scrypt.Key(password, salt, blockIter, blockSize<<1, blockSize>>2, len(hashed)-ilen)
	if err != nil {
		return false
	}

	phash := append(salt, dk...) //nolint:gocritic // ignore
	return bytes.Equal(hashed, phash)
}

func generateSaltFromPassword(password []byte) ([]byte, error) {
	ilen := blockSize + len(password)%blockSize
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
