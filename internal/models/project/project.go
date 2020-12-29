package project

import (
	"encoding/json"
	"github.com/zerodays/sistem-payments/internal/config"
	"net/http"
)

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DateCreated string `json:"date_created"`
}

type Projects []*Project

// All returns all projects from projects microservice.
func All() (Projects, error) {
	url := config.Microservices.ProjectsUrl() + "/projects"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	projects := make(Projects, 0)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&projects)

	return projects, err
}
