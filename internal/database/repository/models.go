// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

type Link struct {
	ID       int64  `json:"id"`
	Location string `json:"location"`
	Slug     string `json:"slug"`
}
