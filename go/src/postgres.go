package main

import (
  "database/sql"
  _ "github.com/lib/pq"
  "time"
  "log"
  "fmt"
)

const (
  CONNECTION_STRING string = "postgres://postgres:postgres@postgres/test?sslmode=disable"
  SQL_CREATE_USER_TABLE string = `DROP TABLE IF EXISTS public.tbl_user;
                                  CREATE TABLE public.tbl_user (
                                    id bigserial PRIMARY KEY,
                                    username varchar(30),
                                    password varchar(30),
                                    cellphone character(11),
                                    gender character(1),
                                    age smallint,
                                    created_at timestamp WITH TIME ZONE,
                                    updated_at timestamp WITH TIME ZONE,
                                    last_login_at timestamp WITH TIME ZONE
                                  ) WITHOUT OIDS;
                                  ALTER TABLE public.tbl_user OWNER TO postgres;
                                  COMMENT ON TABLE public.tbl_user IS '用户表';
                                  COMMENT ON COLUMN public.tbl_user.id IS '编号';
                                  COMMENT ON COLUMN public.tbl_user.username IS '用户名';
                                  COMMENT ON COLUMN public.tbl_user.password IS '密码';
                                  COMMENT ON COLUMN public.tbl_user.cellphone IS '手机';
                                  COMMENT ON COLUMN public.tbl_user.gender IS '性别';
                                  COMMENT ON COLUMN public.tbl_user.age IS '年龄';
                                  COMMENT ON COLUMN public.tbl_user.created_at IS '创建时间';
                                  COMMENT ON COLUMN public.tbl_user.updated_at IS '更新时间';
                                  COMMENT ON COLUMN public.tbl_user.last_login_at IS '最后登录时间';`
  SQL_INSERT_USER string = `INSERT INTO tbl_user (username,
                                                  password,
                                                  cellphone,
                                                  gender,
                                                  age,
                                                  created_at,
                                                  updated_at,
                                                  last_login_at)
                            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
                            RETURNING id;`
  SQL_GET_USER string = `SELECT id,
                                username,
                                password,
                                cellphone,
                                gender,
                                age,
                                created_at,
                                updated_at,
                                last_login_at
                         FROM tbl_user
                         WHERE id = $1
                         OFFSET 0
                         LIMIT 1;`
  SQL_QUERY_USER_LIST string = `SELECT id,
                                  username,
                                  password,
                                  cellphone,
                                  gender,
                                  age,
                                  created_at,
                                  updated_at,
                                  last_login_at
                           FROM tbl_user
                           ORDER BY id ASC;`
  SQL_UPDATE_USER string = `UPDATE tbl_user
                            SET username = $1,
                                password = $2,
                                cellphone = $3,
                                gender = $4,
                                age = $5,
                                updated_at = $6
                            WHERE id = $7;`
)

func connect() *sql.DB {
  db, err := sql.Open("postgres", CONNECTION_STRING)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Connecting postgres successed!")
  return db
}

func createUserTable(db *sql.DB) {
  _, err := db.Exec(SQL_CREATE_USER_TABLE)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Create user table successed!")
  }
}

type User struct {
  Id int64
  Username string
  Password string
  Cellphone string
  Gender string
  Age int
  CreatedAt time.Time
  UpdatedAt time.Time
  LastLoginAt time.Time
}

func newUser(username string,
             password string,
             cellphone string,
             gender string,
             age int) *User {
  return &User {
    Username: username,
    Password: password,
    Cellphone: cellphone,
    Gender: gender,
    Age: age,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
}

func insertUser(db *sql.DB, user *User) {
  err := db.QueryRow(
    SQL_INSERT_USER, user.Username,
    user.Password,
    user.Cellphone,
    user.Gender,
    user.Age,
    user.CreatedAt,
    user.UpdatedAt,
    user.LastLoginAt,
  ).Scan(&user.Id)

  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("Insert user(Id: %d) successed!\n", user.Id)
}

func getUser(db *sql.DB, id int64) (*User, error) {
  user := &User {}
  err := db.QueryRow(SQL_GET_USER, id).
    Scan(
      &user.Id,
      &user.Username,
      &user.Password,
      &user.Cellphone,
      &user.Gender,
      &user.Age,
      &user.CreatedAt,
      &user.UpdatedAt,
      &user.LastLoginAt,
    )
  return user, err
}

func updateUser(db *sql.DB, user *User) {
  result, err := db.Exec(
    SQL_UPDATE_USER,
    user.Username,
    user.Password,
    user.Cellphone,
    user.Gender,
    user.Age,
    user.UpdatedAt,
    user.Id,
  )
  if err != nil {
    fmt.Println(err)
  }
  if rows, _ := result.RowsAffected(); rows == 1 {
    fmt.Println("Update successed!")
  }
}

func queryUserList(db *sql.DB) {
  rows, err := db.Query(SQL_QUERY_USER_LIST)
  if err != nil {
    fmt.Println(err)
    return
  }
  for rows.Next() {
    user := &User {}
    if err := rows.Scan(&user.Id,
                        &user.Username,
                        &user.Password,
                        &user.Cellphone,
                        &user.Gender,
                        &user.Age,
                        &user.CreatedAt,
                        &user.UpdatedAt,
                        &user.LastLoginAt); err != nil {
      fmt.Println(err)
      break
    }
    fmt.Println(user)
  }
}

func main() {
  db := connect()
  defer db.Close()

  //createUserTable(db)

  //user := &User {
  //  Username: "Michael",
  //  Password: "123",
  //  Cellphone: "13000000000",
  //  Gender: "M",
  //  Age: 21,
  //  CreatedAt: time.Now(),
  //  UpdatedAt: time.Now(),
  //}
  //insertUser(db, user)

  //user := newUser("Coral",
  //                "123",
  //                "15222222222",
  //                "F",
  //                22)
  //insertUser(db, user)

  queryUserList(db)

  user, err := getUser(db, 1)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Print("Get user(1):")
    fmt.Println(user)
  }

  user.Cellphone = "13666666666"
  user.UpdatedAt = time.Now()
  updateUser(db, user)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Print("Get user(1):")
    fmt.Println(user)
  }

  queryUserList(db)
}
