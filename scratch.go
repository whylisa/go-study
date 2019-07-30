package main

import "fmt"

//ch < - v
//v := <-ch
//func say (s string ) {
//	for  i:= 0; i < 5;i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}


//}
//func main () {
//	go say("world")
//	say("hello")
//}
//import "fmt"
//func main() {
//	var a string = "roob"
//	print(a)
//	var b,c int =  1,2
//	var s int
//	var e bool
//	print(b,c,s,e)
//}
//
//func max(num1,num2 int) int {
//	var result int
//	if(num1 > num2) {
//		result = num1
//	}else {
//		result = num2
//	}
//	return result
//}
//func main() {
//	var a int = 20
//	var ip *int
//	ip = &a
//	print("a变量的地址是：%x\n",&a)
//	print("a变量的地址是：%x\n",ip)
//	print("a变量的地址是：%x\n",*ip)
//}
//func main() {
//	var ptr *int
//	print("ptr的值为：%x\n",ptr)
////}
//type Books struct {
//	title string
//	author string
//	subject string
//	book_id int
//}

//func main() {
//	var Book1 Books        /* 声明 Book1 为 Books 类型 */
//	var Book2 Books        /* 声明 Book2 为 Books 类型 */
//
//	/* book 1 描述 */
//	Book1.title = "Go 语言"
//	Book1.author = "www.runoob.com"
//	Book1.subject = "Go 语言教程"
//	Book1.book_id = 6495407
//
//	/* book 2 描述 */
//	Book2.title = "Python 教程"
//	Book2.author = "www.runoob.com"
//	Book2.subject = "Python 语言教程"
//	Book2.book_id = 6495700
//
//	/* 打印 Book1 信息 */
//	printBook(Book1)
//
//	/* 打印 Book2 信息 */
//	printBook(Book2)
//}
//
//func printBook( book Books ) {
//	fmt.Printf( "Book title : %s\n", book.title);
//	fmt.Printf( "Book author : %s\n", book.author);
//	fmt.Printf( "Book subject : %s\n", book.subject);
//	fmt.Printf( "Book book_id : %d\n", book.book_id);
//}
//type Books struct {
//	title string
//	anthor string
//	subject string
//	book_id int
//}

//func main() {
//	var book1 Books
//	var book2 Books
//
//	book1.title = "go语言"
//	book1.anthor = "www"
//	book1.subject = "教程"
//	book1.book_id = 3
//
//	book2.title = "go语言"
//	book2.anthor = "www"
//	book2.subject = "教程"
//	book2.book_id = 3
//
//	printBook(&book1)
//	printBook(&book2)
//
//}
//
//func printBook( book *Books) {
//	fmt.Println(book.title)
//}
//var slice1 []types = make([]type,len)
//slice1 := make([]type,len)

//s := [] int {1,2,3}

//func main() {
//	var numbers = make([]int, 3,5)
//	printSlice(numbers)
//}
//func printSlice(x []int) {
//	fmt.Println("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}
//func main() {
//	numbers := []int{0,1,2,3,4,5,6,7,8}
//	printSlice(numbers)
//	fmt.Println(numbers[1:4])
//	fmt.Println(numbers[:3])
//	fmt.Println(numbers[4:])
//
//}


//func main() {
//	 var numbers []int
//	 printSlice(numbers)
//	 numbers = append(numbers,1)
//	 numbers = append(numbers,1,2,3,4,5)
//	 numbers1 := make([]int, len(numbers),(cap(numbers))*2)
//	 copy(numbers1,numbers)
//
//}
//func printSlice(x []int) {
//	fmt.Println("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}

//func main() {
//	nums := []int{2,3,4}
//	sum := 0
//	for _,num := range nums {
//		sum += num
//	}
//	fmt.Println("sum:",sum)
//	for i,num := range nums {
//		if num == 3 {
//			fmt.Println("index",i)
//		}
//	}
//	kvs := map[string]string{"a": "apple","b": "banana"}
//	for k,v := range kvs {
//		fmt.Println(k,v)
//	}
//}

//func main() {
//	var a map[string]string
//	a = make(map[string]string)
//	a["fr"] = "dfs"
//	delete(a,"fr")
//
//	capital,ok := a["fr"]
//	if(ok) {
//		fmt.Println(capital)
//	}else {
//		fmt.Println("buc")
//	}
//}
//
//func a () {
//	a()
//}
//func main () {
//	a()
//}
//func a( n uint64)(result uint64) {
//	if(n > 0) {
//		result = n * a(n-1)
//		return result
//	}
//	return 1
//}
//func main () {
//	var i int  = 15
//	fmt.Println(a(uint64(i)))
//}

//ch := make(chan int )
//func sum(s []int,c chan int ) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum
//}
//
//func main() {
//	s := []int{7,34,234,432,42,4}
//	c :=  make(chan int)
//	go sum(s[:len(s)/2],c)
//	go sum(s[len(s)/2:],c)
//	x,y := <-c,<-c
//
//	fmt.Println(x,y,x+ y)
//}
//type Retriever struct {
//	UserAgent string
//	TimeOut time.Duration
//}
//
//func (r *Retriever) get(url string) string {
//	resp,err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	result, err := httputil.DumpResponse(
//		resp,true)
//	resp.Body.Close()
//	if err != nil {
//		panic(err)
//	}
//	return string (result)
//}
//func gerneration() chan int {
//	out := make(chan int)
//	go func() {
//		i := 0
//		for {
//			time.Sleep(
//				time.Duration(rand.Intn(1000)* time.Millisecond))
//		}
//	}
//}

//type Queue []int
//func (q *Queue) Push(v int) {
//	*q = append(*q,v)
//}
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head
//}
//func (q *Queue) IsEmpty()bool {
//	return len(*q) == 0
//}

//func 函数名（形式参数列表）（返回值列表） {
//	函数体
//}
//func hypot(x,y,float64) float64 {
//	return math.Sqrt(x*x+y*y)
//}
//func f(i,j,k,int,s,t,string){}
//func f(i int ,j string,k int,s string ,t string){}

//const (
//	SecondsPerMinute = 60
//	SecondsPerHour = SecondsPerMinute * 60
//	secondsPerDay = SecondsPerHour * 24
//)
//func resolveTime(seconds int) (day int, hour int ,minute int,) {
//	day = seconds / secondsPerDay
//	hour = seconds / SecondsPerHour
//	minute = seconds / SecondsPerMinute
//
//	return
//}
//
//func main() {
//	_,hour,minute := resolveTime(18000)
//	fmt.Println(hour,minute)
//	day,_,_ := resolveTime(90000)
//	fmt.Println(day)
//}
//
//func fire() {
//	fmt.Println("fire")
//}
//func main () {
//	var f func()
//	f = fire
//
//	f()
//}
//func StringProcess(list []string,chain []func(string) string) {
//	for index,str := range list {
//		result := str
//		for _,proc := range chain {
//			result = proc((result))
//		}
//		list[index] = result
//	}
//}
//
//func removePrefix(str string) string {
//	return string.TrimPrefix(str,"go")
//}
//func main () {
//	list := []string {
//		"go scanner",
//		"go parser"
//	}
//	chain := []func(string) string {
//		removePrefix,
//		strings.TrimSpace,
//		strings..ToUpper
//	}
//	StringProcess(list,chain)
//	for _, str := range list {
//		fmt.Println(str)
//	}
//}
//
//func StringProcess(list []string, chain []func(string) string) {
//	for index,str := range list {
//		result := str
//		for _,proc := range chain {
//			result = proc(result)
//		}
//		list[index] = result
//	}
//}
//func removePrefix(str string) string {
//	return string.TrimPrefix(str,"go")
//}
//func (data int) {
//	fmt.Println("hello",data)
//}(100)
//f := func(data int ) {
//	fmt.Println("hello", data)
//}
//f(100)
//func visit(list []int, f func(int) ) {
//	for _,v := range list {
//		f(v)
//	}
//}
//func main() {
//	visit([]int{1,2,3,4}, func(v int) {
//		fmt.Println(v)
//	})
//}

//var skillParam = flag.String("skill","","skil to perform")
//
//func main() {
//	flag.Parse()
//	var skill = map[string]func() {
//		"fire": func() {
//			fmt.Println("chicken fire")
//	},
//		"run": func() {
//			fmt.Println("sloider run")
//		},
//		"fly": func() {
//			fmt.Println("skill not found")
//		},
//	}
//	if f , ok := skill[*skillParam];ok {
//		f()
//	}else {
//		fmt.Println("skill not found")
//	}
//}
//type Invoke interface {
//	Call(interface{})
//}
//
//type Struct struct {
//
//}
//func (s *Struct)Call(p interface{}){
//	fmt.Println("from struct", p)
//
//}
//type FuncCaller func(interface{})
//
//func (f FuncCaller)Call(p interface{}) {
//	f(p)
//}
//func main() {
//	var invoke Invoke
//	s := new(Struct)
//	invoke = s
//	invoke.Call("hello")
//
//	invoke = FuncCaller((func(v interface{}) {
//		fmt.Println("from function ",v)
//	}))
//	invoke.Call("hello")
//}
//
//type Invoke interface {
//	Call(interface{})
//}
//type Struct struct {
//
//}
//func (s *Struct)Call (p interface{}) {
//	fmt.Println("from struct",p)
//}
//var invoker Invoker
//s:=new(Struct)
//invaker = s
//invoker.Call("hello")
//
//type FuncCaller func(interface{})
//
//func (f FuncCaller)Call(p interface{})  {
//	f(p)
//}
//
//var invoker Invoker
//invoker = funcCaller(func(v ast.InterfaceType) {
//	fmt.Println("from function", v)
//})
//invoker.Call("hello")

//type HandlerFunc func(w http.ResponseWriter,*Request)
//func (f HandlerFunc) ServeHTTP(w http.ResponseWriter,r *Request) {
//	f(w,r)
//}
//type Point struct {
//	X int
//	Y int
//}
//type Color struct {
//	R,G,B byte
//}
//type Point struct {
//	 X int
//	 Y int
//}

//var p Point
//p.X = 10
//p.Y = 20
//type People struct {
//	name string
//	child * People
//}
//
//relation := &People {
//	name: "也饿",
//	child: &People {
//		name: "爸爸",
//		child: &People {
//			name: "我"
//}
//}
//}type
//
//type Cat struct {
//	Color string
//
//}
//type DataWriter interface {
//	WriteData(data interface{}) error
//}
//type file struct {
//
//}
//func (d *file) WriteData(data interface{}) error {
//	fmt.Println("WriteData",data)
//	return nil
//}
//
//func main() {
//	f := new(file)
//	var writer DataWriter
//	writer = f
//	writer.WriteData("data")
//}


//cha1 := make(chan int)
//cha2 := make(chan interface{})
//type Equip struct{}
//ch2 := make(chan *Equip)
//ch := make(chan interface{})
//ch <- 0
//ch <- "hello"
//
//func main () {
//	ch := make(chan int)
//	ch <- 0
//}

//func main () {
//	ch := make(chan int )
//	go func () {
//		fmt.Println("start goroutine")
//		ch <- 0
//		fmt.Println("exit gorountine")
//	}()
//	fmt.Println("wait goroutine")
//	<-ch
//	fmt.Println("all done")
//}
//
//for data := range ch {
//
//}
//
//func main () {
//	ch := make(chan int )
//	go func () {
//		for i := 3; i >=0 ;i--{
//			ch <- i
//			time.Sleep(time.Second)
//		}
//	}()
//	for data := range ch {
//		fmt.Println(data)
//		if data == 0 {
//			break
//		}
//	}
//}
//

//type Retriever struct {
//	Contents string
//}
//func (r *Retriever) String() string {
//	return fmt.Sprintf(
//		"Retriver: {Contents=%s}",r.Contents
//		)
//}
//type Student struct {
//	name string
//	age int
//	Class string
//}
//
//func main () {
//	var stu1 Student
//	stu1.age = 22
//	stu1.name = "wd"
//	stu1.Class = "class1"
//	fmt.Println(stu1.name)
//
//	var stu2 *Student = new(Student)
//	stu2.name = "jack"
//	stu2.age = 32
//	fmt.Println(stu2.name,(*stu2).name)
//
//	var stu3 *Student = &Student{name:"rose",age:18,Class:"class3"}
//	fmt.Println(stu3.name,(*stu3).name)
//}

//type Student struct {
//	name string
//	age int
//	Class string
//}
//
//func Newstu(name1 string ,age1 int ,class1 string ) *Student {
//	return &Student{name:name1,age:age1,Class:class1}
//}
//func main () {
//	stu1  := Newstu("wd",22,"math")
//	fmt.Println(stu1.name)
//}
//
//type Student struct {
//	Name string "the name of student"
//	Age int "the age of student"
//	Class string "the class of student"
//
//}
//type Student struct {
//	Name string `json:"name"`
//	Age int `json:"age"`
//
//}
//
//func main () {
//	var stu = Student{Name:"wd",Age:22}
//	data,err := json.Marshal(stu)
//	if err != nil {
//		fmt.Println("json encode failed err:",err)
//		return
//	}
//	fmt.Println(string(data))
//}
// type Person struct {
// 	Name string
// 	Age int
// }
//
//type Teacher struct {
//	Salary int
//	Classes string
//}
//type man struct {
//	sex string
//	job Teacher
//	Person
//
//}
//
//func main () {
//	var man1 = new(man)
//	man1.Age = 22
//	man1.Name = "fd"
//	man1.job.Salary = 5555
//	fmt.Println(man1,man1.job.Salary)
//}

type Person struct {
	Name string
	Age int
}

func (p Person) Getname() string{
	fmt.Println(p.Name)
	return p.Name
}
func main () {
	var person1 = new(Person)
	person1.Age = 22
	person1.Name = "ss"
	person1.Getname()
}






