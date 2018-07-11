package Helper

import (
	"fmt"
	"strings"
)

const GROUP_NODE_DELIMITER  = "."

func GetContextSqlKey() string {
	return "swoft-sql"
}

func GetPool()  {
	
}
func GetStatementClassNameByInstance(instance string)string  {

}

func GetContextTransactionsKey() string  {
	return fmt.Sprintf("transactions")
}


func GetTsInstanceKey(instance string) string {
	return instance
}


func GetDriverByInstance(instance string) string {

}

func GetPoolName(group string, node string) string  {
	groupNode := strings.Split(group, GROUP_NODE_DELIMITER)
	if len(groupNode) == 2{
		return group
	}
	return fmt.Sprintf("%s.%s", group, node)
}