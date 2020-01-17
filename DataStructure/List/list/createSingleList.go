/*
@author : Aeishen
@data :  19/07/15, 14:19

@description : 创建单链表
*/

package List

type Object interface{}

type Node struct {
	Data Object   //数据域
	Next *Node    //指针域，指向下一个节点
}

type MyList struct {
	size int // 链表长度
	head *Node  // 头指针
	tail *Node  // 尾指针
}

// 初始化链表
func InitList() *MyList{
	var l MyList
	l.size = 0     // 此时链表是空的
	l.head = nil   // 没有头指针
	l.tail = nil   // 没有尾指针
	return &l
}

//创建一个链表
func CreateList(arr []int) *MyList{
	l := InitList()
	for _,v:= range arr{
		l.Add(v)
	}
	return l
}


// 添加元素
func (this *MyList)Add(data Object) bool{
	if data == nil {
		return false
	}

	newNode := &Node{data,nil}

	if this.size == 0{
		this.head = newNode
	}else{
		oldTail := this.tail
		oldTail.Next = newNode
	}

	// 调整尾部位置，及链表元素数量
	this.size ++
	this.tail = newNode
	return true
}

// 在index位置后插入元素
func (this *MyList)Insert(index int,data Object) bool{
	index = index - 1
	if index < 0 || data == nil || index >= this.size {
		return false
	}

	newNode := &Node{data,nil}
	if index == 0 {
		newNode.Next = this.head
		this.head = newNode
	}else{
		preItem := this.head           // 找到前一个元素
		for i := 1;i <= index ;i++{    // i <= index 为了找到index位置上的节点
			preItem  = preItem.Next
		}
		newNode.Next = preItem.Next
		preItem.Next = newNode
		if index == this.size - 1{     // 若插入在尾部，尾部指针需要调整
			this.tail = newNode
		}
	}

	this.size ++
	return true
}


// 删除index位置的节点
func (this *MyList)Delete(index int) bool{
	index = index - 1
    if index < 0 || index >= this.size{
    	return false
	}

    if index == 0{
    	newHead := this.head.Next
		this.head = newHead
		if this.size == 1{     // 如果只有一个元素，那尾部也要调整
			this.tail = nil
		}
	}else{
		preItem := this.head
        for i:= 1;i < index ;i++{      // i < index 为了找到index位置上的节点的前一个节点
			preItem = preItem.Next
		}

        if index == (this.size - 1){   // 若删除的尾部，尾部指针需要调整
			preItem.Next = nil
			this.tail = preItem
		}else{
			preItem.Next = preItem.Next.Next
		}
	}

	this.size --
	return true
}

// 获取index位置的元素
func (this *MyList)Get(index int) (bool,Object){
	index = index - 1
	if index < 0 || index >= this.size {
		return false,nil
	}
	preItem := this.head
	for i:= 1;i <= index ;i++{   // i <= index 为了找到index位置上的节点
		preItem = preItem.Next
	}
	return true,preItem.Data
}

// 设置index位置的元素的值
func (this *MyList)Set(index int, data Object) bool{
	index = index - 1
	if index < 0 || index >= this.size || data == nil{
		return false
	}
	preItem := this.head
	for i:= 1;i <= index ;i++{   // i <= index 为了找到index位置上的节点
		preItem = preItem.Next
	}
	preItem.Data = data
	return true
}

// 搜索元素
func (this *MyList)Search(data Object) (isFound bool,indexs []int) {
	indexs = []int{}
	isFound = false
	if this.size == 0{
		return
	}

	preItem := this.head
	for i:= 0; i<this.size; i++{
		if preItem.Data == data {
			indexs = append(indexs,i + 1)
		}
		preItem = preItem.Next
	}
	if len(indexs) > 0{
		isFound = true
	}
	return
}

//遍历链表
func (this *MyList)Traverse() (all []Object){
	if this.size == 0{
		return nil
	}
	preItem := this.head
	for i:= 0;i < this.size;i++{
		all = append(all,preItem.Data)
		preItem = preItem.Next
	}
	return
}

//链表逆序
func (this *MyList)Reversal() {
	if this.size == 0 || this.size == 1 {
		return
	}

	if this.size == 2{
		tempNode := this.tail
		this.tail = this.head
		this.head = tempNode
		this.head.Next = this.tail
	}else{
		preItem := this.head
		curItem := preItem.Next
		for i:= 1;i < this.size ;i++{
			nextItem := curItem.Next
			curItem.Next = preItem
			preItem = curItem
			curItem = nextItem
		}
		tempNode := this.tail
		this.tail = this.head
		this.head = tempNode
	}
}

// 获取链表长度
func (this *MyList)GetSize() int{
	return this.size
}

// 获取链表头指针
func (this *MyList)GetHead() (bool, *Node){
	if this.size <= 0{
		return false,&Node{nil,nil}
	}
	return true ,this.head
}

// 获取链表尾指针
func (this *MyList)GetTail() (bool, *Node){
	if this.size <= 0{
		return false,&Node{nil,nil}
	}
	return true,this.tail
}

// 置空链表
func (this *MyList)RemoveAll() {
	this.size = 0     // 此时链表是空的
	this.head = nil   // 没有头指针
	this.tail = nil   // 没有尾指针
}

