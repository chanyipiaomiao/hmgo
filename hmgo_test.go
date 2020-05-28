package hmgo

import (
	"fmt"
	"testing"
)

type User struct {
	Username string
	Name     string
	Age      int
	Address  string
}

func TestInsert(t *testing.T) {

	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}

	users := []*User{
		{Username: "zhangsan", Name: "张三", Age: 18, Address: "北京"},
		{Username: "lisi", Name: "李四", Age: 20, Address: "上海"},
		{Username: "wangwu", Name: "王五", Age: 27, Address: "北京"},
		{Username: "mutouliu", Name: "木头六", Age: 37, Address: "上海"},
	}
	m := New("test2", "user")
	for _, user := range users {
		if err := m.Save(NewObjectId(), user); err != nil {
			t.Error(err)
			break
		}
	}

}

func TestQueryOne(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}
	m := New("test", "user")
	var user User
	if err := m.QueryOne(D{"username": "zhangsan"}, nil, &user); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(user)
}

func TestQuery(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}

	m := New("test", "user")

	var users []*User
	if err := m.Query(nil, nil, &users); err != nil {
		t.Error(err)
		return
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

}

func TestQueryWithPage(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}
	m := New("test", "user")
	var userp []*User
	page, err := m.QueryWithPage(nil, nil, &userp, 1, 1)
	if err != nil {
		t.Error(err)
		return
	}
	for _, user := range userp {
		fmt.Printf("%+v\n", user)
	}
	fmt.Printf("%+v\n", page)
}

func TestUpdateOne(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}

	m := New("test", "user")
	if err := m.UpdateOne(D{"username": "zhangsan"}, D{"$set": D{"age": 28}}); err != nil {
		t.Error(err)
	}

}

func TestUpdateMany(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}

	m := New("test", "user")
	if err := m.UpdateMany(nil, D{"$set": D{"address1": "上海"}}); err != nil {
		t.Error(err)
	}

}

func TestDeleteOne(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}

	m := New("test", "user")
	if err := m.DeleteOne(D{"username": "wangwu"}); err != nil {
		t.Error(err)
	}
}

func TestDeleteMany(t *testing.T) {
	if err := InitMongo("127.0.0.1:27017", 10); err != nil {
		t.Error(err)
		return
	}

	m := New("test", "user")
	if err := m.DeleteMany(D{"username": "zhangsan"}); err != nil {
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
