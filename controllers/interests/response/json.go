package response

import (
	"weatherit/usecases/interests"
)

type Interest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromDomain(domain interests.Domain) Interest {
	return Interest{
		ID:   domain.ID,
		Name: domain.Name,
	}
}
