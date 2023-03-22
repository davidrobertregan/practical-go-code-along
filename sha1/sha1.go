package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

// $ cat http.log.gz| gunzip | sha1sum
func sha1Sum(fileName string) (string, error){
	// idiom: acquire a resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer r.Close()

	// io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}