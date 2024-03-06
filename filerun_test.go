package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestFilerun(t *testing.T) {
	//baseUrl := "https://imgo.erps.bio:88/filerun/"
	baseUrl := "https://pan.tiit.ga:88/imgo/"

	reqUrl := "?module=fileman&section=do&page=delete"

	reqUrl = baseUrl + reqUrl

	params := []string{
		"paths[]=/ROOT/HOME/imgGo_upload/2023/06/28_15-23-56.837.png.webp",
		"paths[]=/ROOT/HOME/imgGo_upload/2023/06/28_16-29-24.814.png.webp",
		"paths[]=/ROOT/HOME/imgGo_upload/2023/06/28_18-27-00.515.png.webp",
		"paths[]=/ROOT/HOME/imgGo_upload/2023/06/28_15-22-18.233.png.webp",
		"paths[]=/ROOT/HOME/imgGo_upload/2023/06/28_13-54-41.926.png.webp",
		"csrf=508488d7dcc43864916c674187056c25931a05145995f414",
	}

	data := strings.Join(params, "&")

	// Parse data
	parsedData, err := url.ParseQuery(data)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}
	formData := io.NopCloser(bytes.NewReader([]byte(parsedData.Encode())))

	// Create a new request
	req, err := http.NewRequest(http.MethodPost, reqUrl, formData)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://pan.tiit.ga:88")
	req.Header.Set("Referer", "https://pan.tiit.ga:88")
	req.Header.Set("Cookie", "language=chinese; FileRun[token]=f2453c83e222a5e64edd5b1a33e2b8ad1c0b2f4a0031a5b9")
	req.Header.Set("Authority", "pan.tiit.ga:88")

	// Create an HTTP client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response body:", string(body))
}
