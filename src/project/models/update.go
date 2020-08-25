package models

import (
	"context"
	"fmt"
	"strconv"
)

type Update struct {
	id int64
	//key string
}

func NewUpdate(userId int64, body string) (*Update, error) {
	id, err := client.Incr(context.TODO(), "update:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("update:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(context.TODO(), key, "id", id)
	pipe.HSet(context.TODO(), key,"user_id", userId)
	pipe.HSet(context.TODO(), key,"body", body)
	//pipe.HSet(context.TODO(), "user:by-username", username, id)
	pipe.LPush(context.TODO(), "updates", id)
	pipe.LPush(context.TODO(), fmt.Sprintf("user:%d:updates", userId), id)
	_, err = pipe.Exec(context.TODO())
	if err != nil {
		return nil, err
	}
	//return &Update{key}, nil
	return &Update{id}, nil
}

func (update *Update) GetBody() (string, error) {
	key := fmt.Sprintf("update:%d", update.id)
	//return client.HGet(context.TODO(), update.key, "body").Result()
	return client.HGet(context.TODO(), key, "body").Result()
}

func (update *Update) GetUser() (*User, error) {
	key := fmt.Sprintf("update:%d", update.id)
	//userId, err := client.HGet(context.TODO(), update.key, "user_id").Int64()
	userId, err := client.HGet(context.TODO(), key, "user_id").Int64()
	if err != nil {
		return nil, err
	}
	return GetUserById(userId)
}

//func GetComments() ([]string, error) {
//	return client.LRange(context.TODO(),"comments", 0, 10).Result()
//}
//
//func PostComment(comment string) error {
//	return client.LPush(context.TODO(), "comments", comment).Err()
//}

func queryUpdates(key string) ([]*Update, error) {
	updateIds, err := client.LRange(context.TODO(),key, 0, 10).Result()
	if err != nil {
		return nil, err
	}
	updates := make([]*Update, len(updateIds))
	for i, strId := range updateIds {
		//key := "update:" + id
		id, err := strconv.Atoi(strId)
		if err != nil {
			return nil, err
		}
		//updates[i] = &Update{key}
		updates[i] = &Update{int64(id)}
	}
	//return client.LRange(context.TODO(),"updates", 0, 10).Result()
	return updates, err
}

func GetAllUpdates() ([]*Update, error) {
	return queryUpdates("updates")
	//updateIds, err := client.LRange(context.TODO(),"updates", 0, 10).Result()
	//if err != nil {
	//	return nil, err
	//}
	//updates := make([]*Update, len(updateIds))
	//for i, id := range updateIds {
	//	key := "update:" + id
	//	updates[i] = &Update{key}
	//}
	////return client.LRange(context.TODO(),"updates", 0, 10).Result()
	//return updates, err
}

func GetUpdates(userId int64) ([]*Update, error) {
	key := fmt.Sprintf("user:%d:updates", userId)
	return queryUpdates(key)
	//updateIds, err := client.LRange(context.TODO(),key, 0, 10).Result()
	//if err != nil {
	//	return nil, err
	//}
	//updates := make([]*Update, len(updateIds))
	//for i, id := range updateIds {
	//	key := "update:" + id
	//	updates[i] = &Update{key}
	//}
	////return client.LRange(context.TODO(),"updates", 0, 10).Result()
	//return updates, err
}

func PostUpdate(userId int64, body string) error {
	//return client.LPush(context.TODO(), "updates", body).Err()
	_, err := NewUpdate(userId, body)
	return err
}