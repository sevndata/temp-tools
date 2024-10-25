package github_tools

import (
	"testing"
	"time"
)

func TestDeleteDeployment(t *testing.T) {
	DeleteAllDeployments()
}

func TestTime(t *testing.T) {
	println(time.Now().UTC().Unix() + 24*3600)
}
