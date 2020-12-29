package project

import (
	"encoding/json"
	"github.com/zerodays/sistem-payments/internal/config"
	"github.com/zerodays/sistem-payments/internal/logger"
	"net/http"
	"time"
)

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DateCreated string `json:"date_created"`
}

type Projects []*Project

var projectsCache []*Project
var projectsCachePopulated = false

// All returns all projects from projects microservice.
func All() (Projects, error) {
	url := config.Microservices.ProjectsUrl() + "/projects"

	cli := &http.Client{}
	cli.Timeout = 2 * time.Second

	resp, err := cli.Get(url)
	if err != nil {
		if projectsCachePopulated {
			logger.Log.Warn().Err(err).Msg("using fallback for projects")
			return projectsCache, nil
		} else {
			return nil, err
		}
	}
	defer resp.Body.Close()

	projects := make(Projects, 0)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&projects)
	if err != nil {
		if projectsCachePopulated {
			logger.Log.Warn().Err(err).Msg("using fallback for projects")
			return projectsCache, nil
		} else {
			return nil, err
		}
	}

	projectsCache = projects
	projectsCachePopulated = true

	return projects, err
}
