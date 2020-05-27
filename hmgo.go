package hmgo

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

var (
	mgoSession  *mgo.Session
	ErrNotFound = errors.New("not found")
)

type D map[string]interface{}

func getMgoSession() *mgo.Session {
	return mgoSession.Clone()
}

// NewObjectId 生成一个ObjectId
func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}

type Client struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

type Index struct {
	DB       string
	Table    string
	Key      []string
	Unique   bool
	DropDups bool
	Sparse   bool
}

func makeIndex(c *mgo.Collection, index *Index) error {
	idx := mgo.Index{
		Key:        index.Key,
		Unique:     index.Unique,
		DropDups:   index.DropDups,
		Background: true,
		Sparse:     index.Sparse,
	}

	if err := c.EnsureIndex(idx); err != nil {
		return err
	}

	return nil
}

func InitMongo(url string, poolSize int, indexes ...Index) error {
	var err error

	if url == "" {
		return fmt.Errorf("url can not be empty")
	}

	if mgoSession, err = mgo.DialWithTimeout(url, 10*time.Second); err != nil {
		return err
	}

	mgoSession.SetMode(mgo.Monotonic, true)
	mgoSession.SetPoolLimit(poolSize)

	c := getMgoSession()
	for _, index := range indexes {
		collection := c.DB(index.DB).C(index.Table)
		if err = makeIndex(collection, &index); err != nil {
			return err
		}
	}
	c.Close()

	return nil
}

func New(db, table string) *Client {
	var (
		session    *mgo.Session
		collection *mgo.Collection
	)
	session = getMgoSession()
	collection = session.DB(db).C(table)
	return &Client{Session: session, Collection: collection}
}

func (c *Client) Close() {
	c.Session.Close()
}

func (c *Client) Save(docId, doc interface{}) error {
	if _, err := c.Collection.UpsertId(docId, bson.M{"$set": doc}); err != nil {
		return err
	}

	return nil
}

func (c *Client) Query(query, selector, records interface{}) error {
	if err := c.Collection.Find(query).Select(selector).All(records); err != nil {
		return err
	}

	return nil
}

func (c *Client) QueryOne(query, selector, records interface{}) error {
	if err := c.Collection.Find(query).Select(selector).One(records); err != nil {
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (c *Client) QueryWithPage(query, selector, records interface{},
	pageNo, pageSize int, sortFields ...string) (*Page, error) {
	var (
		err   error
		total int
		q     *mgo.Query
		page  Page
	)

	q = c.Collection.Find(query).Select(selector)
	if total, err = q.Count(); err != nil {
		return nil, err
	}

	if len(sortFields) > 0 {
		q = q.Sort(sortFields...)
	}

	page = PageUtil(total, pageNo, pageSize)
	if err = q.Skip((pageNo - 1) * pageSize).Limit(pageSize).All(records); err != nil {
		return nil, err
	}

	return &page, nil
}

func (c *Client) Update(selector, update interface{}) error {
	if err := c.Collection.Update(selector, update); err != nil {
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (c *Client) Delete(selector interface{}) error {
	if err := c.Collection.Remove(selector); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteAll(selector interface{}) error {
	if _, err := c.Collection.RemoveAll(selector); err != nil {
		return err
	}
	return nil
}
