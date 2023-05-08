package tests

import (
	"encoding/json"
	"fmt"

	"testing"

	"github.com/chnykn/bimface/v3/bean/request"
)

func TestQueryDSL(t *testing.T) {

	dsl := request.QueryDSLRequest{
		TargetType: "file",
		TargetIds:  []string{"1124890692330272"},
		Query: request.BoolOp{
			Contain: map[string]any{"floor": "B01", "familyType": "标准"},
			Match:   map[string]any{"family": "family1"},
			BoolAnd: []request.BoolOp{
				{Match: map[string]any{"categoryId": "id111"}},
				{Match: map[string]any{"boundingBox.min.x": 167899.9999999998}},
			},
			BoolOr: []request.BoolOp{
				{Match: map[string]any{"productID": "KDKE-B-9947-#kL5"}},
				{
					BoolAnd: []request.BoolOp{
						{Match: map[string]any{"productID": "JODL-X-1937-#pV7"}},
						{Match: map[string]any{"price": 30}},
					},
				},
			},
		},
	}

	bts, _ := json.MarshalIndent(dsl, "", "\t")
	fmt.Println(string(bts))
}
