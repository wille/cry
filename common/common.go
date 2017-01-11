package common

// Extensions to encrypt
var Extensions = [...]string{
	"png",
	"jpg",
	"jpeg",
	"tif",
}

const (
	// LockedExtension to append to file name when encrypted
	LockedExtension string = ".locked"

	// ProcessMax X files, then stop
	ProcessMax int = 10
)
