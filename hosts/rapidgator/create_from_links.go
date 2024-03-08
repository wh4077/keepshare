// Copyright 2023 The KeepShare Authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// https://rapidgator.net/article/api/remote

package rapidgator

import (
	"context"
	"net/http"
	"strings"
	"time"

	"errors"
	"fmt"

	"github.com/KeepShareOrg/keepshare/pkg/log"
	"github.com/KeepShareOrg/keepshare/pkg/share"
)

type RemoteUploadFile struct {
	FileID string `json:"file_id"`
	Hash   string `json:"hash"`
	Size   int64  `json:"Size"`
	URL    string `json:"url"`
}

type RemoteUploadJob struct {
	JobID           int              `json:"job_id"`
	RemoteUploadURL string           `json:"url"`
	Name            string           `json:"name"`
	StateLabel      string           `json:"state_label"`
	File            RemoteUploadFile `json:"file"`
	Error           string           `json:"error"`
}

// https://rapidgator.net/article/api/remote#create
const remoteUploadCreateAPIFormat = "/api/v2/remote/create?url=%s&token=%s"

// https://rapidgator.net/article/api/remote#info
const remoteUploadJobsInfoAPIFormat = "/api/v2/remote/info?job_id=%s&token=%s"
const remoteUploadInfoAPIFormat = "/api/v2/remote/info?token=%s"

// https://rapidgator.net/article/api/remote#delete
const remoteUploadJobDeleteAPIFormat = "/api/v2/remote/delete?job_id=%d&token=%s"

// CreateFromLinks create shared links based on the input original links.
func (p *RapidGator) CreateFromLinks(ctx context.Context, keepShareUserID string, originalLinks []string, createBy string) (sharedLinks map[string]*share.Share, err error) {
	defer func() {
		if err != nil {
			log.WithContext(ctx).Error("RapidGator CreateFromLinks err:", err)
		}
	}()

	if limitRateChecker.check() {
		// fmt.Println("RapidGator CreateFromLinks is limitted by rate.")
		return nil, nil
	}

	// todo: get account token by keepShareUserID;
	token, err := getToken()
	if err != nil {
		return nil, fmt.Errorf("RapidGator CreateFromLinks getToken err: %w", err)
	}

	// get existed jobs;
	remoteUploadJobsByOriginLink, _ := getRemoteUploadInfo(token)

	// only create files for new links.
	sharedLinks = map[string]*share.Share{}
	for _, link := range originalLinks {
		status := share.StatusNotFound
		if job, ok := remoteUploadJobsByOriginLink[strings.ToLower(link)]; ok {
			status = jobStateLabel2State(job.StateLabel)
		}

		switch status {
		case share.StatusNotFound, share.StatusDeleted:
			job, err := createRemoteUpload(token, link)
			if err != nil {
				return nil, err
			}
			sharedLinks[link] = &share.Share{
				State:          share.StatusCreated,
				Title:          job.Name,
				HostSharedLink: "",
				OriginalLink:   link,
				CreatedBy:      createBy,
				CreatedAt:      time.Now(),
				Size:           job.File.Size,
			}

		case share.StatusOK:
			job := remoteUploadJobsByOriginLink[strings.ToLower(link)]
			sharedLinks[link] = &share.Share{
				State:          status,
				Title:          job.Name,
				HostSharedLink: job.File.URL,
				OriginalLink:   link,
				CreatedBy:      createBy,
				CreatedAt:      time.Now(),
				Size:           job.File.Size,
			}

		case share.StatusCreated, share.StatusPending:
			job := remoteUploadJobsByOriginLink[strings.ToLower(link)]
			sharedLinks[link] = &share.Share{
				State:          share.StatusCreated,
				Title:          job.Name,
				HostSharedLink: "",
				OriginalLink:   link,
				CreatedBy:      createBy,
				CreatedAt:      time.Now(),
				Size:           job.File.Size,
			}

		default:
			fmt.Println("RapidGator CreateFromLinks, unknown share.Status:", status)
		}
	}

	return sharedLinks, nil
}

func createRemoteUpload(token string, remoteUploadURL string) (job *RemoteUploadJob, err error) {
	remoteUploadCreateURL := baseURL + fmt.Sprintf(remoteUploadCreateAPIFormat, remoteUploadURL, token)

	var remoteUploadCreateResponse struct {
		Status   int `json:"status"`
		Response struct {
			Jobs []RemoteUploadJob `json:"jobs"`
		} `json:"response"`
		Details string `json:"details"`
	}

	_, respErr := restyClient.R().
		SetResult(&remoteUploadCreateResponse).
		Post(remoteUploadCreateURL)

	// maybe unmarshal JSON failed;
	if respErr != nil {
		// fmt.Println("RapidGator createRemoteUpload respErr:", respErr)
	}
	// so check remoteUploadCreateResponse.Status;
	if remoteUploadCreateResponse.Status != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("RapidGator createRemoteUpload response status:%d, details:%s",
			remoteUploadCreateResponse.Status,
			remoteUploadCreateResponse.Details))
	}

	jobCount := len(remoteUploadCreateResponse.Response.Jobs)
	if jobCount != 1 {
		return nil, fmt.Errorf("RapidGator createRemoteUpload error job count:%d", jobCount)
	}

	job = &remoteUploadCreateResponse.Response.Jobs[0]

	return job, nil
}

func deleteRemoteUploadInfo(token string, jobID int) {
	remoteUploadJobDeleteURL := baseURL + fmt.Sprintf(remoteUploadJobDeleteAPIFormat, jobID, token)

	// resp, err := client.R().
	_, respErr := restyClient.R().
		Post(remoteUploadJobDeleteURL)

	if respErr != nil {
		fmt.Println("RapidGator DeleteRemoteUploadInfo respErr:", respErr)
	}
}
