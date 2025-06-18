package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"scheduler/internal/models"
	"scheduler/internal/service"
	"scheduler/internal/validator"

	"github.com/gorilla/mux"
)

type SchedulerHandler struct {
	service service.ISchedulerService
}

// NewPlayerHandler initializes a new PlayerHandler with the provided BettorService.
func NewSchedulerHandler(service service.ISchedulerService) SchedulerHandler {
	return SchedulerHandler{service: service}
}

func (scheduler *SchedulerHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/get_jobs", scheduler.GetJobs).Methods("GET")
	r.HandleFunc("/create_job", scheduler.CreateJob).Methods("POST")
	// r.HandleFunc("/get_jobs/{jobID}", scheduler.GetJobByID).Methods("GET")
	r.HandleFunc("/get_jobs/{jobID}", scheduler.UpdateJob).Methods("PATCH")
	r.HandleFunc("/delete_job/{jobID}", scheduler.DeleteJob).Methods("DELETE")
}

func (scheduler *SchedulerHandler) CreateJob(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Create the job in the database and tag them to the scheduler.
	job := models.Job{}
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		SendErrorResponse(ctx, w, http.StatusInternalServerError, err)
	}

	// Validating the request object and returning the errors accordingly as a response.
	if err := validator.RequestValidator(job); err != nil {
		SendErrorResponse(ctx, w, http.StatusUnprocessableEntity, err)
	}

	cmd := func() {}

	scheduler.service.AddJobToCronScheduler(ctx, job, cmd)
}

func (scheduler *SchedulerHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	// Fetch the jobs from the database and tag them to the scheduler.
}

func (scheduler SchedulerHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	// Delete the job from the database and untag from the schedulers scheduled

	ctx := r.Context()

	params := mux.Vars(r)
	job := models.Job{
		Name: params["jobID"],
	}

	scheduler.service.DeleteJob(ctx, job)
}

func (scheduler SchedulerHandler) UpdateJob(w http.ResponseWriter, r *http.Request) {
	// Update the job in the database and tag to the schedulers scheduled
}

func SendErrorResponse(ctx context.Context, w http.ResponseWriter, errCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	if err := json.NewEncoder(w).Encode(models.Response{
		Status:           "Failure",
		ErrorCode:        http.StatusText(errCode),
		ErrorDescription: err.Error(),
	}); err != nil {
		// TODO: Write the log line....
	}
}
