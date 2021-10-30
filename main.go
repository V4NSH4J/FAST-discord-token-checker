package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

func main() {
	color.Red("\u2584\u2584\u2584\u2588\u2588\u2588\u2588\u2588\u2593 \u2592\u2588\u2588\u2588\u2588\u2588   \u2588\u2588 \u2584\u2588\u2580\u2593\u2588\u2588\u2588\u2588\u2588  \u2588\u2588\u2588\u2584    \u2588   \u2588\u2588\u2588\u2588\u2588\u2588   \u2584\u2588\u2588\u2588\u2588  \u2592\u2588\u2588\u2588\u2588\u2588  \r\n\u2593  \u2588\u2588\u2592 \u2593\u2592\u2592\u2588\u2588\u2592  \u2588\u2588\u2592 \u2588\u2588\u2584\u2588\u2592 \u2593\u2588   \u2580  \u2588\u2588 \u2580\u2588   \u2588 \u2592\u2588\u2588    \u2592  \u2588\u2588\u2592 \u2580\u2588\u2592\u2592\u2588\u2588\u2592  \u2588\u2588\u2592\r\n\u2592 \u2593\u2588\u2588\u2591 \u2592\u2591\u2592\u2588\u2588\u2591  \u2588\u2588\u2592\u2593\u2588\u2588\u2588\u2584\u2591 \u2592\u2588\u2588\u2588   \u2593\u2588\u2588  \u2580\u2588 \u2588\u2588\u2592\u2591 \u2593\u2588\u2588\u2584   \u2592\u2588\u2588\u2591\u2584\u2584\u2584\u2591\u2592\u2588\u2588\u2591  \u2588\u2588\u2592\r\n\u2591 \u2593\u2588\u2588\u2593 \u2591 \u2592\u2588\u2588   \u2588\u2588\u2591\u2593\u2588\u2588 \u2588\u2584 \u2592\u2593\u2588  \u2584 \u2593\u2588\u2588\u2592  \u2590\u258C\u2588\u2588\u2592  \u2592   \u2588\u2588\u2592\u2591\u2593\u2588  \u2588\u2588\u2593\u2592\u2588\u2588   \u2588\u2588\u2591\r\n  \u2592\u2588\u2588\u2592 \u2591 \u2591 \u2588\u2588\u2588\u2588\u2593\u2592\u2591\u2592\u2588\u2588\u2592 \u2588\u2584\u2591\u2592\u2588\u2588\u2588\u2588\u2592\u2592\u2588\u2588\u2591   \u2593\u2588\u2588\u2591\u2592\u2588\u2588\u2588\u2588\u2588\u2588\u2592\u2592\u2591\u2592\u2593\u2588\u2588\u2588\u2580\u2592\u2591 \u2588\u2588\u2588\u2588\u2593\u2592\u2591\r\n  \u2592 \u2591\u2591   \u2591 \u2592\u2591\u2592\u2591\u2592\u2591 \u2592 \u2592\u2592 \u2593\u2592\u2591\u2591 \u2592\u2591 \u2591\u2591 \u2592\u2591   \u2592 \u2592 \u2592 \u2592\u2593\u2592 \u2592 \u2591 \u2591\u2592   \u2592 \u2591 \u2592\u2591\u2592\u2591\u2592\u2591 \r\n    \u2591      \u2591 \u2592 \u2592\u2591 \u2591 \u2591\u2592 \u2592\u2591 \u2591 \u2591  \u2591\u2591 \u2591\u2591   \u2591 \u2592\u2591\u2591 \u2591\u2592  \u2591 \u2591  \u2591   \u2591   \u2591 \u2592 \u2592\u2591 \r\n  \u2591      \u2591 \u2591 \u2591 \u2592  \u2591 \u2591\u2591 \u2591    \u2591      \u2591   \u2591 \u2591 \u2591  \u2591  \u2591  \u2591 \u2591   \u2591 \u2591 \u2591 \u2591 \u2592  \r\n             \u2591 \u2591  \u2591  \u2591      \u2591  \u2591         \u2591       \u2591        \u2591     \u2591 \u2591  ")
	color.Red("Made by https://github.com/V4NSH4J")
	color.Blue("\nMake sure everything is configured and press ENTER to continue!")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	tokenSlice, err := ReadLines("tokens.txt")
	if err != nil {
		color.Red("%v\n", err)
		return
	}
	type Config struct {
		Capture    bool   `json:"capture"`
		ID         bool   `json:"id"`
		Username   bool   `json:"username"`
		Discrim    bool   `json:"discrim"`
		Bio        bool   `json:"bio"`
		Loc        bool   `json:"location_lang"`
		MFA        bool   `json:"mfa"`
		Email      bool   `json:"email"`
		Verified   bool   `json:"verified"`
		Phone      bool   `json:"phone"`
		Nitro      bool   `json:"nitro"`
		Delimiiter string `json:"delimiter"`
	}
	var config Config
	ex, err := os.Executable()
	if err != nil {
		return
	}
	ex = filepath.ToSlash(ex)
	file, err := os.Open(path.Join(path.Dir(ex) + "/" + "config.json"))
	if err != nil {
		return
	}
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	errr := json.Unmarshal(bytes, &config)
	if errr != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(tokenSlice))
	color.Blue("Starting Checking - Initializing - Please wait!")
	start := time.Now()
	for i := 0; i < len(tokenSlice); i++ {
		go func(i int) {
			defer wg.Done()
			r := rand.Intn(1500)
			valid := 0
			invalid := 0

			time.Sleep(time.Duration(r) * time.Millisecond)
			a := CheckToken(tokenSlice[i])
			if a == 200 {
				color.Green("[VALID]   %v [%v]\n", tokenSlice[i], a)
				valid++
				f, err := os.OpenFile(path.Join(path.Dir(ex)+"/"+"working.txt"), os.O_RDWR|os.O_APPEND, 0660)

				if err != nil {
					log.Fatal(err)
				}

				defer f.Close()

				_, err2 := f.WriteString(tokenSlice[i] + "\n")

				if err2 != nil {
					log.Fatal(err2)
				}
			} else {
				color.Red("[INVALID] %v [%v]", tokenSlice[i], a)
				invalid++
			}
			if config.Capture && a == 200 {
				b := TokenCapture(tokenSlice[i])
				c := tokenSlice[i]
				if config.ID {
					c = c + config.Delimiiter + b.ID
				}
				if config.Username {
					c = c + config.Delimiiter + b.Username
				}
				if config.Discrim {
					c = c + config.Delimiiter + b.Discrim
				}
				if config.Bio {
					c = c + config.Delimiiter + b.Bio
				}
				if config.Loc {
					c = c + config.Delimiiter + b.Locale
				}
				if config.MFA {
					c = c + config.Delimiiter + strconv.FormatBool(b.MFA)
				}
				if config.Email {
					if b.Email != "" {
						c = c + config.Delimiiter + b.Email
					} else {
						c = c + config.Delimiiter + "No Email"
					}

				}
				if config.Verified {
					c = c + config.Delimiiter + strconv.FormatBool(b.Verified)
				}
				if config.Phone {
					if b.Phone != "" {
						c = c + config.Delimiiter + b.Phone
					} else {
						c = c + config.Delimiiter + "No Phone"
					}

				}
				if config.Nitro {
					if b.PremiumType == 0 {
						nitro := "No Nitro"
						c = c + config.Delimiiter + nitro
					}
					if b.PremiumType == 1 {
						nitro := "Classic Nitro"
						c = c + config.Delimiiter + nitro
					}
					if b.PremiumType == 2 {
						nitro := "Nitro Boost"
						c = c + config.Delimiiter + nitro
					}

				}
				f, err := os.OpenFile(path.Join(path.Dir(ex)+"/"+"capture.txt"), os.O_RDWR|os.O_APPEND, 0660)

				if err != nil {
					log.Fatal(err)
				}

				defer f.Close()

				_, err2 := f.WriteString(c + "\n")

				if err2 != nil {
					log.Fatal(err2)
				}

			}

		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	color.Blue("Checked %v tokens in %v with an Average CPM of %v", len(tokenSlice), elapsed, 60*(len(tokenSlice)/int(elapsed.Seconds())))

}

func ReadLines(filename string) ([]string, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	ex = filepath.ToSlash(ex)
	fmt.Println(ex)
	file, err := os.Open(path.Join(path.Dir(ex) + "/" + filename))
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

func CheckToken(auth string) int {
	url := "https://discord.com/api/v9/users/@me/affinities/guilds"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		color.Red("%v", err)
		return -1
	}
	req.Header.Set("authorization", auth)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(CommonHeaders(req))
	if err != nil {
		color.Red("%v", err)
		return -1
	}

	return resp.StatusCode

}

type TokenCap struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	Discrim     string `json:"discriminator"`
	Bio         string `json:"bio"`
	Locale      string `json:"locale"`
	MFA         bool   `json:"mfa"`
	Email       string `json:"email"`
	Verified    bool   `json:"verified"`
	Phone       string `json:"phone"`
	PremiumType int    `json:"premium_type,omitempty"`
}

func TokenCapture(auth string) TokenCap {
	url := "https://discord.com/api/v9/users/@me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		color.Red("%v", err)
		log.Fatal(err)
	}
	req.Header.Set("authorization", auth)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(CommonHeaders(req))
	if err != nil {
		color.Red("%v", err)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var TokenCaptured TokenCap

	json.Unmarshal(body, &TokenCaptured)
	return TokenCaptured

}

func CommonHeaders(req *http.Request) *http.Request {

	req.Header.Set("x-super-properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRmlyZWZveCIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJlbi1VUyIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQ7IHJ2OjkzLjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvOTMuMCIsImJyb3dzZXJfdmVyc2lvbiI6IjkzLjAiLCJvc192ZXJzaW9uIjoiMTAiLCJyZWZlcnJlciI6IiIsInJlZmVycmluZ19kb21haW4iOiIiLCJyZWZlcnJlcl9jdXJyZW50IjoiIiwicmVmZXJyaW5nX2RvbWFpbl9jdXJyZW50IjoiIiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTAwODA0LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("x-context-properties", "eyJsb2NhdGlvbiI6IkpvaW4gR3VpbGQiLCJsb2NhdGlvbl9ndWlsZF9pZCI6Ijg4NTkwNzE3MjMwNTgwOTUxOSIsImxvY2F0aW9uX2NoYW5uZWxfaWQiOiI4ODU5MDcxNzIzMDU4MDk1MjUiLCJsb2NhdGlvbl9jaGFubmVsX3R5cGUiOjB9")
	req.Header.Set("sec-ch-ua", "'Chromium';v='92', ' Not A;Brand';v='99', 'Google Chrome';v='92'")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-GB")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.16 Chrome/91.0.4472.164 Electron/13.4.0 Safari/537.36")
	return req
}
