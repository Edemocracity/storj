// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package mutlinodedb

import (
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/multinode"
	"storj.io/storj/private/dbutil"
	"storj.io/storj/private/dbutil/pgutil"
)

var (
	// ensures that multinodeDB implements multinode.DB.
	_ multinode.DB = (*multinodeDB)(nil)

	// Error is the default multinodedb errs class.
	Error = errs.Class("multinodedb internal error")
)

// multinodeDB combines access to different database tables with a record
// of the db driver, db implementation, and db source URL.
// Implementation of multinode.DB interface.
//
// architecture: Master Database
type multinodeDB struct {
	//*dbx.DB

	log            *zap.Logger
	driver         string
	implementation dbutil.Implementation
	source         string
}

// Close closes database.
func (m multinodeDB) Close() error {
	// TODO: will be implemented in dbx.
	panic("implement me")
}

// New creates instance of database supports postgres.
func New(log *zap.Logger, databaseURL string) (multinode.DB, error) {
	driver, source, implementation, err := dbutil.SplitConnStr(databaseURL)
	if err != nil {
		return nil, err
	}
	// TODO: do we need cockroach implementation?
	if implementation != dbutil.Postgres && implementation != dbutil.Cockroach {
		return nil, Error.New("unsupported driver %q", driver)
	}

	source = pgutil.CheckApplicationName(source)

	// dbxDB, err := dbx.Open(driver, source)
	// if err != nil {
	// 	return nil, Error.New("failed opening database via DBX at %q: %v",
	// 		source, err)
	// }
	// log.Debug("Connected to:", zap.String("db source", source))

	// dbutil.Configure(dbxDB.DB, "multinodedb", mon)

	core := &multinodeDB{
		// DB: dbxDB,

		log:            log,
		driver:         driver,
		implementation: implementation,
		source:         source,
	}

	return core, nil
}
