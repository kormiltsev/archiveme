/*
# ArchiveMe

### ArchiveMe is Ote-Button-Application.
Usage:

Just do the archive with default password (0000)

	archiveme ./directory/to/be/archived
	archiveme ./address/to/file
	archiveme -f /usr/address/to/file

Using password

	archiveme -f /usr/address/to/file -p myPassword

Unzip

	archiveme ./folsername.nau

Unzip with password

	archiveme -p myPassword ./file.nau

Unzip with specific folder for result

	archiveme -d ./put.result/here -p myPassword ./file.nau
	archiveme --type .archiveme
*/
package main
