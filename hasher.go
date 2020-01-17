package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

type HashFunc struct {
	Name     string
	HashFunc crypto.Hash
}

func doHash(hasher hash.Hash, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := io.Copy(hasher, file); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func doFile(filename string) {
	// Many of these are deprecated and for whatever reason, not included in the default
	// library. They exist in x/crypt, google it for more info.
	algs := []HashFunc{
		{"MD4", crypto.MD4},
		{"MD5", crypto.MD5},
		{"SHA1", crypto.SHA1},
		{"SHA224", crypto.SHA224},
		{"SHA256", crypto.SHA256},
		{"SHA384", crypto.SHA384},
		{"SHA512", crypto.SHA512},
		{"MD5SHA1", crypto.MD5SHA1},
		{"RIPEMD160", crypto.RIPEMD160},
		{"SHA3_224", crypto.SHA3_224},
		{"SHA3_256", crypto.SHA3_256},
		{"SHA3_384", crypto.SHA3_384},
		{"SHA512_224", crypto.SHA512_224},
		{"SHA512_256", crypto.SHA512_256},
		{"BLAKE2s_256", crypto.BLAKE2s_256},
		{"BLAKE2b_256", crypto.BLAKE2b_256},
		{"BLAKE2b_384", crypto.BLAKE2b_384},
		{"BLAKE2b_512", crypto.BLAKE2b_512},
	}

	fmt.Println(filename)
	for _, alg := range algs {
		if alg.HashFunc.Available() {
			h := alg.HashFunc.New()
			hash := doHash(h, filename)
			fmt.Printf("    %-16s: %s\n", alg.Name, hash)
		}
	}

	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No filenames given, aborting")
	}

	for _, filename := range os.Args[1:] {
		doFile(filename)
	}
}
