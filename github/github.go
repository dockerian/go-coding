package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API interface
type API interface {
	GetLatestReleaseTag(string) (string, error)
	GetReleaseTags(string) ([]string, error)
}

// Github struct
type Github struct{}

// Release struct
type Release struct {
	ID      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

// GetLatestReleaseTag returns the lastest release tag
func (gh *Github) GetLatestReleaseTag(repo string) (string, error) {
	tags, err := gh.GetReleaseTags(repo)
	if err != nil {
		return "", err
	}
	return tags[0], nil
}

// GetReleaseTags returns all releases tag name per a repository
func (gh *Github) GetReleaseTags(repo string) ([]string, error) {
	tags := []string{}
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	res, err := http.Get(apiURL)
	if err != nil {
		return tags, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return tags, err
	}

	releases := []Release{}
	if err := json.Unmarshal(body, &releases); err != nil {
		return tags, err
	}

	tags = make([]string, len(releases))
	for i, release := range releases {
		tags[i] = release.TagName
	}

	return tags, nil
}
