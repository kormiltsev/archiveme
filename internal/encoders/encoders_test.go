package encoders

import (
	"bytes"
	"testing"
)

func TestEncrypto(t *testing.T) {
	// Test data
	plaintext := []byte("This is the plaintext to be encrypted.")
	keystring := []byte("my-secret-key")

	// Call the Encrypto function to encrypt the plaintext
	ciphertext, err := Encrypto(plaintext, keystring)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Test that the ciphertext is not equal to the plaintext
	if bytes.Equal(ciphertext, plaintext) {
		t.Errorf("Ciphertext is equal to plaintext")
	}

	// Test that decrypting the ciphertext using the same key produces the original plaintext
	decryptedText, err := Decrypto(ciphertext, keystring)
	if err != nil {
		t.Fatalf("Unexpected error during decryption: %v", err)
	}

	if !bytes.Equal(decryptedText, plaintext) {
		t.Errorf("Decrypted text does not match the original plaintext")
	}
}
