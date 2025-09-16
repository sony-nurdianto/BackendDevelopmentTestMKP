package repo

import (
	"context"
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"

	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/concurent"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/constants"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/data"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/models"
	"github.com/sony-nurdianto/BackendDevelopmentTest/skill_test/service/gRPC/auth/internal/pbgen"
)

const (
	InsertUserStmtType      = "InsertUserStmtType"
	GetUserByCardIDStmtType = "GetUserByCardIDStmtType"
	GetUserByEmailStmtType  = "GetUserByEmailStmtType"
)

type AuthRepo interface {
	InsertUser(ctx context.Context, data *pbgen.RegisterUserRequest) (usr models.User, _ error)
	GetUserByCardID(ctx context.Context, cardID string) (usr models.User, _ error)
	GetUserByEmail(ctx context.Context, email string) (usr models.User, _ error)
}

type authRepo struct {
	insertUserStmt      *sql.Stmt
	getUserByCardIDStmt *sql.Stmt
	getUserByEmailStmt  *sql.Stmt
}

type userStmt struct {
	stmtType string
	stmt     *sql.Stmt
}

func insertUserStmt(ctx context.Context, db *sql.DB) <-chan any {
	out := make(chan any, 1)
	go func() {
		defer close(out)
		var res concurent.Result[userStmt]

		stmt, err := db.Prepare(constants.QueryInsertUser)
		if err != nil {
			res.Error = err
			concurent.SendResult(ctx, out, res)
			return
		}

		res.Value.stmtType = InsertUserStmtType
		res.Value.stmt = stmt
		concurent.SendResult(ctx, out, res)
	}()
	return out
}

func getUserByCardIDStmt(ctx context.Context, db *sql.DB) <-chan any {
	out := make(chan any, 1)
	go func() {
		defer close(out)
		var res concurent.Result[userStmt]
		stmt, err := db.Prepare(constants.QueryGetUserByID)
		if err != nil {
			res.Error = err
			concurent.SendResult(ctx, out, res)
			return
		}

		res.Value.stmtType = GetUserByCardIDStmtType
		res.Value.stmt = stmt
		concurent.SendResult(ctx, out, res)
	}()
	return out
}

func getUserByEmailStmt(ctx context.Context, db *sql.DB) <-chan any {
	out := make(chan any, 1)
	go func() {
		defer close(out)
		var res concurent.Result[userStmt]
		stmt, err := db.Prepare(constants.QueryGetUserByEmail)
		if err != nil {
			res.Error = err
			concurent.SendResult(ctx, out, res)
			return
		}

		res.Value.stmtType = GetUserByEmailStmtType
		res.Value.stmt = stmt
		concurent.SendResult(ctx, out, res)
	}()
	return out
}

func NewAuthRepo(ctx context.Context) (ar authRepo, _ error) {
	dbAddr := os.Getenv("DATABASE_ADDR")

	db, err := data.OpenPostgres(ctx, dbAddr)
	if err != nil {
		return ar, err
	}

	chs := []<-chan any{
		insertUserStmt(ctx, db),
		getUserByCardIDStmt(ctx, db),
		getUserByEmailStmt(ctx, db),
	}

	for v := range concurent.FanIn(ctx, chs...) {
		res, ok := v.(concurent.Result[userStmt])
		if !ok {
			return ar, errors.New("wrong data type prepare repo")
		}

		if res.Error != nil {
			return ar, res.Error
		}

		switch res.Value.stmtType {
		case InsertUserStmtType:
			ar.insertUserStmt = res.Value.stmt
		case GetUserByCardIDStmtType:
			ar.getUserByCardIDStmt = res.Value.stmt
		case GetUserByEmailStmtType:
			ar.getUserByEmailStmt = res.Value.stmt
		}
	}

	return ar, nil
}

func (ar authRepo) InsertUser(
	ctx context.Context,
	data *pbgen.RegisterUserRequest,
) (usr models.User, _ error) {
	row := ar.insertUserStmt.QueryRowContext(
		ctx,
		data.GetCardId(),
		data.GetFullName(),
		data.GetEmail(),
		data.GetPhone(),
		data.GetBalance(),
	)

	if row.Err() != nil {
		return usr, row.Err()
	}

	if err := row.Scan(
		&usr.ID,
		&usr.CardID,
		&usr.FullName,
		&usr.Email,
		&usr.Phone,
		&usr.Balance,
		&usr.Status,
		&usr.CurrentTripTerminalID,
		&usr.RegisteredAt,
		&usr.UpdatedAt,
	); err != nil {
		return usr, err
	}

	return usr, nil
}

func (ar authRepo) GetUserByEmail(
	ctx context.Context,
	email string,
) (usr models.User, _ error) {
	row := ar.getUserByEmailStmt.QueryRowContext(ctx, email)
	if row.Err() != nil {
		return usr, row.Err()
	}

	if err := row.Scan(
		&usr.ID,
		&usr.CardID,
		&usr.FullName,
		&usr.Email,
		&usr.Phone,
		&usr.Balance,
		&usr.Status,
		&usr.CurrentTripTerminalID,
		&usr.RegisteredAt,
		&usr.UpdatedAt,
	); err != nil {
		return usr, err
	}

	return usr, nil
}

func (ar authRepo) GetUserByCardID(
	ctx context.Context,
	cardID string,
) (usr models.User, _ error) {
	row := ar.getUserByCardIDStmt.QueryRowContext(ctx, cardID)
	if row.Err() != nil {
		return usr, row.Err()
	}

	if err := row.Scan(
		&usr.ID,
		&usr.CardID,
		&usr.FullName,
		&usr.Email,
		&usr.Phone,
		&usr.Balance,
		&usr.Status,
		&usr.CurrentTripTerminalID,
		&usr.RegisteredAt,
		&usr.UpdatedAt,
	); err != nil {
		return usr, err
	}

	return usr, nil
}
