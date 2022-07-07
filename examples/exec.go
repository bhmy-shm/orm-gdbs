package examples

import (
	"context"
	"fmt"
	"github.com/bhmy-shm/orm-gdbs/pbfiles"
	"github.com/bhmy-shm/orm-gdbs/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type UserAddResult struct {
	UserID       int `mapstructure:"user_id"`
	RowsAffected int `mapstructure:"_RowsAffected"`
}

//程序员在囧途(www.jtthink.com)出品 咨询群：98514334
func ExecTest() {

	client, _ := grpc.DialContext(context.Background(),
		"localhost:8082",
		grpc.WithInsecure(),
	)
	c := pbfiles.NewDBServiceClient(client)

	//构建 参数
	paramBuilder := builder.NewParamBuilder().
		Add("userName", "shenyi").
		Add("userPass", "123456")
	//构建API对象
	api := builder.NewApiBuilder("adduser", builder.APITYPE_EXEC)

	//构建一个 结果集对象----必须是 地址
	ret := &UserAddResult{}

	err := api.Invoke(context.Background(), paramBuilder, c, ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

}

//程序员在囧途(www.jtthink.com)出品 咨询群：98514334
