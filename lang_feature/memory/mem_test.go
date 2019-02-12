package memory

//禁用内联优化时，会执行runtime.newObject的汇编指令，但使用内联则不会
func test() *int {
	x := new(int)
	*x = 0xAABB
	return x
}
func main() {
	println(*test())
}
