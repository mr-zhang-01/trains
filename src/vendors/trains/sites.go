/*
	站点功能
*/
package trains

import (
	"dbsource"
)

var(
	TableName = "site"
)

func init()(){

}

/*
	站点数据结构
	@param id 站点id
	@param areaId 地区id
	@param name 站点名称
*/
type Site struct{
	ID int
	AreaId int
	Name string
}

/*
	站点集合数据结构
*/
type Sites struct{
	Data []Site
}

func NewSites()(*Sites){
	return &Sites{
	}
}

/*
	添加站点
	@param d 站点数据结构
*/

func(s *Sites)Add(d Site)(bool){
	b := map[string]interface{}{
		"area_id": d.AreaId,
		"name": d.Name,
	}
	dbsource.MySqlHanle.Add(TableName, b)
	return false
}

/*
	查找站点
	@param search 含有areaid或者name等信息
*/
func(s *Sites)Get(search map[int]interface{})([]Site){
	return nil
}
