package trains

import (
	"testing"
	"fmt"
)

func Test_site(t *testing.T){
	siteData := Site{
		AreaId: 1,
		Name: "测试",
	}

	fmt.Println(siteData)

	//s := NewSites()
	//s.Add(siteData)

	t.Errorf("test")
}