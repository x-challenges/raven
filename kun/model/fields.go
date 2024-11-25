package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

// ID
type ID = string

// DateTime
type DateTime = time.Time

/*
	YDB support

	1.	Json
	2.	JsonString
	3.	JsonStrings

	4. JsonValuer ->
	5. JsonScanner ->

*/

func JSONScanner[T any](e T, src interface{}) error {
	return json.Unmarshal(src.([]byte), &e)
}

func JSONValuer[T any](e T) (driver.Value, error) {
	val, err := json.Marshal(e)
	return types.JSONValue(string(val)), err
}
