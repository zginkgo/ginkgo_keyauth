package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/infraboard/mcube/exception"
	"github.com/zginkgo/ginkgo_keyauth/apps/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *impl) save(ctx context.Context, ins *token.Token) error {
	// s.col.InsertMany()
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			ins.AccessToken, err)
	}
	return nil
}

// GET, Describe
// filter 过滤器(Collection), 类似于MySQL, Where条件
// 调用Decode方法来进行 反序列化 bytes ---> Object (通过JSON Tag)
//func (s *impl) get(ctx context.Context, accessToken string) (*token.Token, error) {
//	filter := bson.M{"_id": accessToken}
//
//	ins := token.NewDefaultToken()
//	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
//		if err == mongo.ErrNoDocuments {
//			return nil, exception.NewNotFound("access token %s not found", accessToken)
//		}
//		return nil, exception.NewInternalServerError("find access token %s error, %s", accessToken, err)
//	}
//	return ins, nil
//}

// 带缓存逻辑
func (s *impl) get(ctx context.Context, accessToken string) (*token.Token, error) {
	// 1. 先从缓存中获取
	resp, err := s.redis.Get(accessToken).Bytes()
	if err != nil {
		s.log.Warnf("redis get error, %s", err)
	} else {
		tk := token.NewDefaultToken()
		if err := json.Unmarshal(resp, tk); err != nil {
			s.log.Warnf("unmarshal redis data error, %s", err)
		}
	}

	// 直接从数据库获取
	tk, err := s.getFromDB(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	// 设置环境
	jtk, err := json.Marshal(tk)
	if err != nil {
		s.log.Warnf("marshal token error, %s", err)
	}
	s.redis.Set(tk.AccessToken, string(jtk), 600)

	return tk, nil
}

// GET, Describe
// filter 过滤器(Collection),类似于MYSQL Where条件
// 调用Decode方法来进行 反序列化  bytes ---> Object (通过BSON Tag)
func (s *impl) getFromDB(ctx context.Context, accessToken string) (*token.Token, error) {
	filter := bson.M{"_id": accessToken}

	ins := token.NewDefaultToken()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("access token %s not found", accessToken)
		}

		return nil, exception.NewInternalServerError("find access token %s error, %s", accessToken, err)
	}

	return ins, nil
}

func (s *impl) delete(ctx context.Context, ins *token.Token) error {
	if ins == nil || ins.AccessToken == "" {
		return fmt.Errorf("access token is nil")
	}

	// delete from access token where id = ?
	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.AccessToken})
	if err != nil {
		return exception.NewInternalServerError("delete access token(%s) error, %s", ins.AccessToken, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("access token %s not found", ins.AccessToken)
	}

	return nil
}

// UpdateByID, 通过主键来更新对象
func (s *impl) update(ctx context.Context, ins *token.Token) error {
	// SQL update obj(SET f=v,f=v) where id=?
	// s.col.UpdateOne(ctx, filter(), ins)
	if _, err := s.col.UpdateByID(ctx, ins.AccessToken, ins); err != nil {
		return exception.NewInternalServerError("update token(%s) document error, %s",
			ins.AccessToken, err)
	}

	return nil
}
