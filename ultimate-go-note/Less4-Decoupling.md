## 4.1 Methods

- method là function với việc declare "receiver". Receiver là parameter nằm giữa keyword "func" và func name. 
- Receiver có 2 loại: value và pointer
- method cho phép data có khả năng behavior

- value receiver làm việc trên copy of data
- pointer receiver làm việc trên shared access to data
- các method của 1 type ko nên mix giữa value và pointer, gây khó đọc và maintain code. Ngoại trừ 1 vài TH đặc biệt và cần phải dùng


## 4.2 Method Calls

khi call method, compiler cần xác định xem đúng kiểu dữ liệu, còn nó sẽ tự điều chỉnh pointer hoặc value


## 4.3 Data Semantic Guideline For Internal Types

- nếu data đang làm việc là internal type (slice, map, channel, function, interface) thì dùng value


- 1 trường hợp sử dụng pointer ví dụ như cần share slice hoặc map với function mà làm nhiệm vụ unmarshalling hoặc decoding.


## 4.4 Data Semantic Guideline For Struct Types

- khi làm việc với struct, cần xác định dùng 1 trong 2 loại value or pointer.
- nếu safe to be copied -> dùng value

## 4.5 Methods Are Just Functions

- recommend ko nên implement setter và getter trong Golang
- sử dụng variable để gọi chứ ko nên dùng \*variable để gọi hàm 

## 4.6 Know The Behavior of the Code
- ví dụ


## 4.7 Interfaces

- interface giúp cho các data có cách thể hiện 1 behavior theo cách khác nhau với từng kiểu dữ liệu 

- sử dụng khi 1 hàm, api có nhiều cách implement dựa vào các kiểu dữ liệu khác nhau

- 1 phần logic trong hàm, api  có thể thay đổi và yêu cầu decoupling 

## 4.8 Interfaces Are Valueless

- interface chỉ định nghĩa behavior chứ ko có value, ko có vùng nhớ

- do vậy ko khai báo variable với kiểu dữ liệu là interface (khác so với OOP)


## 4.9 Implementing Interfaces

- Go is a language that is about convention over configuration.

- khi sử dụng interface trong Golang thì ko cần dùng từ khóa implement, hay chỉ rõ rằng type A có interface là I hay gì đó tương tự. Mà chỉ cần A có method giống signature với method của I là đủ


## 4.10 Polymorphism

- do vậy trong các func khác, khi nhận tham số I thì khi sử dụng có thể truyền vào data với kiểu dữ liệu A, bởi vì khi đó compiler hiểu rằng A có behavior của I


## 4.11 Method Set Rules
```
type notifier interface {
	notify()
}
type user struct {
	name string
	email string
}
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}
func sendNotification(n notifier) {
	n.notify()
}
func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(u)
}
```

- nếu dùng pointer receiver trong hàm implemented từ interface thì khi sử dụng bắt buộc phải dùng pointer
- vì compiler phải chắc chắn rằng type tương ứng với interface đó có address
- nên nếu đưa value vào sẽ bị lỗi
- VD: như đoạn code trên thì hàm sendNotification gọi hàm notify của interface notifier, khi sử dụng type user thì hàm đó lại là pointer receiver nên nếu đưa value vào sẽ bị lỗi ở main, mà phải dùng pointer


## 4.12 Slice of Interface

- nếu sử dụng value receiver:
	- nếu đưa value vào thì sẽ dùng copy of data
	- nếu đưa pointer vào sẽ dùng chung với original data


## 4.13 Embedding
- khai báo embedding
```
type user struct {
	name string
	email string
}
type admin struct {
	user			// Value Semantic Embedding
	level string
}
```

```
type user struct {
	name string
	email string
}
type admin struct {
	*user			// Pointer Semantic Embedding
	level string
}
```

- khi khai báo embedding, admin type có thể sử dụng các hàm, cũng như hàm interface của user type


## 4.14 Exporting

- để export thì hãy viết hoa chữ cái đầu của hàm, struct type
- khi export thì bên ngoài package sẽ có thể access được
