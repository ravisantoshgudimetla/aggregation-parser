package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type jobStatus struct {
	url    string
	status string
}

type jobs map[string]jobStatus

func getJobURLs(jobURLs []string, jobName string) string {
	trimmedJobName := strings.Split(jobName, "master-")[1]
	for _, jobURL := range jobURLs {
		if strings.Contains(jobURL, trimmedJobName) {
			return jobURL
		}
	}
	return ""
}

func Parse(inputURL string) {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	jobs := make(jobs)
	// collect all URLs
	jobUrls := make([]string, 0)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "prow") {
			jobUrls = append(jobUrls, link)
		}
	})
	// get all the failed jobs
	c.OnHTML("span.text-danger", func(h *colly.HTMLElement) {
		jobName := h.Text
		js := jobStatus{}
		js.status = "failed"
		js.url = getJobURLs(jobUrls, jobName)
		jobs[h.Text] = js
	})
	// get all successful jobs
	c.OnHTML("span.text-success", func(h *colly.HTMLElement) {
		jobName := h.Text
		js := jobStatus{}
		js.status = "Success"
		js.url = getJobURLs(jobUrls, jobName)
		jobs[h.Text] = js
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})
	c.Visit(inputURL)
	for jobName, jobStatus := range jobs {
		fmt.Println(jobName, jobStatus)
	}
}
