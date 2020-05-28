## hmgo
hmgo是基于 https://github.com/globalsign/mgo 的封装

## 示例

### 插入

```go
package main

import (
	"fmt"
	"github.com/chanyipiaomiao/hmgo"
)

type User struct {
	Username string
	Name     string
	Age      int
}


func main() {

	if err := hmgo.InitMongo("127.0.0.1:27017", 5); err != nil {
		fmt.Println(err)
		return
	}

	users := []*User{
		{Username: "zhangsan",Name:"张三",Age:18},
		{Username: "lisi",Name: "李四", Age: 20},
		{Username: "wangwu",Name: "王五", Age: 27},
		{Username: "mutouliu",Name: "木头六", Age: 37},
	}

	m := hmgo.New("test", "user")
	defer m.Close()

	for _, user := range users {
		if err := m.Save(hmgo.NewObjectId(), user); err != nil {
			fmt.Println(err)
			break
		}
	}
}
```

### 查询

#### 查询单个

```go
var user User
if err := m.QueryOne(hmgo.D{"username": "zhangsan"}, nil, &user); err != nil {
    fmt.Println(err)
    return
}
fmt.Println(user)
```

#### 查询多个

```go
var users []*User
if err := m.Query(nil, nil, &users); err != nil {
    fmt.Println(err)
    return
}
for _, user := range users {
    fmt.Printf("%+v\n", user)
}
```

#### 分页查询

```go
var userp []*User
page, err := m.QueryWithPage(nil, nil, &userp, 1, 1)
if err != nil {
    fmt.Println(err)
    return
}
for _, user := range userp {
    fmt.Printf("%+v\n", user)
}
fmt.Printf("%+v\n", page)
```

#### 更新

更新1条
```go
if err := m.UpdateOne(hmgo.D{"username": "zhangsan"}, hmgo.D{"$set": hmgo.D{"age": 38}}); err != nil {
    fmt.Println(err)
}
```

更新多条
```go
if err := m.UpdateMany(nil, D{"$set": D{"address": "上海"}}); err != nil {
    t.Error(err)
}
```

#### 删除

删除1条
```go
if err := m.DeleteOne(hmgo.D{"username": "wangwu"}); err != nil {
    fmt.Println(err)
}
```
删除多条

```go
if err := m.DeleteMany(hmgo.D{"address": "北京"}); err != nil {
    fmt.Println(err)
}
```

#### 创建索引

在初始化的时候创建索引

```go
indexs := []hmgo.Index{
    {
        DB:       "test",
        Table:    "user",
        Key:      []string{"username"},
        Unique:   true,
        DropDups: true,
        Sparse:   true,
    },
    {
        DB:       "test2",
        Table:    "user",
        Key:      []string{"username"},
        Unique:   true,
        DropDups: true,
        Sparse:   true,
    },
}

if err := hmgo.InitMongo("127.0.0.1:27017", 5, indexs...); err != nil {
    fmt.Println(err)
    return
}
```