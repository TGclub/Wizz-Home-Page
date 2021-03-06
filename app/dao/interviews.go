// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// interviewsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type interviewsDao struct {
	*internal.InterviewsDao
}

var (
	// Interviews is globally public accessible object for table interviews operations.
	Interviews = &interviewsDao{
		internal.Interviews,
	}
)

// Fill with you ideas below.
