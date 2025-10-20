package request

import ()

type StoreStudent struct {
	StudentID string `json:"student_id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Address string `json:"address"`
	EntryYear uint `json:"entry_year"`
}

type UpdateStudent struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Address string `json:"address"`
	EntryYear uint `json:"entry_year"`
}