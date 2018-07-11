package Db

type DbResult struct {
	Type int
	ClassName string
	Decorators []string

}

func (d * DbResult)SetType(Drtype int)  {
	d.Type = Drtype
}
func (d *DbResult)SetClassName(className string)  {
	d.ClassName = className
}
func (d *DbResult)SetDecorators(decorators []string)  {
	d.Decorators = decorators
}

func (d *DbResult)GetResultByClassName()  {
	className := d.ClassName
	result := getResultByType()
}
func (d *DbResult)getResultByType()[]string  {

}