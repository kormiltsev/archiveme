package app

// print logo and some info
func About() string {
	return `
                _     _                          
  __ _ _ __ ___| |__ (_)_   _____ _ __ ___   ___ 
 / _' | '__/ __| '_ \| \ \ / / _ \ '_ ' _ \ / _ \
| (_| | | | (__| | | | |\ V /  __/ | | | | |  __/
 \__,_|_|  \___|_| |_|_| \_/ \___|_| |_| |_|\___|
   $ archiveme /address/dir/or/file                  // to encode
   $ archiveme -f /usr/address/to/file               // to encode
   $ archiveme -p myPassword ./filename              // with password
   $ archiveme --type .mytype ./file                 // set suffix
   $ archiveme --type _suffix ./file                 // set suffix

   $ archiveme ./filename.archiveme                  // to decode
   $ archiveme -p myPassword ./file.archiveme        // with password
   $ archiveme -d ./put/result/here ./file.archiveme // put result here
   `
}
