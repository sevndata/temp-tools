package github_tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const githubToken = "Personal access tokens (classic)"
const orgName = "Organization name"

type Repository struct {
	Name string `json:"name"`
}

type Commit struct {
	Sha    string `json:"sha"`
	Commit struct {
		Author struct {
			Name string `json:"name"`
			Date string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
}

// get repositories by orgName
func getRepositories() ([]Repository, error) {
	reposURL := fmt.Sprintf("https://api.github.com/orgs/%s/repos", orgName)
	req, err := http.NewRequest("GET", reposURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+githubToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var repos []Repository
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

// get commits by repoName
func getCommits(repoName string) ([]Commit, error) {
	commitsURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", orgName, repoName)
	req, err := http.NewRequest("GET", commitsURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+githubToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var commits []Commit
	if err := json.Unmarshal(body, &commits); err != nil {
		return nil, err
	}
	return commits, nil
}

// GetAllCommit get all commits by orgName
func GetAllCommit() {
	repos, err := getRepositories()
	if err != nil {
		log.Fatalf("Error getting repositories: %v", err)
	}
	//outputFile, err := os.Create("git-commits.txt")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	//defer outputFile.Close()
	for _, repo := range repos {
		fmt.Printf("\nFetching commits for repository: \033[31m%s\n", repo.Name)
		commits, err := getCommits(repo.Name)
		if err != nil {
			log.Printf("Error getting commits for repo %s: %v", repo.Name, err)
			continue
		}
		for _, commit := range commits {
			commitDate, _ := time.Parse(time.RFC3339, commit.Commit.Author.Date)
			//output := fmt.Sprintf("Repository: %s\nCommit SHA: %s\nAuthor: %s\nDate: %s\nMessage: %s\n\n",
			//	repo.Name, commit.Sha, commit.Commit.Author.Name, commitDate.Format("2006-01-02 15:04:05"), commit.Commit.Message)
			//outputFile.WriteString(output)
			fmt.Printf("Date: %s,Author: %s,Message: %s\n",
				commitDate.Format("2006-01-02 15:04:05"), commit.Commit.Author.Name, commit.Commit.Message)
		}
	}
}
