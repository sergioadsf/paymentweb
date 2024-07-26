package configuration

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/lib/pq"
)

type DatabaseConf struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Database string `json:"dbname"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (d DatabaseConf) toParams() *pq.Connector {
	fmt.Println(d.Database)
	fmt.Println(d.Name)
	fmt.Println(d.Password)
	fmt.Println(d.Username)
	con, err := pq.NewConnector(fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", d.Username, d.Password, d.Database))
	if err != nil {
		panic("Não foi possível conectar ao banco!")
	}

	return con
}

const startToken = "${"
const endToken = "}"

func (d DatabaseConf) Parse() DatabaseConf {

	valueOf := reflect.ValueOf(&d).Elem()

	for i := 0; i < valueOf.NumField(); i++ {
		ptr := valueOf.Field(i)
		if ptr.Kind() == reflect.Ptr {
			ptr = ptr.Addr()
		}

		str, ok := ptr.Interface().(string)
		if !ok {
			continue
		}

		hasEnv := strings.Index(str, startToken)
		if hasEnv < 0 {
			continue
		}

		endEnv := strings.Index(str[hasEnv:], endToken)
		if endEnv < 0 {
			continue
		}

		sub := str[hasEnv+len(startToken) : endEnv]
		val, ok := os.LookupEnv(sub)

		if ok && ptr.CanSet() {
			ptr.SetString(val)
		}
	}
	return valueOf.Interface().(DatabaseConf)
}
