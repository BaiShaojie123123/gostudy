package pointer

//*int代表指针
func swap(pa *int, pb *int) {
	var tmp int
	tmp = *pa
	*pa = *pb
	*pb = tmp
}

//交换
func main() {
	var a int = 1
	var b int = 2
	//&a代表传递的是a的地址
	swap(&a, &b)

}
