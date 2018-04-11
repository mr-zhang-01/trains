/*
	地区功能
*/
package trains

/*
	地区结构
	@param id 地区id
	@param name 地区名称
	@param parentId 父级地区id
	@param childIds 子级地区id字符串
*/
type area struct{
	id int
	name string
	parentId int
	childIds string
}

/*
	地区集合
*/
type areas struct{
	data []area
}

func NewAreas()(*areas){
	return &areas{

	}
}