package repo

import (
	"context"
	"scheduler/internal/models"

	"gorm.io/gorm"
)

type SchedulerRepo struct {
	jobsCache []models.Job
	db        *gorm.DB
}

type ISchedulerRepo interface {
	AddJob(ctx context.Context, job models.Job) (int, error)
	GetJobs(ctx context.Context) (int64, []models.Job, error)
	DeleteJob(ctx context.Context, job models.Job) (int, error)
	UpdateJob(ctx context.Context, job models.Job) (int, int64, error)
	GetJob(ctx context.Context, job models.Job) (int64, models.Job, error)
}

func NewSchedulerRepo(db *gorm.DB) (ISchedulerRepo, error) {
	sr := SchedulerRepo{
		db: db,
	}

	sr.jobsCache = []models.Job{}

	// rows, jobsCache, err := sr.GetJobs(context.Background())
	// if err != nil {
	// 	return nil, err
	// }

	// if rows > 0 {
	// 	sr.jobsCache = jobsCache
	// }
	return sr, nil
}

func (schRepo SchedulerRepo) AddJob(ctx context.Context, job models.Job) (jobID int, err error) {
	result := schRepo.db.Create(&job)
	return job.ID, result.Error
}

func (schRepo SchedulerRepo) GetJobs(ctx context.Context) (int64, []models.Job, error) {
	jobs := []models.Job{}
	result := schRepo.db.Find(&jobs)
	return result.RowsAffected, jobs, result.Error
}

func (schRepo SchedulerRepo) DeleteJob(ctx context.Context, job models.Job) (int, error) {
	result := schRepo.db.Delete(job.ID)
	return int(result.RowsAffected), result.Error
}

func (schRepo SchedulerRepo) UpdateJob(ctx context.Context, job models.Job) (int, int64, error) {
	result := schRepo.db.Save(&job)
	return job.ID, result.RowsAffected, result.Error
}

func (schRepo SchedulerRepo) GetJob(ctx context.Context, job models.Job) (int64, models.Job, error) {
	result := schRepo.db.Find(&job)
	return result.RowsAffected, job, result.Error
}
