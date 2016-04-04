package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Event struct {
	Id            string  `db:"id"`
	Kind          string  `db:"kind"`
	Status        string  `db:"status"`
	TriggerTag    string  `db:"trigger_tag"`
	TriggerEvent  string  `db:"trigger_event"`
	TransactionId string  `db:"transaction_id"`
	Data          *string `db:"data"`
}

func (db *Database) GetEvents(status string) ([]Event, error) {
	var events []Event
	err := db.selectObjs(&events, "SELECT * FROM teleport.event WHERE status = $1 ORDER BY id ASC;", status)
	return events, err
}

func (e *Event) UpdateQuery(tx *sqlx.Tx) {
	tx.MustExec(
		"UPDATE teleport.event SET status = $1, data = $2 WHERE id = $3;",
		e.Status,
		e.Data,
		e.Id,
	)
}

// Implement Stringer
func (e *Event) String() string {
	return fmt.Sprintf(
		"%s,%s,%s,%s,%s,%s",
		e.Id,
		e.Kind,
		e.TriggerTag,
		e.TriggerEvent,
		e.TransactionId,
		*e.Data,
	)
}

