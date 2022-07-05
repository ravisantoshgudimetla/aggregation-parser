package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type job struct {
	URL     string
	Status  string
	Variant string
	Name    string
}

type jobs map[string]job

func getJobURLs(jobURLs []string, jobName string) string {
	trimmedJobName := strings.Split(jobName, "master-")[1]
	for _, jobURL := range jobURLs {
		if strings.Contains(jobURL, trimmedJobName) {
			return jobURL
		}
	}
	return ""
}

// filterJobsBasedOnVariants filters the jobs in the required format as presented in spreadsheet
func filterJobsBasedOnVariants(totalJobs jobs) jobs {
	variants := make(map[string][]string)
	platforms := []string{"aws", "azure", "gcp", "metal-ipi", "vsphere"}
	for _, platform := range platforms {
		variants[platform] = []string{
			platform + "-upi-serial", platform + "-upi", // upi specific jobs
			platform + "-serial",                              // serial
			platform + "-csi",                                 // csi
			platform + "-single-serial", platform + "-single", // single node
			platform + "-techpreview-serial", platform + "-techpreview", // tech preview jobs
			platform + "-ovn-serial", platform + "-ovn", // ovn specific jobs
			platform + "-rt",                                                               //rt specific jobs
			"console-" + platform,                                                          // console
			platform + "-upgrade-ovn",                                                      // ovn upgrade jobs
			platform + "-upgrade-single", platform + "-upgrade-cnv", platform + "-upgrade", // upgrade jobs
			"-upgrade-from-stable-" + platform,
			platform, // e2e-parallel
		}
	}
	jobsCollected := make(jobs)
	for _, platform := range platforms {
		for _, probableJob := range variants[platform] {
			for jobName, jobStatus := range totalJobs {
				if strings.Contains(jobName, probableJob) {
					if _, ok := jobsCollected[jobName]; ok {
						continue
					} else {
						jobStatus.Name = jobName
						jobStatus.Variant = probableJob
						jobsCollected[jobName] = jobStatus
					}
				}
			}
		}
	}
	return jobsCollected
}

func Parse(inputURL string) []job {
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
		js := job{}
		js.Status = "failed"
		js.URL = getJobURLs(jobUrls, jobName)
		jobs[h.Text] = js
	})
	// get all successful jobs
	c.OnHTML("span.text-success", func(h *colly.HTMLElement) {
		jobName := h.Text
		js := job{}
		js.Status = "Success"
		js.URL = getJobURLs(jobUrls, jobName)
		jobs[h.Text] = js
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})
	c.Visit(inputURL)
	finalJobs := filterJobsBasedOnVariants(jobs)
	var jobCollection []job
	for _, jobStatus := range finalJobs {
		jobCollection = append(jobCollection, jobStatus)
	}
	return jobCollection
	//fmt.Printf("variant: %v Jobs: %v Status: %v\n", jobStatus.variant, jobStatus.url, jobStatus.status)
}
