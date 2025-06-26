package main

import (
	"fmt"
	"log"
	"regexp"
	"scraper/utils"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type UniversityCoordsFetcher struct {
	Universities []University
}

type Coord struct {
	place string
	lat   string
	lng   string
}

func NewUniversityCoordsFetcher(universities []University) *UniversityCoordsFetcher {
	return &UniversityCoordsFetcher{
		Universities: universities,
	}
}

func webScrapeCoord(universityName string, e *colly.HTMLElement) Coord {
	coords := []Coord{}

	e.ForEach("tr:nth-child(n+2):nth-child(-n+7)", func(i int, h *colly.HTMLElement) {
		place := h.ChildText("td:nth-child(2) a:nth-child(1)")
		lat := h.ChildText("td:nth-child(5)")
		lng := h.ChildText("td:nth-child(6)")
		coordObj := Coord{
			place: place,
			lat:   lat,
			lng:   lng,
		}
		coords = append(coords, coordObj)

		place = h.ChildText("td[nowrap] > b > a:nth-child(1)")
		latlng := h.ChildText("td[nowrap] > a")
		if latlng != "" {
			llArr := strings.Split(latlng, " / ")
			lat = llArr[0]
			lng = llArr[1]
		}
		coordObj = Coord{
			place: place,
			lat:   lat,
			lng:   lng,
		}
		coords = append(coords, coordObj)
	})

	universityNames := []string{}
	for _, coord := range coords {
		universityNames = append(universityNames, coord.place)
	}

	index, _ := utils.FindMostSimilarText(universityName, universityNames)
	return coords[index]
}

func ConvertDMS(dms string) (string, error) {
	// Normalize different prime symbols (′ and ″) to ASCII-friendly characters (') and (")
	dms = strings.ReplaceAll(dms, "′", "'")
	dms = strings.ReplaceAll(dms, "″", "\"")

	// Match patterns for DMS coordinates
	dmsPattern := regexp.MustCompile(`([NSEW])?\s*(\d+)°\s*(\d+)'?\s*(\d+)"?`)
	matches := dmsPattern.FindStringSubmatch(dms)

	if matches == nil {
		// If not in DMS, try to parse directly as a decimal degree
		if dms != "" {
			f, err := strconv.ParseFloat(strings.TrimSpace(dms), 64)
			return fmt.Sprintf("%.4f", f), err
		} else {
			return "", nil
		}
	}

	// Extract degree, minute, second, and direction
	dir := matches[1]
	degree, _ := strconv.Atoi(matches[2])
	minute, _ := strconv.Atoi(matches[3])
	second, _ := strconv.Atoi(matches[4])

	// Convert DMS to decimal degrees
	decimal := float64(degree) + float64(minute)/60 + float64(second)/3600

	// Apply direction
	switch dir {
	case "S", "W":
		decimal = -decimal
	}

	return fmt.Sprintf("%.4f", decimal), nil
}

func scrapeUniversityCoord(universities []University) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.geonames.org", "geonames.org"),
		colly.Async(true),
	)

	scraperQueue, err := queue.New(90, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		log.Fatal("Cannot instantiate web scraper queue:", err)
	}

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*.geonames.org",
		Parallelism: 90,
	})

	var mu sync.Mutex

	c.OnHTML("table.restable ", func(e *colly.HTMLElement) {
		universityName := strings.ReplaceAll(e.Request.URL.Query().Get("q"), "+", " ")
		index, err := strconv.Atoi(e.Request.URL.Query().Get("uniIndex"))
		if err != nil {
			log.Fatalln("cannot parse uniIndex query param:", err)
		}
		coord := webScrapeCoord(universityName, e)
		dcLat, err := ConvertDMS(coord.lat)
		if err != nil {
			log.Fatalln("Unable to convert latitude to DD:", err)
		}
		dcLng, err := ConvertDMS(coord.lng)
		if err != nil {
			log.Fatalln("Unable to convert longitude to DD:", err)
		}

		mu.Lock()
		universities[index].Latitude = dcLat
		universities[index].Longitude = dcLng
		mu.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting Url:", r.URL)
	})

	for i, university := range universities {
		uniNameFormat := strings.ReplaceAll(university.Name, " ", "+")
		url := fmt.Sprintf("https://www.geonames.org/search.html?q=%s&country=&uniIndex=%d", uniNameFormat, i)
		if err = scraperQueue.AddURL(url); err != nil {
			fmt.Printf("Unable to add url to scraper queue with university %d: %s", i, err)
		}
	}
	scraperQueue.Run(c)
	c.Wait()
}

func (u *UniversityCoordsFetcher) GetUniversityCoords() {
	start := time.Now()
	fmt.Println("Scraping university coords...")
	scrapeUniversityCoord(u.Universities)
	end := time.Since(start)

	noCoordCount := 0
	fmt.Println("100 university samples and their coordinates:")
	for i, university := range u.Universities {
		if len(university.Latitude) < 1 {
			noCoordCount++
		}
		if i == 100 {
			continue
		}
		fmt.Print(university.Name, ": ")
		fmt.Printf("%s | %s\n", university.Latitude, university.Longitude)
	}

	fmt.Println("no of missing data:", noCoordCount)
	fmt.Println("time it took to scrape all university coordinates:", end.Seconds(), "secs")
}
