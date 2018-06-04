package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var feedURL = "http://feeds.nightvalepresents.com/welcometonightvalepodcast"

type feedData struct {
	Channel feedChannel `xml:"channel"`
}

type feedChannel struct {
	Items []feedItem `xml:"item"`
}

type feedItem struct {
	Title string    `xml:"title"`
	URL   enclosure `xml:"enclosure"`
}

type enclosure struct {
	URL string `xml:"url,attr"`
}

func main() {
	resp, err := http.Get(feedURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data feedData
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range data.Channel.Items {
		if err := download(item); err != nil && !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
}

func download(item feedItem) error {
	file := filename(item.Title)

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		log.Printf("Skipping %s, already exists\n", file)
		return err
	}

	fmt.Printf("Downloading to %s\n", file)

	// Create output file
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get HTTP resp
	resp, err := http.Get(item.URL.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func filename(title string) string {
	// Remove spaces and append mp3 suffix
	title = strings.Replace(title, " ", "_", -1) + ".mp3"

	// Adjust episode number to maintain ascii sorting
	num, text := 0, ""
	_, err := fmt.Sscanf(title, "%d_-_%s", &num, &text)
	if err != nil {
		//log.Println(err)
		return title
	}

	prefix := ""
	switch {
	case num < 10:
		prefix = "00"
	case num < 100:
		prefix = "0"
	}

	return strings.Replace(
		title,
		fmt.Sprintf("%d_-", num),
		fmt.Sprintf("%s%d_-", prefix, num),
		1)
}
