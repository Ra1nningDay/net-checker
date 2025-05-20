package main

import "fmt"

func ping(host string) (string, error) {
	if host == "" {
		return "", fmt.Errorf("host is empty")
	}

	return host, nil
}

func main() {
	host := ""
	result, err := ping(host)
	if err != nil {
		fmt.Printf("Error %s: %v", host, err)
		return
	}
	fmt.Printf("Ping: %s: \n%s\n", host, result)
}
