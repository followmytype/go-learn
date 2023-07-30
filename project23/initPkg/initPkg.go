package initpkg

var Name = "initPkg"

func init()  {
	println("Name is", Name)
	Name = "initPkg here"
	println("Name is", Name)
}
