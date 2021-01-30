// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// productsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type productsDao struct {
	*internal.ProductsDao
}

var (
	// Products is globally public accessible object for table products operations.
	Products = &productsDao{
		internal.Products,
	}
)

// Fill with you ideas below.