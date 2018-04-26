package trains

import (
	"testing"
	"fmt"
)

func Test_Sites(t *testing.T){
	s := NewSites()

	siteData := Site{
		AreaId: 1,
		Name: "测试",
	}

	fmt.Println(123)

	s.Add(siteData)

	t.Fatal("exit")
}