package user

import (
	"database/sql"
	"luvletter/conf"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken ...
func GenerateToken(name string, admin bool) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		"Jon Snow", // Name
		true,       // Admin
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))
}

// SaveUser ...
func SaveUser(u NewUser) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`INSERT INTO user (account, nickname, password) VALUES (?, ?, ?)`)
	res, err := stmt.Exec(u.Account, u.NickName, u.Password)
	defer stmt.Close()
	_, err = res.LastInsertId()
	return err
}

// GetUserByID ...
func GetUserByID(id int16) (User, error) {
	var res User
	db, err := sql.Open("mysql", conf.DBConfig)
	row := db.QueryRow(`SELECT id, avator, account, nickname, password,create_time,update_time FROM user WHERE id=?`, id)
	err = row.Scan(&res.ID, &res.Avator, &res.Account, &res.Nickname, &res.Password, &res.CreateTime, &res.UpdateTime)
	return res, err
}

// GetUserByAccount ...
func GetUserByAccount(account string) (User, error) {
	var res User
	db, err := sql.Open("mysql", conf.DBConfig)
	row := db.QueryRow(`SELECT id, avator, account, nickname, password,create_time,update_time FROM user WHERE account=?`, account)
	err = row.Scan(&res.ID, &res.Avator, &res.Account, &res.Nickname, &res.Password, &res.CreateTime, &res.UpdateTime)
	return res, err
}

// UpdateUser ...
func UpdateUser(u User) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`UPDATE user SET avator=?,nickname=?,password=?,update_time=? WHERE id=?`)
	res, err := stmt.Exec(u.Avator, u.Nickname, u.Password, u.UpdateTime, u.ID)
	defer stmt.Close()
	_, err = res.LastInsertId()
	return err
}

// TrackUserAction ...
func TrackUserAction(track TrackAction) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`INSERT INTO trace (account, time, action, extra) VALUES (?, ?, ?, ?)`)
	res, err := stmt.Exec(track.Account, track.Time, track.Action, track.Extra)
	defer stmt.Close()
	_, err = res.LastInsertId()
	return err
}
