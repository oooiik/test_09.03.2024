package service

import (
	"github.com/oooiik/test_09.03.2024/internal/http/request"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"github.com/oooiik/test_09.03.2024/internal/model"
	"github.com/oooiik/test_09.03.2024/internal/repository"
)

type Good interface {
	Index(p request.GoodIndex) ([]*model.Good, error)
	Show(id uint32) (*model.Good, error)
	Create(req request.GoodCreate) (*model.Good, error)
	Update(req request.GoodUpdate) (*model.Good, error)
	Delete(id uint32) (*model.Good, error)

	RePrioritize(req request.GoodRePrioritize) ([]*model.Good, error)
}

type good struct {
	repository repository.Good
}

func NewGood(r repository.Good) Good {
	return &good{
		repository: r,
	}
}

func (s good) Index(req request.GoodIndex) ([]*model.Good, error) {
	list, err := s.repository.ListWithPagination(req.Limit, req.Offset)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return list, nil
}

func (s good) Show(id uint32) (*model.Good, error) {
	//TODO implement me
	return nil, nil
}

func (s good) Create(req request.GoodCreate) (*model.Good, error) {
	m := &model.Good{}
	m.Fill(&model.Good{
		ProjectId:   req.ProjectId,
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
	})

	c, err := s.repository.Create(m)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return c, nil
}

func (s good) Update(req request.GoodUpdate) (*model.Good, error) {
	m, err := s.repository.GetById(req.Id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	m.Fill(&model.Good{
		ProjectId:   req.ProjectId,
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
	})
	u, err := s.repository.Update(m)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s good) Delete(id uint32) (*model.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (s good) RePrioritize(req request.GoodRePrioritize) ([]*model.Good, error) {
	//TODO implement me
	panic("implement me")
}
