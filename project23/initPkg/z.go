package initpkg

var Namez = "zzzzz"

func init()  {
	println("Namez is", Namez)
	Namez = "zzzzz here"
	println("Namez is", Namez)
}
