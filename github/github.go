package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println(githubInfo("davidrobertregan"))
}

// func demo() {
// 	resp, err := http.Get("https://api.github.com/users/davidrobertregan")

// 	if err != nil {
// 		log.Fatalf("error: %s", err)
// 		/*
// 			log.Pringf("error: %s", err)
// 			os.Exit(1)
// 		*/
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("error: %s", resp.Status)
// 	}
// 	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
// 	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
// 	// 	log.Fatalf("error: can't copy - %s", err)
// 	// }
// 	var r Reply 
// 	dec := json.NewDecoder(resp.Body)
// 	if err := dec.Decode(&r); err != nil {
// 		log.Fatalf("error: can't decode - %s", err)
// 	}
// 	fmt.Printf("%#v\n", r)
// }

func githubInfo(login string) (string, int, error) {
	// "string" + string_variable
	url := fmt.Sprintf("https://api.github.com/users/%v", url.PathEscape(login))
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf(resp.Status)
	}
	
	var r struct {
		Name string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}

/*
	type Reply struct {
		Name         string
		// Public_Repos int
		NumRepos int `json:"public_repos`
	}
*/

/*
	json -> io.Reader -> Go: json.Decoder
	json -> 
*/