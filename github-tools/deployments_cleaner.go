package github_tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

//[doc](https://docs.github.com/en/rest/deployments/deployments?apiVersion=2022-11-28)

// Replace with your GitHub username
var username = ""

// Replace with your GitHub token,[personal access token](https://github.com/settings/tokens)
var token = ""

// Replace with your GitHub repositoryName
var repositoryName = ""

var baseURL = "https://api.github.com/repos/" + username + "/" + repositoryName
var deploymentsURL = baseURL + "/deployments"
var statusesURL = baseURL + "/deployments/%d/statuses"

// Function to get deployments
func getDeployments() (*http.Response, error) {
	req, err := http.NewRequest("GET", deploymentsURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, token)

	client := &http.Client{}
	return client.Do(req)
}

// Function to mark a deployment as inactive
func inactiveDeployment(deploymentId int) (*http.Response, error) {
	url := fmt.Sprintf(statusesURL, deploymentId)
	data := map[string]string{"state": "inactive"}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.ant-man-preview+json")

	client := &http.Client{}
	return client.Do(req)
}

// Function to delete a deployment
func deleteDeployment(deploymentId int) (*http.Response, error) {
	inactivateResponse, err := inactiveDeployment(deploymentId)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(inactivateResponse.Body)

	if inactivateResponse.StatusCode == 201 {
		fmt.Printf("Deployment %d marked as inactive.\n", deploymentId)

		url := fmt.Sprintf("%s/%d", deploymentsURL, deploymentId)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth(username, token)

		client := &http.Client{}
		return client.Do(req)
	} else {
		fmt.Printf("Failed to update deployment %d. Status Code: %d\n", deploymentId, inactivateResponse.StatusCode)
		return nil, fmt.Errorf("failed to update deployment")
	}
}

// Function to delete all deployments found in this query
func deleteDeploymentsThisQuery() bool {
	response, err := getDeployments()
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if response.StatusCode != 200 {
		fmt.Printf("Failed to retrieve deployments. Status Code: %d\n", response.StatusCode)
		return false
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var deployments []map[string]interface{}
	if err := json.Unmarshal(body, &deployments); err != nil {
		log.Fatal(err)
	}

	if len(deployments) == 0 {
		return false
	}

	for _, deployment := range deployments {
		id, ok := deployment["id"].(float64)
		if !ok {
			continue
		}
		_, err := deleteDeployment(int(id))
		if err != nil {
			log.Fatal(err)
		}
	}

	return true
}

// DeleteAllDeployments Function to delete all deployments
func DeleteAllDeployments() {
	for {
		if !deleteDeploymentsThisQuery() {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
