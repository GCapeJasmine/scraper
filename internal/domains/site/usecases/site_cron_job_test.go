package usecases

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetListSite(t *testing.T) {
	filePath := "sites.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var sites []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sites = append(sites, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %v", err)
	}
	fmt.Println(sites)
}
