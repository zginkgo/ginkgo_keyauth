package impl

import (
	"context"
	"github.com/zginkgo/ginkgo_keyauth/apps"
	"github.com/zginkgo/ginkgo_keyauth/apps/book"
	"github.com/zginkgo/ginkgo_keyauth/conf"
	"testing"
)

var (
	ins book.ServiceServer
)

func TestQueryBook(t *testing.T) {
	ss, err := ins.QueryBook(context.Background(), book.NewQueryBookRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestDescribeBook(t *testing.T) {
	ss, err := ins.DescribeBook(context.Background(), book.NewDescribeBookRequest("cd18qoeg26u22ffcon30"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestCreateBook(t *testing.T) {
	req := book.NewCreateBookRequest()
	req.CreateBy = "youmen3"
	req.Name = "三体"
	req.Author = "刘慈溪"
	ss, err := ins.CreateBook(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	if err := apps.InitAllApp(); err != nil {
		panic(err)
	}

	ins = apps.GetGrpcApp(book.AppName).(book.ServiceServer)
}
