package go_micro_srv_user

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.Must(uuid.NewV4()).String()
	return scope.SetColumn("id", uuid)
}

/*
What happen here?
 1. change our ORM's behaviour to generate a UUID on creation, instead of trying to generate an integer ID.
 2. In case you didn't know, a UUID is a randomly generated set of hyphenated strings, used as an ID or primary key. This is more secure than just using auto-incrementing ID's, because it stops people from guessing or traversing through your API endpoints.
 3. MongoDB already uses a variation of this, but we need to tell our Postgres models to use UUID's.
 4. Hooks into GORM's event lifecycle so that we generate a UUID for our Id column, before the entity is saved.
*/
