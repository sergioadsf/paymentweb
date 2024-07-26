package repository

import (
	"context"
	"log/slog"
	"paymentweb/internal/configuration"
	"paymentweb/internal/core/domain"
)

type UserRepository struct {
	db configuration.DBContext
}

func NewUserRepository(db configuration.DBContext) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Save(ctx *context.Context, user *domain.User) error {
	db := ur.db.Get("pagamento")
	// defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		slog.Warn("Não foi possível abrir transação para salvar usuário." + err.Error())
		return err
	}
	stmt, err := tx.Prepare("insert into users(id, name, cpf, email, phone) values ($1,$2,$3,$4,$5) ")
	if err != nil {
		slog.Warn("Não foi possível preparar usuário." + err.Error())
		return err
	}
	_, err = stmt.Exec(user.Id, user.Name, user.Cpf, user.Email, user.Phone)
	if err != nil {
		slog.Warn("Não foi possível inserir usuário." + err.Error())
		return err
	}
	if err := stmt.Close(); err != nil {
		slog.Warn("Não foi possível fechar Statemnt." + err.Error())
		return err
	}
	if err := tx.Commit(); err != nil {
		slog.Warn("Não foi possível finalizar transação." + err.Error())
		return err
	}
	return nil
}

func (ur *UserRepository) List(ctx *context.Context) ([]domain.User, error) {
	db := ur.db.Get("pagamento")
	// defer db.Close()

	rows, err := db.Query("select id, name, cpf, email, phone from users")
	if err != nil {
		slog.Warn("Não foi possível listar os clientes." + err.Error())
	}
	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name, &user.Cpf, &user.Email, &user.Phone)
		if err != nil {
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) ListUser(ctx *context.Context, id int32) (*domain.User, error) {

	return nil, nil
}
