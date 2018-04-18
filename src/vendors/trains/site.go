/*
	站点功能
*/
package trains

import (
	"dbsource"
)

var(

)

func init()(){

}

/*
	站点数据结构
	@param id 站点id
	@param areaId 地区id
	@param name 站点名称
*/
type site struct{
	id int
	areaId int
	name string
}

/*
	站点集合数据结构
*/
type sites struct{
	data []site
}

func NewSites()(*sites){
	return &sites{
	}
}

/*
	添加站点
	@param d 站点数据结构
*/

func(s *sites)Add(d site)(bool){
	s.data = append(s.data, d)

	dbsource.MySqlHanle.Add()
}

/*
	查找站点
	@param search 含有areaid或者name等信息
*/
func(s *sites)Get(search map[int]interface{})([]site){

}