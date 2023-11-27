// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package credentials

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func _getDeprecatedHasher() Hasher {
	return md5.New()
}

func _getHasher() Hasher {
	return sha256.New()
}

const (
	filenamePrefix      = "filename"
	encryptionKeyPrefix = "encryption"
)

type Hasher interface {
	Write(p []byte) (n int, err error)
	Sum(b []byte) []byte
}

func hashWith(hasher Hasher, key []byte) (string, error) {
	if _, err := hasher.Write(key); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// HashKeyToFilename creates a filename using the default hasher and provided key
// as input. It returns this filename and an error.
func HashKeyToFilename(key string) (string, error) {
	return HashKeyToFilenameWith(_getHasher(), key)
}

// HashKeyToFilenameWith creates a filename using the provided key as input and
// using the provided hasher.
func HashKeyToFilenameWith(hasher Hasher, key string) (string, error) {
	return hashWith(hasher, []byte(filenamePrefix+key))
}

// HashKeyToEncryptionKey creates an encryption key using a default hasher.
// It returns the created key and an error.
func HashKeyToEncryptionKey(key string) ([]byte, error) {
	return HashKeyToEncryptionKeyWith(_getHasher(), key)
}

// HashKeyToEncryptionKeyWith creates a 32 bytes long key from the provided
// key using the provided hasher.
func HashKeyToEncryptionKeyWith(hasher Hasher, key string) ([]byte, error) {
	h, err := hashWith(hasher, []byte(encryptionKeyPrefix+key))
	if err != nil {
		return nil, err
	}
	b := []byte(h)
	return b[:32], nil
}

// encrypt encrypts provided byte slice with AES using the encryption key.
func encrypt(data []byte, encryptionKey []byte) ([]byte, error) {
	f := func(hasher Hasher, data []byte, encryptionKey []byte) ([]byte, error) {
		block, err := aes.NewCipher(encryptionKey)
		if err != nil {
			return nil, err
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}
		nonce := make([]byte, gcm.NonceSize())
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			return nil, err
		}
		ciphertext := gcm.Seal(nonce, nonce, data, nil)
		return ciphertext, nil
	}

	if ret, err := f(_getHasher(), data, encryptionKey); err == nil {
		return ret, nil
	}

	ret, err := f(_getDeprecatedHasher(), data, encryptionKey)
	if err == nil {
		return ret, nil
	}
	return nil, err
}

// decrypt decrypts provided byte slice with AES using the encryptionKey.
func decrypt(data []byte, encryptionKey []byte) ([]byte, error) {
	f := func(hasher Hasher, data []byte, encryptionKey []byte) ([]byte, error) {
		block, err := aes.NewCipher(encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("unable tocreate new aes cipher: %w", err)
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, fmt.Errorf("unable to create new cipher gcm: %w", err)
		}
		nonceSize := gcm.NonceSize()
		if nonceSize > len(data) {
			return nil, fmt.Errorf("unable to decrypt credentials")
		}
		nonce, ciphertext := data[:nonceSize], data[nonceSize:]
		plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt: %w", err)
		}
		return plaintext, nil
	}

	if ret, err := f(_getHasher(), data, encryptionKey); err == nil {
		return ret, nil
	}

	ret, err := f(_getDeprecatedHasher(), data, encryptionKey)
	if err == nil {
		return ret, nil
	}
	return nil, err
}
