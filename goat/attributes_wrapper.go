package goat

type AttributesWrapper struct {
	actual *Attributes
	changes map[string]string
}

func (attr *AttributesWrapper) Init(at *Attributes){
	attr.actual = at
	attr.changes = nil
}

func (attr *AttributesWrapper) Get(x string) (string, bool){
	var out string 
	has := false
	var val string
	if attr.changes != nil{
		if val, has = attr.changes[x]; has {
			out = val
			return out, has
		}
	}
	return attr.actual.Get(x)
}

func (attr *AttributesWrapper) GetValue(x string) string{
    val, _ := attr.Get(x)
    return val
}

func (attr *AttributesWrapper) Has(x string) bool{
    _, has := attr.Get(x)
    return has
}

func (attr *AttributesWrapper) Set(key string, val string){
	if attr.changes == nil{
		attr.changes = map[string]string{key: val}
	} else {
		attr.changes[key] = val
	}
}

func (attr *AttributesWrapper) Commit() bool{
	if attr.actual == nil{
		panic("invalid attributes pointer!")
	} 
	if attr.changes != nil {
		anyChange := len(attr.changes) > 0
		for k, v := range attr.changes{
			attr.actual.Set(k, v)
		}
		attr.changes = nil
		return anyChange
	} else {
		return false
	}
}

func (attr *AttributesWrapper) Rollback(){
	attr.changes = nil
}
