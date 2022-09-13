package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
)

func main() {
	fmt.Println("Hello")
	client, err := echo.NewClient("echo", client.WithHostPorts("[::1]:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {

		/*
			req := &api.Request{}
		*/

		/*
			type Request struct {
				SilentClientCount *SilentClientCount `thrift:"SilentClientCount,1,optional" json:"SilentClientCount,omitempty"`
			}

			type SilentClientCount struct {
				Count     *int64      `thrift:"count,1,optional" json:"count,omitempty"`
				IdcCounts []*IDCCount `thrift:"idcCounts,2,optional" json:"idcCounts,omitempty"`
			}
		*/

		// var count int64
		// count = 4
		// sc := &api.SilentClientCount{
		// 	// Count:     &count,
		// 	IdcCounts: []*api.IDCCount{},
		// }

		req := &api.Request{
			// SilentClientCount: sc,
		}
		log.Printf("kitex req: %v", req)

		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("resp: %v", resp)
		time.Sleep(3 * time.Second)
	}
}
