package server

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/entities"
	entitiesGrpcGen "github.com/ischenkx/innotech-backend/services/entities/implementation/grpc/pb/generated"
	"github.com/ischenkx/innotech-backend/services/entities/service"
)

type Server struct {
	entitiesGrpcGen.UnimplementedEntitiesServer
	service *service.Service
}


func (s *Server) Create(ctx context.Context, req *entitiesGrpcGen.CreateEntityRequest) (*entitiesGrpcGen.Entity, error) {
	ent, err := s.service.Create(ctx, entities.CreateEntityDto{
		Title:     req.Title,
		ShortDesc: req.ShortDesc,
		LongDesc:  req.LongDesc,
		MoneyGoal: req.MoneyGoal,
		OwnerID:   req.OwnerId,
	})

	if err != nil {
		return nil, err
	}

	return &entitiesGrpcGen.Entity{
		Id:        ent.ID,
		Title:     ent.Title,
		ShortDesc: ent.ShortDesc,
		LongDesc:  ent.LongDesc,
		MoneyGoal: ent.MoneyGoal,
		OwnerId:   ent.OwnerID,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *entitiesGrpcGen.DeleteEntityRequest) (*entitiesGrpcGen.Entity, error) {
	ent, err := s.service.Delete(ctx, entities.DeleteEntityDto{
		ID:      req.Id,
		OwnerID: req.OwnerId,
	})

	if err != nil {
		return nil, err
	}

	return &entitiesGrpcGen.Entity{
		Id:        ent.ID,
		Title:     ent.Title,
		ShortDesc: ent.ShortDesc,
		LongDesc:  ent.LongDesc,
		MoneyGoal: ent.MoneyGoal,
		OwnerId:   ent.OwnerID,
	}, nil
}

func (s *Server) Update(ctx context.Context, req *entitiesGrpcGen.UpdateEntityRequest) (*entitiesGrpcGen.Entity, error) {
	ent, err := s.service.Update(ctx, entities.UpdateEntityDto{
		ID:        req.Id,
		Title:     req.Title,
		ShortDesc: req.ShortDesc,
		LongDesc:  req.LongDesc,
		MoneyGoal: req.MoneyGoal,
		OwnerID:   req.OwnerId,
	})

	if err != nil {
		return nil, err
	}

	return &entitiesGrpcGen.Entity{
		Id:        ent.ID,
		Title:     ent.Title,
		ShortDesc: ent.ShortDesc,
		LongDesc:  ent.LongDesc,
		MoneyGoal: ent.MoneyGoal,
		OwnerId:   ent.OwnerID,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *entitiesGrpcGen.GetEntityRequest) (*entitiesGrpcGen.EntityArray, error) {
	arr, err := s.service.Get(ctx, entities.GetEntitiesDto{
		IDs:       req.Ids.Ids,
		IsPreview: req.IsPreview,
	})

	if err != nil {
		return nil, err
	}

	var entArr entitiesGrpcGen.EntityArray

	for _, info := range arr {
		entArr.Entities = append(entArr.Entities, &entitiesGrpcGen.Entity{
			Id:        info.ID,
			Title:     info.Title,
			ShortDesc:info.ShortDesc,
			LongDesc:  info.LongDesc,
			MoneyGoal: info.MoneyGoal,
			OwnerId:   info.OwnerID,
		})
	}

	return &entArr, nil
}

func (s *Server) GetByOwnerID(ctx context.Context, req *entitiesGrpcGen.GetEntityByOwnerIDRequest) (*entitiesGrpcGen.EntityArray, error) {
	arr, err := s.service.GetByOwnerID(ctx, entities.GetEntitiesByOwnerIdDto{
		OwnerID:   req.OwnerId,
		IsPreview: req.IsPreview,
		Offset:    req.Offset,
		Limit:     req.Limit,
	})

	if err != nil {
		return nil, err
	}

	var entArr entitiesGrpcGen.EntityArray

	for _, info := range arr {
		entArr.Entities = append(entArr.Entities, &entitiesGrpcGen.Entity{
			Id:        info.ID,
			Title:     info.Title,
			ShortDesc:info.ShortDesc,
			LongDesc:  info.LongDesc,
			MoneyGoal: info.MoneyGoal,
			OwnerId:   info.OwnerID,
		})
	}

	return &entArr, nil
}

func New(srv *service.Service) *Server {
	return &Server{
		service: srv,
	}
}