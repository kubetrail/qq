package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// EncryptWithAesKey encrypts data using AES key
func EncryptWithAesKey(data, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		err := fmt.Errorf("could not create a new aes cipher: %w", err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		err := fmt.Errorf("could not create new gcm from cipher: %w", err)
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		err := fmt.Errorf("could not populate nonce: %w", err)
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// DecryptWithAesKey decrypts data using AES key
func DecryptWithAesKey(data, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		err := fmt.Errorf("could not create a new aes cipher: %w", err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		err := fmt.Errorf("could not create new gcm from cipher: %w", err)
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		err := fmt.Errorf("invalid cipher text, length less than nonce")
		return nil, err
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		err := fmt.Errorf("could not decrypt cipher text: %w", err)
		return nil, err
	}

	return plaintext, nil
}
