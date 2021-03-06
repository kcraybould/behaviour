// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// MissedClass represents a row from 'public.missed_classes'.
type MissedClass struct {
	MissedID   int       `json:"missed_id"`   // missed_id
	MissedDate time.Time `json:"missed_date"` // missed_date
	ClassID    int       `json:"class_id"`    // class_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MissedClass exists in the database.
func (mc *MissedClass) Exists() bool {
	return mc._exists
}

// Deleted provides information if the MissedClass has been deleted from the database.
func (mc *MissedClass) Deleted() bool {
	return mc._deleted
}

// Insert inserts the MissedClass to the database.
func (mc *MissedClass) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if mc._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.missed_classes (` +
		`missed_date, class_id` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING missed_id`

	// run query
	XOLog(sqlstr, mc.MissedDate, mc.ClassID)
	err = db.QueryRow(sqlstr, mc.MissedDate, mc.ClassID).Scan(&mc.MissedID)
	if err != nil {
		return err
	}

	// set existence
	mc._exists = true

	return nil
}

// Update updates the MissedClass in the database.
func (mc *MissedClass) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mc._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if mc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.missed_classes SET (` +
		`missed_date, class_id` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE missed_id = $3`

	// run query
	XOLog(sqlstr, mc.MissedDate, mc.ClassID, mc.MissedID)
	_, err = db.Exec(sqlstr, mc.MissedDate, mc.ClassID, mc.MissedID)
	return err
}

// Save saves the MissedClass to the database.
func (mc *MissedClass) Save(db XODB) error {
	if mc.Exists() {
		return mc.Update(db)
	}

	return mc.Insert(db)
}

// Upsert performs an upsert for MissedClass.
//
// NOTE: PostgreSQL 9.5+ only
func (mc *MissedClass) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if mc._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.missed_classes (` +
		`missed_id, missed_date, class_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (missed_id) DO UPDATE SET (` +
		`missed_id, missed_date, class_id` +
		`) = (` +
		`EXCLUDED.missed_id, EXCLUDED.missed_date, EXCLUDED.class_id` +
		`)`

	// run query
	XOLog(sqlstr, mc.MissedID, mc.MissedDate, mc.ClassID)
	_, err = db.Exec(sqlstr, mc.MissedID, mc.MissedDate, mc.ClassID)
	if err != nil {
		return err
	}

	// set existence
	mc._exists = true

	return nil
}

// Delete deletes the MissedClass from the database.
func (mc *MissedClass) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mc._exists {
		return nil
	}

	// if deleted, bail
	if mc._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.missed_classes WHERE missed_id = $1`

	// run query
	XOLog(sqlstr, mc.MissedID)
	_, err = db.Exec(sqlstr, mc.MissedID)
	if err != nil {
		return err
	}

	// set deleted
	mc._deleted = true

	return nil
}

// Class returns the Class associated with the MissedClass's ClassID (class_id).
//
// Generated from foreign key 'missed_classes_class_id_fkey'.
func (mc *MissedClass) Class(db XODB) (*Class, error) {
	return ClassByClassID(db, mc.ClassID)
}

// MissedClassByMissedID retrieves a row from 'public.missed_classes' as a MissedClass.
//
// Generated from index 'missed_classes_pkey'.
func MissedClassByMissedID(db XODB, missedID int) (*MissedClass, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`missed_id, missed_date, class_id ` +
		`FROM public.missed_classes ` +
		`WHERE missed_id = $1`

	// run query
	XOLog(sqlstr, missedID)
	mc := MissedClass{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, missedID).Scan(&mc.MissedID, &mc.MissedDate, &mc.ClassID)
	if err != nil {
		return nil, err
	}

	return &mc, nil
}
