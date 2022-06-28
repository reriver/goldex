package goldex

var localGoldexVar int

func Return42(i int) int {
	return 42
}
func Add42(i int) int {
	return i + 42
}

func GetLocalVar() int {
	return localGoldexVar
}

func SetLocalVar(i int) {
	localGoldexVar = i
}
