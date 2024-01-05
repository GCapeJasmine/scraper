package usecases

import (
	"bufio"
	"context"
	"log"
	"os"
	"sync"

	"github.com/cj/scraper/internal/domains/site/repos"
)

const (
	BatchSize = 5
)

type SiteCronJob struct {
	sitesRepo repos.ISitesRepo
}

func NewSiteCronJob(sitesRepo repos.ISitesRepo) *SiteCronJob {
	return &SiteCronJob{
		sitesRepo: sitesRepo,
	}
}

func (s *SiteCronJob) UpdateSiteState(ctx context.Context) {
	sites := getListSite()
	if len(sites) == 0 {
		return
	}
	var wg sync.WaitGroup
	totalWorker := (len(sites) + BatchSize - 1) / BatchSize
	wg.Add(totalWorker)

	for i := 0; i < len(sites); i++ {
		startIdx := i * BatchSize
		endIdx := startIdx + BatchSize
		if endIdx > len(sites) {
			endIdx = len(sites)
		}
		go s.updateSiteStateByBatch(sites[startIdx:endIdx], &wg)
	}
	wg.Wait()
}

func (s *SiteCronJob) updateSiteStateByBatch(batch []string, wg *sync.WaitGroup) {
	defer wg.Done()
	return
}

func getListSite() []string {
	filePath := "sites.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil
	}
	defer file.Close()

	var sites []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sites = append(sites, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %v", err)
		return nil
	}
	return sites
}
