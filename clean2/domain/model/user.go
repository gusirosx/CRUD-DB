package model

/*
Domain encapsulate enterprise wide business rules.

An entity (or domain) can be an object with methods, or it can be a set
of data structures and functions. It doesn’t matter so long as the entities
could be used by many different applications in the enterprise.
If you don’t have an enterprise, and are just writing a single application,
then these entities are the business objects of the application.
They encapsulate the most general and high-level rules. They are the least
likely to change when something external changes. No operational change to
any particular application should affect the entity layer.

*/

type User struct {
	Id     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}

// type User struct {
// 	ID        uint       `gorm:"primary_key" json:"id"`
// 	Name      string     `json:"name"`
// 	Age       string     `json:"age"`
// 	CreatedAt *time.Time `json:"created_at"`
// 	UpdatedAt *time.Time `json:"updated_at"`
// 	DeletedAt *time.Time `json:"deleted_at"`
// }
