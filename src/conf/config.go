package conf
var (
	DbConf = make(map[string]string)
)
func init(){
	DbConf["driver"] = "mysql"
	DbConf["user"] = "mrzhang"
	DbConf["password"] = "123456"
	DbConf["dbname"] = "mrzhang"
}

func GetConf()(map[string]string){
	return DbConf
}