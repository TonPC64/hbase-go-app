package repository

import (
	"context"
	"fmt"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

// HbaseClient client
type HbaseClient struct {
	client gohbase.Client
	table  string
	cf     string
}

// NewHbase constuct hbase client
func NewHbase(host, table, columnFamily string) *HbaseClient {
	return &HbaseClient{
		client: gohbase.NewClient(host),
		table:  table,
		cf:     columnFamily,
	}
}

func (h HbaseClient) GetRow(row string) {
	getRequest, err := hrpc.NewGetStr(context.Background(), h.table, row)
	getRsp, err := h.client.Get(getRequest)

	fmt.Println(getRsp, err)
}

func (h HbaseClient) GetCell(row, cell string) {
	family := map[string][]string{h.cf: []string{cell}}
	getRequest, err := hrpc.NewGetStr(context.Background(), h.table, row, hrpc.Families(family))
	getRsp, err := h.client.Get(getRequest)

	fmt.Println(getRsp, err)
}

func (h HbaseClient) Put(row, cell string, value []byte) {
	values := map[string]map[string][]byte{
		h.cf: map[string][]byte{
			cell: value,
		},
	}
	putRequest, err := hrpc.NewPutStr(context.Background(), h.table, row, values)
	rsp, err := h.client.Put(putRequest)

	fmt.Println(rsp, err)
}
