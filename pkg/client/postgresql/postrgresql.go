package postgresql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/EXClub/test_project.git/internal/utils/config"
	"github.com/EXClub/test_project.git/pkg/utils/repeatable"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//файл инициализации клиента постгреса

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	/*
			DSN (Data Source Name) — это строка, используемая для определения источника данных,
		к которому вы хотите подключиться,в контексте подключения к базе данных.
			DSN содержит информацию о типе базы данных, местоположении, имени пользователя,
		пароле и других параметрах, необходимых для установления соединения с базой данных.
	*/

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	err = repeatable.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		// Подключаемся к пулу соединений
		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		// Если подключение успешно, возвращаем nil
		return nil
	}, maxAttempts, 5*time.Second)

	// Обработка ошибки после попыток подключения
	if err != nil {
		log.Fatal("error do with tries postgresql", err)
	}

	return pool, nil
}
