# ArchiveMe

archiveme is a command-line tool that archive folder in tar, than compress into gzip and ecode it using AES. Should works on Windows, Linux and macOS.

Send archive file and your password to other device in separate ways. 

## ArchiveMe is One-Button-Application.

### Usage:

Just do the archive with default password (0000)
```
archiveme ./directory/to/be/archived
archiveme ./address/to/file
archiveme -f /usr/address/to/file
```

Using password
```
archiveme -f /usr/address/to/file -p myPassword
```

Set other file type (example file.archived)
```
archiveme --type .archived
```

Unzip
```
archiveme ./folsername.nau
```

Unzip with password
```
archiveme -p myPassword ./file.nau
```

Unzip with specific folder for result
```
archiveme -d ./put.result/here -p myPassword ./file.nau
```

### Build examples:

64bit

Linux:
`GOOS=linux GOARCH=amd64 go build -o bin/archiveme ./cmd/encoder/main.go`

Windows:
`GOOS=windows GOARCH=amd64 go build -o bin/archiveme ./cmd/encoder/main.go`

MacOS:
`GOOS=darwin GOARCH=amd64 go build -o bin/archiveme ./cmd/encoder/main.go`