package examples

import (
	"context"
	"github.com/bhmy-shm/orm-gdbs/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type ProdModel struct {
	Id        int32   `mapstructure:"id,omitempty"`           //商品ID
	ProdName  string  `mapstructure:"prod_name,omitempty"`   //商品名
	ProdPrice float32 `mapstructure:"prod_price,omitempty"` //价格
}

func QueryProds() {

	// 客户端构建器
	c, _ := builder.NewClientBuilder("localhost:8082").WithOption(grpc.WithInsecure()).Build()
	// 参数构建器
	paramBuilder := builder.NewParamBuilder()
	// api 构建器
	api := builder.NewApiBuilder("getProdList", builder.APITYPE_QUERY)

	// 查询结果集
	depts := make([]*ProdModel, 0)

	err := api.Invoke(context.Background(), paramBuilder, c, &depts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("depts", depts[0])
}
