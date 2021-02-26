package response

/*
   "createTime" : "2018-03-29 18:26:01",
   "databagVersion" : "3.1",
   "length" : 0,
   "reason" : "reason",
   "status" : "success"
*/

//Bake ***
type Bake struct {
	Status         string `json:"status"`
	Reason         string `json:"reason"`
	Length         int64  `json:"length"`
	DatabagVersion string `json:"databagVersion"`
	CreateTime     string `json:"createTime"`
}
