package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/howeyc/gopass"
	"github.com/json-iterator/go"
)

type GHRepo struct {
	Stars int `json:"stargazers_count"`
}

func main() {
	readme := "README.md"
	output := "parser/README.md"

	lines, err := readLines(readme)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your github username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	passB, _ := gopass.GetPasswdMasked()
	pass := string(passB)

	regStrFull := "\\* \\[[-_A-Za-z0-9\\/]+\\]\\(https:\\/\\/github\\.com\\/[-_A-Za-z]+\\/[-_A-Za-z]+\\/?\\)"
	regLink := regexp.MustCompile("https:\\/\\/github\\.com\\/([-_A-Za-z]+\\/[-_A-Za-z]+)\\/?")

	total := 0
	for i, line := range lines {
		match, _ := regexp.MatchString(regStrFull, line)
		if match {
			total = total + 1
			repo := regLink.FindStringSubmatch(line)
			getStarsAndWriteRes(repo[1], lines, i, username, pass, total)
		}
	}

	os.Remove(output)
	writeLines(lines, output)
}

func getStarsAndWriteRes(repo string, lines []string, i int, username, pass string, total int) {
	var client http.Client
	apiUrl := "https://api.github.com/repos/"
	url := apiUrl + repo
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, pass)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("STATUS: %d\n%s\n%s\n", resp.StatusCode, string(bodyBytes), url)
		return
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	ghRepo := GHRepo{}
	json.Unmarshal(bodyBytes, &ghRepo)

	lines[i] = fmt.Sprintf("%s **⭐%d**", lines[i], ghRepo.Stars)
	fmt.Println(total, lines[i])
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
