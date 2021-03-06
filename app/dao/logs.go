// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// logsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type logsDao struct {
	*internal.LogsDao
}

var (
	// Logs is globally public accessible object for table logs operations.
	Logs = &logsDao{
		internal.Logs,
	}
)

// Fill with you ideas below.
