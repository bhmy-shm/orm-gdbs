package helpers

import (
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

//程序员在囧途(www.jtthink.com)出品 咨询群：98514334
func PbstructsToMapList(in []*structpb.Struct) []map[string]interface{} {
	list := make([]map[string]interface{}, len(in))
	for i, item := range in {
		log.Println("asMap",item.AsMap())
		list[i] = item.AsMap()
	}
	return list
}
