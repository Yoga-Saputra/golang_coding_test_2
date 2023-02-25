package members

import (
	"errors"
)

type Service interface {
	GetMembers() ([]Member, error)
	InsertMember(input InputMember) (Member, error)
	GetMemberById(input int) (Member, error)
	UpdateMember(input InputMember, IDMember int) (Member, error)
	DeleteMember(input int) (Member, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMembers() ([]Member, error) {

	member, err := s.repository.FindAll()

	if err != nil {
		return member, err
	}

	return member, nil
}

func (s *service) InsertMember(input InputMember) (Member, error) {
	member := Member{
		Gender:    input.Gender,
		Username:  input.Username,
		Skintype:  input.Skintype,
		Skincolor: input.Skincolor,
	}

	newMember, err := s.repository.Insert(member)

	if err != nil {
		return newMember, err
	}

	return newMember, nil
}

func (s *service) GetMemberById(input int) (Member, error) {
	member, err := s.repository.FindById(input)
	if err != nil {
		return member, err
	}

	if member.IDMember == 0 {
		return member, errors.New("no member found on that id")
	}

	return member, nil
}

func (s *service) UpdateMember(input InputMember, inputIDMember int) (Member, error) {
	member, err := s.repository.FindById(inputIDMember)

	if err != nil {
		return member, err
	}

	if member.IDMember == 0 {
		return member, errors.New("no member found on that id")
	}

	member.Username = input.Username
	member.Gender = input.Gender
	member.Skintype = input.Skintype
	member.Skincolor = input.Skincolor

	updateMember, err := s.repository.Update(member)

	if err != nil {
		return updateMember, err
	}

	return updateMember, nil
}

func (s *service) DeleteMember(input int) (Member, error) {

	member, err := s.repository.FindById(input)

	if err != nil {
		return member, err
	}

	if member.IDMember == 0 {
		return member, errors.New("no member found on that id")
	}

	_, err = s.repository.Delete(input)

	if err != nil {
		return member, err
	}

	return member, nil
}
