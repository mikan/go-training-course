// Copyright 2016 mikan.

// Package crypt provides some utilities for simple crypt & decrypt.
// Limitation 1: Max length is 32 for both password and plain text.
// Limitation 2: Cannot use "*" for plain text because it uses for padding.
package main

import (
	"crypto/aes"
	"errors"
	"strings"
)

const paddingChar = "*"

var initialChars = strings.Repeat(paddingChar, 32) // 32 chars

// Encrypt encrypts plaint text with password.
func Encrypt(plainText, password string) ([]byte, error) {
	// Length check
	if len(plainText) > len(initialChars) {
		return nil, errors.New("Too long plain text.")
	}
	if len(password) > len(initialChars) {
		return nil, errors.New("Too long password.")
	}

	// Setup cipher
	pw := []byte(initialChars)
	bPassword := []byte(password)
	for i := 0; i < len(bPassword); i++ {
		pw[i] = bPassword[i]
	}
	cipher, err := aes.NewCipher(pw)
	if err != nil {
		return nil, err
	}

	// Setup plain text
	pt := []byte(initialChars)
	bPlain := []byte(plainText)
	for i := 0; i < len(bPlain); i++ {
		pt[i] = bPlain[i]
	}

	// Encrypt
	encrypted := make([]byte, aes.BlockSize)
	cipher.Encrypt(encrypted, pt)
	return encrypted, nil
}

// Decrypt decrypts source with password.
func Decrypt(encrypted []byte, password string) (string, error) {
	// Length check
	if len(password) > len(initialChars) {
		return "", errors.New("Too long password.")
	}

	// Setup cipher
	pw := []byte(initialChars)
	bPassword := []byte(password)
	for i := 0; i < len(bPassword); i++ {
		pw[i] = bPassword[i]
	}
	cipher, err := aes.NewCipher(pw)
	if err != nil {
		return "", err
	}

	// Decrypt
	decrypted := make([]byte, aes.BlockSize)
	cipher.Decrypt(decrypted, encrypted)
	return strings.Replace(string(decrypted), paddingChar, "", -1), nil
}
