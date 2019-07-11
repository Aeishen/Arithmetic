package Stack

// 栈元素
type Item interface {}

// 栈
type ItemStack struct {
	Items []Item
}

// 创建栈
func (this *ItemStack)new() *ItemStack{
	this.Items = []Item{}
	return this
}

// 入栈
func (this *ItemStack)Push(i Item)  {
	this.Items = append(this.Items,i)
}

// 出栈
func (this *ItemStack)Pop() {
	this.Items = this.Items[:len(this.Items) - 1]
}

// 获取栈顶元素
func (this *ItemStack)Peek() Item{
	i := this.Items[len(this.Items) - 1]
	return i
}

// 获取栈中元素个数
func (this *ItemStack)GetSize() int{
	return len(this.Items)
}

// 栈判空
func (this *ItemStack)IsEmpty() bool{
	return len(this.Items) <= 0
}

// 栈的初始化
func InitStack() *ItemStack{
	var this ItemStack
	if this.Items == nil{
		this = ItemStack{}   //创建一个空结构体
		this.new()
	}
	return &this
}
