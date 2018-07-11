package Helper

func FormatListByType(result []map[string]string, className string) []map[string]string  {
	if result == nil{
		return nil
	}
	var rowList  []map[string]string
	rowList =    []map[string]string{}
	for i,row := range result{
		rowList[i] = FormatRowByType(row, className)
	}
	return  rowList
}
func FormatRowByType(row map[string]string, className string) map[string]string  {
	return  row
}
func ListToEntity()  {
	
}
func ArrayToEntity()  {
	
}

func TrasferTypes()  {
	
}
func TransferParameter()  {
	
}
func FormatParamsKey()  {
	
}
