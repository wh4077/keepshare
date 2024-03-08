// // Copyright 2024 The KeepShare Authors. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file.

package rapidgator

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/KeepShareOrg/keepshare/pkg/share"
)

// GetStatuses return the statuses of each host shared link.
func (p *RapidGator) GetStatuses(ctx context.Context, userID string, hostSharedLinks []string) (statuses map[string]share.State, err error) {
	// todo: get account token by keepShareUserID;
	token, err := getToken()
	if err != nil {
		return nil, fmt.Errorf("RapidGator GetStatuses getToken err: %w", err)
	}

	// todo: single token or all token?
	_, remoteUploadJobsByHostLink := getRemoteUploadInfo(token)

	statuses = make(map[string]share.State, len(hostSharedLinks))
	for _, link := range hostSharedLinks {
		status := share.StatusNotFound
		if job, ok := remoteUploadJobsByHostLink[strings.ToLower(link)]; ok {
			status = jobStateLabel2State(job.StateLabel)
		}
		statuses[link] = status
	}

	return statuses, nil
}

func getRemoteUploadInfo(token string) (
	remoteUploadJobsByOriginLink map[string]RemoteUploadJob,
	remoteUploadJobsByHostLink map[string]RemoteUploadJob) {
	var remoteUploadInfoResponse struct {
		Status   int `json:"status"`
		Response struct {
			Jobs []RemoteUploadJob `json:"jobs"`
		} `json:"response"`
	}

	remoteUploadInfoURL := baseURL + fmt.Sprintf(remoteUploadInfoAPIFormat, token)
	_, respErr := restyClient.R().
		SetResult(&remoteUploadInfoResponse).
		Post(remoteUploadInfoURL)

	// maybe unmarshal JSON failed.
	if respErr != nil {
		// fmt.Println("RapidGator getRemoteUploadInfo respErr:", respErr)
	}
	// so check remoteUploadInfoResponse.Status;
	if remoteUploadInfoResponse.Status != http.StatusOK {
		fmt.Println(fmt.Errorf("RapidGator getRemoteUploadInfo.Status %d != http.StatusOK", remoteUploadInfoResponse.Status))

		return
	}

	remoteUploadJobsByOriginLink = map[string]RemoteUploadJob{}
	remoteUploadJobsByHostLink = map[string]RemoteUploadJob{}
	jobCount := len(remoteUploadInfoResponse.Response.Jobs)
	if jobCount > 0 {
		for _, job := range remoteUploadInfoResponse.Response.Jobs {
			// todo: *job or job?
			remoteUploadJobsByOriginLink[strings.ToLower(job.RemoteUploadURL)] = job
			remoteUploadJobsByHostLink[strings.ToLower(job.File.URL)] = job
		}
	}

	return remoteUploadJobsByOriginLink, remoteUploadJobsByHostLink
}

func jobStateLabel2State(jobStateLabel string) (status share.State) {
	status = share.StatusUnknown

	switch jobStateLabel {
	case "Done":
		status = share.StatusOK

	case "Downloading":
		status = share.StatusCreated

	case "Fail":
		status = share.StatusError

	case "Canceled":
		status = share.StatusDeleted

	case "Waiting":
		status = share.StatusPending

	default:
		status = share.StatusUnknown
	}

	return status
}
