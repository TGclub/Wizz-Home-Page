// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// interviewersDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type interviewersDao struct {
	*internal.InterviewersDao
}

var (
	// Interviewers is globally public accessible object for table interviewers operations.
	Interviewers = &interviewersDao{
		internal.Interviewers,
	}
)

// Fill with you ideas below.
