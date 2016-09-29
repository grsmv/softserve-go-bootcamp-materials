package main

import

//"strconv"

"gopkg.in/mgo.v2/bson"

type dbMongoDB struct {
	session    *mgo.Session
	dbName     string
	collection string
}

func (db *dbMongoDB) Create(t Task) error {
	c := db.session.DB(db.dbName).C(db.collection)

	// TODO: add id generation
	//hexStr := bson.NewObjectId().Hex()
	//t.Id = strconv.ParseInt(hexStr, 16, 64)

	err := c.Insert(&t)
	return err
}

func (db *dbMongoDB) ReadById(id *int64) (TaskList, error) {

	var tasks TaskList

	c := db.session.DB(db.dbName).C(db.collection)
	err := c.Find(bson.M{"Id": id}).All(&tasks)
	return tasks, err

}

func (db *dbMongoDB) ReadByAlias(alias *string) (TaskList, error) {
	var tasks TaskList

	c := db.session.DB(db.dbName).C(db.collection)
	err := c.Find(bson.M{"Alias": alias}).All(&tasks)
	return tasks, err

}

func (db *dbMongoDB) Update(t Task) error {

	c := db.session.DB(db.dbName).C(db.collection)
	err := c.Update(bson.M{"Id": t.Id}, &t)

	return err

}

func (db *dbMongoDB) Delete(t Task) error {
	c := db.session.DB(db.dbName).C(db.collection)
	err := c.Remove(bson.M{"Id": t.Id})

	return err
}

func (db *dbMongoDB) ReadAll() (TaskList, error) {
	var tasks TaskList

	c := db.session.DB(db.dbName).C(db.collection)
	err := c.Find(nil).All(&tasks)

	return tasks, err
}
