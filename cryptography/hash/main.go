package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " <filepath>")
	fmt.Println("Example: " + os.Args[0] + " document.txt")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	filename := checkArgs()

	// Get bytes from file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Hash the file and output results
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))

	// Hashing large file
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create new hasher, which is a writer interface
	hasher := md5.New()

	// Default buffer size for copying is 32*1024 or 32kb per copy
	// Use io.CopyBuffer() if you want to specify the buffer to use
	// It will write 32kb at a time to the digest/hash until EOF
	// The hasher implements a Write() function making it satisfy
	// the writer interface. The Write() function performs the digest
	// at the time the data is copied/written to it. It digests
	// and processes the hash one chunk at a time as it is received.
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}

	// Now get the final sum or checksum.
	// We pass nil to the Sum() function because
	// we already copied the bytes via the Copy to the
	// writer interface and don't need to pass any new bytes
	checksum := hasher.Sum(nil)

	fmt.Printf("Md5 checksum: %x\n", checksum)
}
