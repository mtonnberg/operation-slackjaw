package main

import (
	"strings"
	"time"
)

var finalStatuses = map[string]string{
	"Succeeded": `{"response_type":"channel", "attachments": [{"text":"Deployment of {appName} to {env} succeeded", "color":"good"},{"color":"good", "title": "logs", "title_link": "http://localhost:82/executions/details?executionId={jobId}"}]}`,
	"Warning":   `{"response_type":"channel", "attachments": [{"text":"Deployment of {appName} to {env} finished with warnings", "color":"warning"},{"color":"warning", "title": "logs", "title_link": "http://localhost:82/executions/details?executionId={jobId}"}]}`,
	"Error":     `{"response_type":"channel", "attachments": [{"text":"Deployment of {appName} to {env} failed!","color":"danger"},{"color":"danger", "title": "logs", "title_link": "http://localhost:82/executions/details?executionId={jobId}"}]}`,
}

func GetOtterStatus(jobId string, otterServer string, apiKey string) (status string) {
	return GetHttpGetAsText(otterServer + "/api/jobs/status?token=" + jobId + "&key=" + apiKey)
}

func TriggerDeploy(name string, otterServer string) (jobId string) {
	return GetHttpGetAsText(otterServer + "/api/jobs/trigger?template=" + name)
}

func FollowFeedbackOfProgressUntilCompletion(appName string, jobId string, env string, callbackUrl string, otterServer string, apiKey string) {
	status := ""
	for i := 0; i < 1; i = 0 {
		newStatus := GetOtterStatus(jobId, otterServer, apiKey)

		if newStatus != status {
			status = newStatus
			PostJsonAsText(callbackUrl, `{"text":"Deploying `+appName+` to `+env+`: `+status+`","response_type":"channel"}`)
		}

		if slackDto, contains := finalStatuses[status]; contains {
			slackDto = strings.Replace(slackDto, "{appName}", appName, -1)
			slackDto = strings.Replace(slackDto, "{jobId}", jobId, -1)
			slackDto = strings.Replace(slackDto, "{env}", env, -1)
			PostJsonAsText(callbackUrl, slackDto)
			break
		}

		time.Sleep(time.Millisecond * 500)
	}
}
