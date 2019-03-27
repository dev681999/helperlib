package helperlib

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
)

var (
	// ErrDBNoID is returned when no ID field or id tag is found in the struct.
	ErrDBNoID = storm.ErrNoID

	// ErrDBZeroID is returned when the ID field is a zero value.
	ErrDBZeroID = storm.ErrZeroID

	// ErrDBBadType is returned when a method receives an unexpected value type.
	ErrDBBadType = storm.ErrBadType

	// ErrDBAlreadyExists is returned uses when trying to set an existing value on a field that has a unique index.
	ErrDBAlreadyExists = storm.ErrAlreadyExists

	// ErrDBNilParam is returned when the specified param is expected to be not nil.
	ErrDBNilParam = storm.ErrNilParam

	// ErrDBUnknownTag is returned when an unexpected tag is specified.
	ErrDBUnknownTag = storm.ErrUnknownTag

	// ErrDBIdxNotFound is returned when the specified index is not found.
	ErrDBIdxNotFound = storm.ErrIdxNotFound

	// ErrDBSlicePtrNeeded is returned when an unexpected value is given, instead of a pointer to slice.
	ErrDBSlicePtrNeeded = storm.ErrSlicePtrNeeded

	// ErrDBStructPtrNeeded is returned when an unexpected value is given, instead of a pointer to struct.
	ErrDBStructPtrNeeded = storm.ErrStructPtrNeeded

	// ErrDBPtrNeeded is returned when an unexpected value is given, instead of a pointer.
	ErrDBPtrNeeded = storm.ErrPtrNeeded

	// ErrDBNoName is returned when the specified struct has no name.
	ErrDBNoName = storm.ErrNoName

	// ErrDBNotFound is returned when the specified record is not saved in the bucket.
	ErrDBNotFound = storm.ErrNotFound

	// ErrDBNotInTransaction is returned when trying to rollback or commit when not in transaction.
	ErrDBNotInTransaction = storm.ErrNotInTransaction

	// ErrDBIncompatibleValue is returned when trying to set a value with a different type than the chosen field
	ErrDBIncompatibleValue = storm.ErrIncompatibleValue
)

// Persistence is used for persisting data
type Persistence struct {
	DBConn *storm.DB `json:"dbConn,omitempty"`
	dbName string
}

// NewPersistence returns a new Persistence handler
func NewPersistence(dbName string) *Persistence {
	return &Persistence{
		dbName: dbName,
	}
}

// Connect establishes connection with the database
func (p *Persistence) Connect() error {
	var err error
	p.DBConn, err = storm.Open(p.dbName)
	return err
}

// Init inits database for data
func (p *Persistence) Init(data interface{}) error {
	return p.DBConn.Init(data)
}

// Close closes the database connection
func (p *Persistence) Close() {
	if p.DBConn != nil {
		p.DBConn.Close()
	}
}

// Save saves the data in the database
func (p *Persistence) Save(data interface{}) error {
	return p.DBConn.Save(data)
}

// One returns one record by the specified index
func (p *Persistence) One(fieldName string, value interface{}, to interface{}) error {
	return p.DBConn.One(fieldName, value, to)
}

// Find returns one or more records by the specified index
func (p *Persistence) Find(fieldName string, value interface{}, to interface{}) error {
	return p.DBConn.Find(fieldName, value, to)
}

// AllByIndex gets all the records from that are indexed in the specified index
func (p *Persistence) AllByIndex(fieldName string, to interface{}) error {
	return p.DBConn.AllByIndex(fieldName, to)
}

// All gets all the records from database
func (p *Persistence) All(to interface{}) error {
	return p.DBConn.All(to)
}

// Delete deletes the data from database
func (p *Persistence) Delete(data interface{}) error {
	return p.DBConn.DeleteStruct(data)
}

// DeleteByField deletes the data from database matching the field values
func (p *Persistence) DeleteByField(field string, value interface{}, dataInterface interface{}) error {
	return p.DBConn.Select(q.Eq(field, value)).Delete(dataInterface)
}

// Update data in database
func (p *Persistence) Update(data interface{}) error {
	return p.DBConn.Update(data)
}

// UpdateField updates a single field
func (p *Persistence) UpdateField(data interface{}, fieldName string, value interface{}) error {
	return p.DBConn.UpdateField(data, fieldName, value)
}

// ClearDB deletes all records
func (p *Persistence) ClearDB(data interface{}) error {
	return p.DBConn.Drop(data)
}

// Prefix finds all records whose specified fields start with given prefix
func (p *Persistence) Prefix(field, prefix string, to interface{}) error {
	return p.DBConn.Prefix(field, prefix, to)
}

// GetKey gets the value for the given key in specified bucket
func (p *Persistence) GetKey(bucket string, key, to interface{}) error {
	return p.DBConn.Get(bucket, key, to)
}

// SetKey sets the value for the given key in specified bucket
func (p *Persistence) SetKey(bucket string, key, value interface{}) error {
	return p.DBConn.Set(bucket, key, value)
}

// DeleteKey deletes the key in specified bucket
func (p *Persistence) DeleteKey(bucket string, key interface{}) error {
	return p.DBConn.Delete(bucket, key)
}
