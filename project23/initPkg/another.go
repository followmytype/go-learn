package initpkg

var Name1 = "another"

func init()  {
	println("Name1 is", Name1)
	Name1 = "another here"
	println("Name1 is", Name1)
}
