package service

import (
	"context"
	"log/slog"
	"paymentweb/internal/core/domain"
	"paymentweb/internal/core/ports/in"
	"paymentweb/internal/core/ports/out"

	"github.com/google/uuid"
)

type UserService struct {
	repo        out.UserRepository
	cardService in.CardService
}

func NewUserService(repo out.UserRepository, cardService in.CardService) *UserService {
	return &UserService{repo, cardService}
}

func (us *UserService) Save(ctx *context.Context, user *domain.User) error {
	slog.Info("Chegou usuário " + user.Name + " cartão " + user.Card.Brand)
	id, err := uuid.NewUUID()
	if err != nil {
		slog.Warn("Não foi possível gerar id para usuário." + err.Error())
		return err
	}
	user.Id = &id
	if err := us.repo.Save(ctx, user); err != nil {
		slog.Warn("Não foi possível salvar usuário." + err.Error())
		return err
	}
	if err := us.cardService.Save(ctx, *user.Id, user.Card); err != nil {
		slog.Warn("Não foi possível salvar usuário." + err.Error())
		return err
	}
	return nil
}

func (us *UserService) List(ctx *context.Context) ([]domain.User, error) {
	users, err := us.repo.List(ctx)
	if err != nil {
		slog.Warn("Não foi possível encontar um usuário com o ID(%is)")
		return nil, err
	}

	return users, nil
}

func (us *UserService) ListUser(ctx *context.Context, id int32) (*domain.User, error) {
	_, err := us.repo.ListUser(ctx, id)
	if err != nil {
		slog.Warn("Não foi possível encontar um usuário com o ID(%is)")
		return nil, err
	}

	return nil, nil
}
