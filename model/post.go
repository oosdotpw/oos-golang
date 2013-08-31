package model

import (
	"labix.org/v2/mgo/bson"
	"oos-go/db"
	"time"
)

type PostModel struct {
	ObjectId   bson.ObjectId `_id,omitempty`
	UserID     bson.ObjectId `user_id`
	Content    string        `content`
	CreateTime time.Time     `create_time`
	Replys     []ReplyModel  `replys,omitempty`
}

type ReplyModel struct {
	ObjectId   bson.ObjectId `_id,omitempty`
	UserID     bson.ObjectId `user_id`
	Content    string        `content`
	CreateTime time.Time     `create_time`
}

type MarkModel struct {
	ObjectId bson.ObjectId "_id,omitempty"
	UserID   bson.ObjectId "user_id"
	PostID   bson.ObjectId "post_id"
	Type     string        "type"
}

func InsertPost(userID bson.ObjectId, content string) bson.ObjectId {
	m := PostModel{
		ObjectId:   bson.NewObjectId(),
		UserID:     userID,
		Content:    content,
		CreateTime: time.Now(),
		Replys:     []ReplyModel{},
	}

	err := Post.Insert(m)
	if err != nil {
		panic(err)
	}

	return m.ObjectId
}

func ExistPost(idStr string) bool {
	return db.Exist(Post, bson.M{"_id": bson.ObjectIdHex(idStr)})
}

func GetPost(postIDstr string) PostModel {
	m := PostModel{}

	postID := bson.ObjectIdHex(postIDstr)

	err := Post.Find(bson.M{"_id": postID}).One(&m)
	if err != nil {
		panic(err)
	}

	return m
}

func InsertReply(userID bson.ObjectId, postIDstr string, content string) bson.ObjectId {
	postID := bson.ObjectIdHex(postIDstr)

	m := ReplyModel{
		ObjectId:   bson.NewObjectId(),
		UserID:     userID,
		Content:    content,
		CreateTime: time.Now(),
	}

	err := Post.UpdateId(postID, bson.M{"$push": bson.M{"replys": m}})
	if err != nil {
		panic(err)
	}

	return m.ObjectId
}

func GetReplys(postIDstr string) []ReplyModel {
	m := PostModel{}

	postID := bson.ObjectIdHex(postIDstr)

	sortStr := "replys.create_time"

	err := Post.FindId(postID).Sort(sortStr).One(&m)
	if err != nil {
		panic(err)
	}

	return m.Replys
}

func InsertMark(userID bson.ObjectId, postIDstr string, markType string) {
	postID := bson.ObjectIdHex(postIDstr)

	findM := bson.M{
		"user_id": userID,
		"post_id": postID,
	}
	changeM := bson.M{"$set": bson.M{"type": markType}}

	_, err := Mark.Upsert(findM, changeM)
	if err != nil {
		panic(err)
	}

}

func FetchNewest(maxnum int) []PostModel {
	postsM := []PostModel{}

	err := Post.Find(bson.M{}).Sort("-create_time").Limit(maxnum).All(&postsM)
	if err != nil {
		panic(err)
	}

	return postsM
}

func FetchNewer(postIDstr string, maxnum int) []PostModel {
	postM := PostModel{}
	postsM := []PostModel{}

	postID := bson.ObjectIdHex(postIDstr)

	err := Post.Find(bson.M{"_id": postID}).One(&postM)
	if err != nil {
		panic(err)
	}

	findM := bson.M{"create_time": bson.M{"$gt": postM.CreateTime}}

	err = Post.Find(findM).Sort("-create_time").Limit(maxnum).All(&postsM)
	if err != nil {
		panic(err)
	}

	return postsM
}

func FetchOlder(postIDstr string, maxnum int) []PostModel {
	postM := PostModel{}
	postsM := []PostModel{}

	postID := bson.ObjectIdHex(postIDstr)

	err := Post.Find(bson.M{"_id": postID}).One(&postM)
	if err != nil {
		panic(err)
	}

	findM := bson.M{"create_time": bson.M{"$lt": postM.CreateTime}}

	err = Post.Find(findM).Sort("-create_time").Limit(maxnum).All(&postsM)
	if err != nil {
		panic(err)
	}

	return postsM
}
