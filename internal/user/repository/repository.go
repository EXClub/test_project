package repository

type PostgresUserRepo struct {
}

type UserRepository interface {
	SaveUser() error
	GetHash() []byte
	GetToken() []byte
}

func (r *PostgresUserRepo) NewPostgresUserRepo() *PostgresUserRepo {
	return &PostgresUserRepo{}
}

func (r *PostgresUserRepo) SaveUser(username string, passwordHash string) error {
	// Логика сохранения юзера в БД
	return nil
}

func (r *PostgresUserRepo) GetHash() []byte {
	// Логика получения хэша пароля юзера
	return []byte("some hash")
}

func (r *PostgresUserRepo) GetToken() []byte {
	// Логика получения токена юзера из БД
	return []byte("some token")
}
