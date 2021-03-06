// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

// Class represents a row from 'public.classes'.
type Class struct {
	ClassID    int         `json:"class_id"`   // class_id
	Name       string      `json:"name"`       // name
	StartDate  time.Time   `json:"start_date"` // start_date
	EndDate    pq.NullTime `json:"end_date"`   // end_date
	Mainstream bool        `json:"mainstream"` // mainstream

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Class exists in the database.
func (c *Class) Exists() bool {
	return c._exists
}

// Deleted provides information if the Class has been deleted from the database.
func (c *Class) Deleted() bool {
	return c._deleted
}

// Insert inserts the Class to the database.
func (c *Class) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.classes (` +
		`name, start_date, end_date, mainstream` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING class_id`

	// run query
	XOLog(sqlstr, c.Name, c.StartDate, c.EndDate, c.Mainstream)
	err = db.QueryRow(sqlstr, c.Name, c.StartDate, c.EndDate, c.Mainstream).Scan(&c.ClassID)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Class in the database.
func (c *Class) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.classes SET (` +
		`name, start_date, end_date, mainstream` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE class_id = $5`

	// run query
	XOLog(sqlstr, c.Name, c.StartDate, c.EndDate, c.Mainstream, c.ClassID)
	_, err = db.Exec(sqlstr, c.Name, c.StartDate, c.EndDate, c.Mainstream, c.ClassID)
	return err
}

// Save saves the Class to the database.
func (c *Class) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Class.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Class) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.classes (` +
		`class_id, name, start_date, end_date, mainstream` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (class_id) DO UPDATE SET (` +
		`class_id, name, start_date, end_date, mainstream` +
		`) = (` +
		`EXCLUDED.class_id, EXCLUDED.name, EXCLUDED.start_date, EXCLUDED.end_date, EXCLUDED.mainstream` +
		`)`

	// run query
	XOLog(sqlstr, c.ClassID, c.Name, c.StartDate, c.EndDate, c.Mainstream)
	_, err = db.Exec(sqlstr, c.ClassID, c.Name, c.StartDate, c.EndDate, c.Mainstream)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Class from the database.
func (c *Class) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.classes WHERE class_id = $1`

	// run query
	XOLog(sqlstr, c.ClassID)
	_, err = db.Exec(sqlstr, c.ClassID)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// ClassByClassID retrieves a row from 'public.classes' as a Class.
//
// Generated from index 'classes_pkey'.
func ClassByClassID(db XODB, classID int) (*Class, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`class_id, name, start_date, end_date, mainstream ` +
		`FROM public.classes ` +
		`WHERE class_id = $1`

	// run query
	XOLog(sqlstr, classID)
	c := Class{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, classID).Scan(&c.ClassID, &c.Name, &c.StartDate, &c.EndDate, &c.Mainstream)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
