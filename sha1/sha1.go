package main

import "os"

// $ cat http.log.gz| gunzip | sha1sum
func sha1Sum(fileName string) (string, error){
	file, err := os.Open(filename)
	if err != nil {
		return "", nil
	}
}