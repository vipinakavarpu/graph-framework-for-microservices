// Code generated by github.com/vmware-tanzu/graph-framework-for-microservices/gqlgen, DO NOT EDIT.

package chat

import (
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
