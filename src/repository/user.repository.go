package repository

type User struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (r Repository) CreateUser(email string, password string) (int, error) {
	insert, err := r.dbPool.Exec(`
		INSERT INTO "user" (email, password) 
		VALUES ($1, $2)
	`, email, password)
	if err != nil {
		return 0, err
	}

	affected, err := insert.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

func (r Repository) FindUserByEmail(email string) (User, error) {
	var user User
	rows, err := r.dbPool.Query(`
		SELECT id, email, password 
		FROM "user" 
		WHERE email=$1 AND deleted_at is null 
		LIMIT 1
	`, email)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (r Repository) UpdatePassword(id string, password string) (int, error) {
	res, err := r.dbPool.Exec(`
		UPDATE "user" 
		SET password=$1 
		where id=$2
	`, password, id)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

func (r Repository) SoftDelete(email string) (int, error) {
	res, err := r.dbPool.Exec(`
		UPDATE "user"
		SET deleted_at=NOW()
		WHERE email=$1;
	`, email)

	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}
