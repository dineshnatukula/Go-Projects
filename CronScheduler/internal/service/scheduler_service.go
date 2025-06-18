package service

import (
	"context"
	"fmt"
	"net/http"
	"scheduler/internal/models"
	"scheduler/internal/repo"
	"sync"

	"github.com/robfig/cron/v3"
)

type SchedulerService struct {
	scheduledjobs sync.Map
	cron          *cron.Cron
	SchedulerRepo repo.ISchedulerRepo
}

type ISchedulerService interface {
	AddJob(ctx context.Context, job models.Job)
	GetJobs(ctx context.Context, jobs []models.Job)
	DeleteJob(ctx context.Context, job models.Job)
	UpdateJob(ctx context.Context, job models.Job)
	AddJobToCronScheduler(ctx context.Context, job models.Job, cmd func()) error
}

func NewSchedulerService(schRepo repo.ISchedulerRepo) ISchedulerService {
	schSVC := &SchedulerService{
		cron:          cron.New(cron.WithSeconds()),
		scheduledjobs: sync.Map{},
		SchedulerRepo: schRepo,
	}
	schSVC.cron.Start()
	return schSVC
}

func (schService *SchedulerService) AddJob(ctx context.Context, job models.Job) {

}

func (schService *SchedulerService) GetJobs(ctx context.Context, jobs []models.Job) {

}

func (schService *SchedulerService) DeleteJob(ctx context.Context, job models.Job) {
	schService.RemoveJob(ctx, job)
}

func (schService *SchedulerService) UpdateJob(ctx context.Context, job models.Job) {

}

func (s *SchedulerService) AddJobToCronScheduler(ctx context.Context, job models.Job, makeRequest func()) error {
	mr := func() {
		fmt.Printf("[CRON RUN] Executing job: %s (%s %s)\n", job.Name, job.HttpMethod, job.URL+job.EndPoint)
		req, err := http.NewRequest(job.HttpMethod, job.URL+job.EndPoint, nil)
		if err != nil {
			fmt.Printf("HTTP request failed: %v\n", err)
			return
		}
		// Optionally, you can add headers or other stuff to req here
		// e.g. req.Header.Set("Authorization", "Bearer ...")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("HTTP request failed: %v\n", err)
			return
		}
		// You can handle response here, e.g., read body or check status
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			fmt.Printf("HTTP request returned status %d\n", resp.StatusCode)
		}
		defer resp.Body.Close()
	}

	cronEntryID, err := s.cron.AddFunc(job.CronSpec, mr)
	fmt.Println("Cron Scheduled...", cronEntryID, err)
	if err != nil {
		return err
	}

	s.scheduledjobs.Store(job.Name, cronEntryID)
	return nil
}

func (s *SchedulerService) RemoveJob(ctx context.Context, job models.Job) {
	value, ok := s.scheduledjobs.Load(job.Name)
	fmt.Println("value", value, "name", job.Name)

	cronEntryID := value.(cron.EntryID)
	if !ok {
		return
	}
	s.cron.Remove(cronEntryID)
	fmt.Println("Removed the job from the cron scheduler...")

	s.scheduledjobs.Delete(job.Name)
	fmt.Println("Removed the job from the map...")
}
