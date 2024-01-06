package usecases

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cj/scraper/internal/domain/models"
	"github.com/cj/scraper/internal/domains/site/repos"
)

const (
	BatchSize    = 5
	NotAvailable = "not_available"
	Available    = "available"
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

	for i := 0; i < totalWorker; i++ {
		startIdx := i * BatchSize
		endIdx := startIdx + BatchSize
		if endIdx > len(sites) {
			endIdx = len(sites)
		}
		go s.updateSiteStateByBatch(ctx, sites[startIdx:endIdx], &wg)
	}
	wg.Wait()
}

func (s *SiteCronJob) updateSiteStateByBatch(ctx context.Context, batch []string, wg *sync.WaitGroup) {
	defer wg.Done()
	sites := make([]*models.Site, 0)
	for _, site := range batch {
		state, latency, err := s.getState(site)
		if err != nil {
			continue
		}
		sites = append(sites, &models.Site{
			Name:       site,
			AccessTime: latency.Milliseconds(),
			State:      state,
		})
	}
	for _, site := range sites {
		s.sitesRepo.Upsert(ctx, site)
	}
}

func (s *SiteCronJob) getState(site string) (string, time.Duration, error) {
	if !strings.HasPrefix(site, "http://") && !strings.HasPrefix(site, "https://") {
		site = "https://" + site
	}
	start := time.Now()
	resp, err := http.Get(site)
	if err != nil {
		log.Printf("getLatency error %v", err)
		return "", 0, err
	}
	defer resp.Body.Close()
	elapsed := time.Since(start)
	if resp.StatusCode != 200 {
		return NotAvailable, 0, nil
	}

	return Available, elapsed, nil
}

func getListSite() []string {
	filePath := "internal/domains/site/usecases/sites.txt"

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
