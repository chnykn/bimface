package response

import "github.com/chnykn/bimface/v2/bean/common"

type ElementNodeBean common.TreeNode

type ElementNodeTree struct {
	Root  string             `json:"root"`
	Items []*ElementNodeBean `json:"items"`
}
