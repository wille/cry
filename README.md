# go-cry

This project was written to show how easy it is to create extremely malicious code.

Ransomware is designed to take your most loved files hostage demanding large amounts of money to unlock them.

Clone of [native-tear](https://github.com/redpois0n/native-tear/) written in Go which is a clone of [hidden-tear](https://github.com/utkusen/hidden-tear/)

## Building

go-cry consists of two parts, a webserver and the client software.
Output files will be placed in `./bin/`

Built with linker flags `-w -s` to minimize file size. To further reduce the size of Go binaries, please see the [UPX project](https://upx.github.io/)

### Building client and webserver
```
$ make
```

Will create the files
- `./bin/web[.exe]`
- `./bin/cry[.exe]`


### Building client for all common operating systems and architectures
```
$ make all
```

Will create the files
- `./bin/windows_amd64.exe`
- `./bin/windows_x86.exe`
- `./bin/linux_amd64`
- `./bin/linux_x86`
- `./bin/macos` (amd64)

### Cleaning

Will remove all files in the bin directory
```
$ make clean
```

# Configuring

### Web server
See [web/web.go](web/web.go) and modify the constant values. They are commented and straight forward.

### Client
See [config.go](config.go) and modify the constant values.

If modifying the RSA key size variable `Bits`, please see `EncryptedHeaderSize`. RSA ciphertext length changes depending on key size used and it is not calculated at runtime.

# This program does not

- Demand any money from the user. It simply encrypts the amount of files specified in [config.go](config.go) 
constant `ProcessMax` and sends it to the server. Encrypt your files and store your encryption key on your 
server.

# Legal Warning

go-cry may be used only for Educational Purposes. Do not use it as a ransomware! You could go to jail on obstruction of justice charges just for running hidden tear, even though you are innocent.
