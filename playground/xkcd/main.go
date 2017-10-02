package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// XkcdImg Image info from xkcd json
type XkcdImg struct {
	Day        string
	Month      string
	Year       string
	Num        int
	Link       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
}

func (img *XkcdImg) Print(showTranscript bool) {
	fmt.Printf("Num: %d (%s-%s-%s)> %s \t%s\n", img.Num,
		img.Year, img.Month, img.Day, img.SafeTitle, img.Img)

	if showTranscript {
		fmt.Println(img.Transcript)
	}
}

const apiURL = "https://xkcd.com"
const imgsToLoad = 1890 // 1890
var images = make([]*XkcdImg, imgsToLoad*2)

var transcript bool
var all bool
var limit int

func init() {
	flag.BoolVar(&transcript, "transcript", false, "Show transcripts")
	flag.BoolVar(&all, "all", false, "Fetch all")
	flag.IntVar(&limit, "max", imgsToLoad, "Max number of items to fetch")
}

func main() {
	flag.Parse()

	os.Mkdir("/tmp/xkcd", 0777)

	if all {
		fetchAll(limit)
	}

	if arg, hasArg := getArg(); hasArg {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fetchAll(limit)
			search(arg, transcript)
			return
		}

		img, err := fetchByID(num)
		if err != nil {
			log.Fatal(err)
		}
		img.Print(transcript)
	} else {
		img, _ := fetchLatest()
		img.Print(transcript)
	}
}

func getArg() (string, bool) {
	for _, arg := range os.Args[1:] {
		if arg[:1] != "-" {
			return arg, true
		}
	}
	return "", false
}

func search(q string, showTranscript bool) {
	// half := len(images) / 2

	// ch := make(chan bool)

	// scanSlice := func(imgs []*XkcdImg) {
	for _, img := range images {
		// fmt.Printf("Searching %d\n", img.Num)
		if img != nil && img.Num > 0 && strings.Contains(img.SafeTitle+img.Transcript, q) {
			img.Print(showTranscript)
		}
	}
	// 	ch <- true
	// }

	// go scanSlice(images[:half])
	// go scanSlice(images[half:])

	// <-ch
	// <-ch
}

func fetchAll(imgsToLoad int) {
	var ch = make(chan int)

	for i := 1; i <= imgsToLoad; i++ {
		go (func(id int) {
			img, _ := fetchByID(id)
			// fmt.Printf("%d| %s\n", id, err)

			images = append(images, img)
			// fmt.Printf("Fetched %v\n", img)

			ch <- id
		})(i)

		<-ch
	}

	fmt.Printf("Total images fetched: %d", len(images))
}

func fetchLatest() (*XkcdImg, error) {
	return fetchByURL(fmt.Sprintf("%s/info.0.json", apiURL))
}

func fetchByID(id int) (*XkcdImg, error) {
	cachedFileName := fmt.Sprintf("/tmp/xkcd/xkcd_%d.json", id)

	if img, err := getCached(cachedFileName); err == true {
		return img, nil
	}

	return fetchByURL(fmt.Sprintf("%s/%d/info.0.json", apiURL, id))
}

func fetchByURL(url string) (*XkcdImg, error) {
	// fmt.Printf("-> %s \n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Cannot fetch status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var imgInfo XkcdImg

	if err := json.Unmarshal(body, &imgInfo); err != nil {
		// fmt.Printf("%s\n", err)
		return nil, err
	}

	defer writeCache(&imgInfo)

	return &imgInfo, nil
}

func getCached(fname string) (*XkcdImg, bool) {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		// fmt.Printf("Cache: %s\n", err)
		return nil, false
	}

	var img XkcdImg
	if err := json.Unmarshal(content, &img); err != nil {
		fmt.Printf("Cache json: %s\n", err)
		return nil, false
	}

	return &img, true
}

func writeCache(img *XkcdImg) {
	cachedFileName := fmt.Sprintf("/tmp/xkcd/xkcd_%d.json", img.Num)

	if body, err := json.Marshal(img); err != nil {
		if err := ioutil.WriteFile(cachedFileName, body, 0644); err != nil {
			fmt.Printf("Cannot write file %s: %s", cachedFileName, err)
		}
	}
}
