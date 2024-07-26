package repository

import (
	"context"
	"log/slog"
	"paymentweb/internal/configuration"
	"paymentweb/internal/core/domain"

	"github.com/google/uuid"
)

type CardRepository struct {
	db configuration.DBContext
}

func NewCardRepository(db configuration.DBContext) *CardRepository {
	return &CardRepository{db}
}

func (c *CardRepository) Save(ctx *context.Context, idUser uuid.UUID, card *domain.Card) error {
	db := c.db.Get("pagamento")
	tx, err := db.Begin()
	if err != nil {
		slog.Warn("Não foi possível abrir transação para salvar em CardRepository." + err.Error())
		return err
	}
	stmt, err := tx.Prepare("insert into cards(id, id_user, brand, number, alias, exp_year, exp_month) values ($1,$2,$3,$4,$5,$6,$7) ")
	if err != nil {
		slog.Warn("Não foi possível preparar em CardRepository." + err.Error())
		return err
	}
	id, err := uuid.NewUUID()
	if err != nil {
		slog.Warn("Não foi possível gerar id para em CardRepository." + err.Error())
		return err
	}
	_, err = stmt.Exec(id, idUser, card.Brand, card.Number, card.Alias, card.ExpYear, card.ExpMonth)
	if err != nil {
		slog.Warn("Não foi possível inserir em CardRepository." + err.Error())
		return err
	}
	if err := stmt.Close(); err != nil {
		slog.Warn("Não foi possível fechar Statemnt em em CardRepository." + err.Error())
		return err
	}
	if err := tx.Commit(); err != nil {
		slog.Warn("Não foi possível finalizar transação em CardRepository." + err.Error())
		return err
	}
	return nil
}

func (ur *CardRepository) List(ctx *context.Context) ([]domain.Card, error) {
	db := ur.db.Get("pagamento")
	// defer db.Close()

	rows, err := db.Query("select id, brand, alias, number, exp_year, exp_month  from cards")
	if err != nil {
		slog.Warn("Não foi possível listar os cartões." + err.Error())
	}
	var cards []domain.Card
	for rows.Next() {
		var card domain.Card
		err := rows.Scan(&card.Id, &card.Brand, &card.Alias, &card.Number, &card.ExpYear, &card.ExpMonth)
		if err != nil {
			slog.Warn(err.Error())
			continue
		}
		cards = append(cards, card)
	}

	return cards, nil
}
