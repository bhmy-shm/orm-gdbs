package examples

import (
	"context"
	"fmt"
	"github.com/bhmy-shm/dbsdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
	"time"
)

func QueryTest() {

	// 客户端构建器
	c,_:=builder.NewClientBuilder("localhost:8082").WithOption(grpc.WithInsecure()).Build()
	// 参数构建器
	paramBuilder:=builder.NewParamBuilder().Add("id",1)
	// api 构建器
	api:=builder.NewApiBuilder("deptlist",builder.APITYPE_QUERY)

	// 查询结果集
	depts:=make([]*Dept,0)

	//执行 和调用  API
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()

	err:=api.Invoke(ctx,paramBuilder,c,&depts)
	if err!=nil{
		log.Fatal(err)
	}
	for _,dept:=range depts{
		fmt.Println(dept)
	}
}
