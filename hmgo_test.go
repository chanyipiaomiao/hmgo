package hmgo

import (
	"fmt"
	"testing"
)

type User struct {
	Username string
	Name     string
	Age      int
}

func TestInsert(t *testing.T) {

	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
	}

	docs := []*User{
		{
			Username: "zhangsan",
			Name:     "张三",
			Age:      18,
		},
		{
			Username: "lisi",
			Name:     "李四",
			Age:      20,
		},
		{
			Username: "wangwu",
			Name:     "王五",
			Age:      27,
		},
	}

	m := New("test", "user")
	for _, doc := range docs {
		if err := m.Save(NewObjectId(), doc); err != nil {
			t.Error(err)
			break
		}
	}

}

func TestQuery(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
	}

	m := New("test", "user")

	var user User
	if err := m.QueryOne(D{"username": "zhangsan"}, nil, &user); err != nil {
		t.Error(err)
	}
	fmt.Println(user)

	var users []*User
	if err := m.Query(nil, nil, &users); err != nil {
		t.Error(err)
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	var userp []*User
	page, err := m.QueryWithPage(nil, nil, &userp, 1, 1)
	if err != nil {
		t.Error(err)
	}
	for _, user := range userp {
		fmt.Printf("%+v\n", user)
	}
	fmt.Printf("%+v\n", page)
}

func TestUpdate(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
	}

	m := New("test", "user")
	if err := m.Update(D{"username": "zhangsan"}, D{"$set": D{"age": 28}}); err != nil {
		t.Error(err)
	}

}

func TestDelete(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
	}

	m := New("test", "user")
	if err := m.Delete(D{"username": "wangwu"}); err != nil {
		t.Error(err)
	}
}

func TestDeleteAll(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
	}

	m := New("test", "user")
	if err := m.DeleteAll(D{"username": "zhangsan"}); err != nil {
		t.Error(err)
	}
}

func TestMakeIndex(t *testing.T) {
	indexs := []Index{
		{
			DB:       "test",
			Table:    "user",
			Key:      []string{"username"},
			Unique:   true,
			DropDups: true,
			Sparse:   true,
		},
	}

	if err := InitMongo("127.0.0.1:27017", 10, indexs...); err != nil {
		t.Error(err)
	}

}
