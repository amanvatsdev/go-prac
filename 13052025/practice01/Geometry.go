package practice

func Square (a int)int{
	return a*a
}
func Calculate(unit int ,shape func (int)int)int{
	return  shape(unit)
}