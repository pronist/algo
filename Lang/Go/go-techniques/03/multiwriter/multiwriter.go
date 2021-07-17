package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	tmp, _ := ioutil.TempFile(os.TempDir(), "tmp")
	defer tmp.Close()

	hash := sha256.New()
	w := io.MultiWriter(tmp, hash)

	written, _ := io.Copy(w, os.Stdin)
	fmt.Printf("Wrote %d bytes to %s \nSHA256: %x\n", written, tmp.Name(), hash.Sum(nil))
}
