package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DBResults to store database operation results
type DBResults struct {
	LastInsertID int64
	AffectedRows int64
}

type database struct {
	ctx context.Context
	db *sqlx.DB
	log *log.Logger
}

func NewDatabase(ctx context.Context, log *log.Logger) *database {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")
	u := url.URL{
		User:     url.UserPassword("keycloak", "password"),
		Host:     "localhost:5432",
		Path:     "wisdom",
		RawQuery: q.Encode(),
	}
	
	db, err := sqlx.Connect("postgres", "postgres:"+u.String())
	if err != nil {
		log.Printf("[DataBase] err %v", err)
	}
	
	log.Printf("[DataBase] Connected to database %v", u.String())
	return &database{
		ctx: ctx,
		db: db,
		log: log,
	}
}

func (d *database) GetDB() *sqlx.DB {
	return d.db
}

// NamedExecContext is a helper function to execute a CUD operation with
// logging and tracing.
func NamedExecContext(ctx context.Context, log *log.Logger, db sqlx.ExtContext, query string, data interface{}) (DBResults, error) {
	q := queryString(query, data)
	log.Printf("[DataBase] Query: %v", q)

	var dbres DBResults
	res, err := sqlx.NamedExecContext(ctx, db, query, data)
	if err != nil {
		return DBResults{}, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Println("<DataBase> LastInsertId error: ", err)
		//return DBResults{}, err
	}
	dbres.LastInsertID = lid

	ra, err := res.RowsAffected()
	if err != nil {
		log.Println("<DataBase> RowsAffected error: ", err)
		// return DBResults{}, err
	}
	dbres.AffectedRows = ra

	return dbres, err
}

// NamedQuerySlice is a helper function for executing queries that return a
// collection of data to be unmarshalled into a slice.
func NamedQuerySlice(ctx context.Context, log *log.Logger, db sqlx.ExtContext, query string, data interface{}, dest interface{}) error {
	q := queryString(query, data)
	log.Printf("[DataBase] Query: %v", q)
	
	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return errors.New("must provide a pointer to a slice")
	}

	rows, err := sqlx.NamedQueryContext(ctx, db, query, data)
	if err != nil {
		return err
	}
	defer rows.Close() //nolint:all

	slice := val.Elem()
	for rows.Next() {
		v := reflect.New(slice.Type().Elem())
		if err := rows.StructScan(v.Interface()); err != nil && !strings.Contains(err.Error(), "unsupported Scan, storing driver.Value type <nil> into type *json.RawMessage") {
			return err
		}
		slice.Set(reflect.Append(slice, v.Elem()))
	}

	return nil
}

// NamedQueryStruct is a helper function for executing queries that return a
// single value to be unmarshalled into a struct type.
func NamedQueryStruct(ctx context.Context, log *log.Logger, db sqlx.ExtContext, query string, data interface{}, dest interface{}) error {
	q := queryString(query, data)
	log.Printf("[DataBase] Query: %v", q)

	rows, err := sqlx.NamedQueryContext(ctx, db, query, data)
	if err != nil {
		return err
	}
	defer rows.Close() //nolint:all

	if !rows.Next() {
		return errors.New("no rows returned")
	}

	if err := rows.StructScan(dest); err != nil && !strings.Contains(err.Error(), "unsupported Scan, storing driver.Value type <nil> into type *json.RawMessage") {
		return err
	}

	return nil
}

// queryString provides a pretty print version of the query and parameters.
func queryString(query string, args ...interface{}) string {
	if args[0] == nil {
		return query
	}
	query, params, err := sqlx.Named(query, args)
	if err != nil {
		return err.Error()
	}

	for _, param := range params {
		var value string
		switch v := param.(type) {
		case *string:
			value = fmt.Sprintf("%v", v)
			if v != nil {
				value = fmt.Sprintf(`'%s'`, *v)
			}
		case string, []byte:
			value = fmt.Sprintf(`'%s'`, v)
		case json.RawMessage:
			value = fmt.Sprintf(`'%s'`, string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}
		query = strings.Replace(query, "?", value, 1)
	}

	singleSpacePattern := regexp.MustCompile(`\s\s+`)
	query = singleSpacePattern.ReplaceAllString(query, " ")

	return strings.Trim(query, " ")
}
