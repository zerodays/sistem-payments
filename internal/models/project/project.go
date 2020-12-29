package project

import (
	"encoding/json"
	"github.com/sony/gobreaker"
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

var projectsCache = make([]*Project, 0)
var projectsCachePopulated = false

var cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
	OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
		logger.Log.Info().Stringer("from_state", from).Stringer("to_state", to).Msg("circuit breaker changed state")
	},
})

// All returns all projects from projects microservice.
func All() (Projects, error) {
	pr, err := cb.Execute(func() (interface{}, error) {
		url := config.Microservices.ProjectsUrl() + "/projects"

		cli := &http.Client{}
		cli.Timeout = 2 * time.Second

		resp, err := cli.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		projects := make([]*Project, 0)
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&projects)
		if err != nil {
			return nil, err
		}

		return projects, err
	})

	if err != nil {
		if projectsCachePopulated {
			logger.Log.Warn().Err(err).Msg("using projects cache")
			return projectsCache, nil
		} else {
			return nil, err
		}
	} else {
		projects := pr.([]*Project)
		projectsCache = projects
		projectsCachePopulated = true

		return projects, nil
	}
}
