package builder

import (
	"context"
	"github.com/bhmy-shm/dbsdk/pbfiles"
	"google.golang.org/grpc"
)

type ClientBuilder struct {

	url string
	opts []grpc.DialOption
}
func NewClientBuilder(url string ) *ClientBuilder {
	return &ClientBuilder{url:url}
}
func(this *ClientBuilder) WithOption(opts ...grpc.DialOption) *ClientBuilder{
	this.opts=append(this.opts,opts...)
	return this
}
func(this *ClientBuilder) Build() ( pbfiles.DBServiceClient,error ){
	client,err:=grpc.DialContext(context.Background(),
		this.url,
		this.opts...
	)
	if err!=nil{
		return nil ,err
	}
	return pbfiles.NewDBServiceClient(client),nil

}

