package svc

import (
	"gopherss/db"
	"gopherss/model"
)

type EntriesService struct{}

var entriesRepository = db.EntriesRepository{}

func (s EntriesService) GetMany(input model.ListEntriesInput) ([]model.Entry, error) {
	entries, err := entriesRepository.GetMany(input)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (s EntriesService) Get(id int) (*model.Entry, error) {
	entry, err := entriesRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (s EntriesService) Update(id int, input model.UpdateEntryInput) (bool, error) {
	updated, err := entriesRepository.Update(id, input)
	if err != nil {
		return false, err
	}
	return updated, nil
}

func (s EntriesService) UpdateMany(input model.UpdateEntriesInput) error {
	err := entriesRepository.UpdateMany(input)
	if err != nil {
		return err
	}
	return nil
}

func (s EntriesService) Mark(id int, input model.MarkEntryInput) (bool, error) {
	updated, err := entriesRepository.Mark(id, input)
	if err != nil {
		return false, err
	}
	return updated, nil
}

func (s EntriesService) MarkMany(input model.MarkEntriesInput) error {
	err := entriesRepository.MarkMany(input)
	if err != nil {
		return err
	}
	return nil
}
