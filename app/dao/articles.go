// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// articlesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type articlesDao struct {
	*internal.ArticlesDao
}

var (
	// Articles is globally public accessible object for table articles operations.
	Articles = &articlesDao{
		internal.Articles,
	}
)

// Fill with you ideas below.