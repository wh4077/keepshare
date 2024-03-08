// // Copyright 2024 The KeepShare Authors. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file.

package rapidgator

import (
	"context"
	"fmt"

	"github.com/KeepShareOrg/keepshare/hosts"
	"github.com/KeepShareOrg/keepshare/pkg/share"
)

// RapidGator official website: https://rapidgator.net/
type RapidGator struct {
	*hosts.Dependencies

	// q   *query.Query
	// m   *account.Manager
	// api *api.API
}

// //go:embed  rawsql/*.sql
// var sqlFS embed.FS

func init() {
	// sql, err := hosts.ReadSQLFileFromFS(sqlFS)
	// if err != nil {
	// 	// panic(fmt.Errorf("read sql files err: %w", err))
	// 	fmt.Println(fmt.Errorf("read sql files err: %w", err))
	// }
	sql := []string{"test"}

	hosts.Register(&hosts.Properties{Name: "rapidgator", New: New, CreateTableStatements: sql})
}

// New create a RapidGator host.
func New(d *hosts.Dependencies) hosts.Host {
	p := &RapidGator{Dependencies: d}

	// p.q = query.Use(p.Mysql)

	// p.api = api.New(p.q, d)

	// p.m = account.NewManager(p.q, p.api, d)

	// go p.deleteFilesBackground()

	// d.Queue.RegisterHandler(taskTypeSyncWorkerInfo, asynq.HandlerFunc(p.syncWorkerHandler))

	fmt.Println("RapidGator host is initialized.")

	return p
}

// GetStatistics return the statistics of each host shared link.
func (p *RapidGator) GetStatistics(ctx context.Context, userID string, hostSharedLinks []string) (details map[string]share.Statistics, err error) {
	fmt.Println("RapidGator GetStatistics is not implemented yet.")

	return
}

// Delete delete shared links by original links.
func (p *RapidGator) Delete(ctx context.Context, userID string, originalLinks []string) error {
	fmt.Println("RapidGator Delete is not implemented yet.")

	return nil
}

func (p *RapidGator) AddEventListener(event hosts.EventType, callback hosts.ListenerCallback) {
	fmt.Println("RapidGator AddEventListener is not implemented yet.")

	return
}
