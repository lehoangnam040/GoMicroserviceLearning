## 5.1 Grouping Different Types of Data

- trong Golang ko có khái niệm sub-typing hoặc sub-classing

- do vậy, ko tạo class dựa trên đặc điểm chung để tái sử dụng, mà tạo interface dựa trên behavior chung để decoupling code

## 5.2 Don’t Design With Interfaces

- hãy tạo các struct, type trước và code hoạt động đúng, sau đó hãy nghĩ đến việc tạo interface


# 5.3 -> 5.6 
- 3 chap này nói về việc áp dụng interface vào decoupling 

## 5.3 Composition
- Việc composition  bắt đầu sau khi đã định nghĩa các type, struct. 
- các struct type đó là primitive layer đưa ra các behavior để giải quyết bài toán
- layer tiếp là các function nhận param các các struct của primitive layer, nhằm thực hiện 1 logic nào đó
- layer cuối cùng là tạo ra các api, hàm để sử dụng build trên function layer


## 5.4 Decoupling With Interfaces
- lúc này, nếu muốn mở rộng code, dựa trên các behavior đã định nghĩa ở primitive layer để tạo ra các interface tương ứng với behavior đó

## 5.5 Interface Composition
- tiếp tục decoupling bằng interface cho các layer trên

## 5.6 Precision Review
- review lại sau khi đã sử dụng interface để có thể remove các đoạn code thừa, ko cần thiết 

## 5.7 Implicit Interface Conversions
- có thể convertions 1 data từ interfaceA sang interfaceB nếu thỏa mãn đk sau

```
type A interface {
	Func1()
}
type C interface {
	Func2()
}
type B interface {
	A
	C
}
```

khi đó 1 data implement cả 2 behavior Func1 và Func2 được coi như là interface B và có thể convert sang A nếu cần. Tuy nhiên nếu chỉ có Func1 thì tức là interface A và ko thể convert sang B

```
// 2 values sau là valueless vì là zero value của interface
var dataB B
var dataA A

dataB = x{}	// giả sử struct x có implement 2 Func1 và Func2
dataA = dataB	// conversion 

// Lỗi
dataB = dataA 	// ko thể convert như này
```


## 5.8 Type assertions
xác định xem 1 data có phải có behavior thuộc 1 interface nào đó ko

```
y := x.(B)		// xác định xem x có phải interface B hay ko
dataB = y		// nếu đúng thì có thể gán, ko thì báo lỗi


y, ok := x.(B)		// assert trả về ok là bool, tránh panic khi chạy
```


## 5.9 Interface Pollution
- trong golang, ko phải lúc nào cũng tạo interface
- luôn bắt đầu với tạo data trước, nếu bài toán cần mới decoupling sử dụng interface
- interface luôn đi với behavior (động từ)


## 5.10 Interface Ownership

- do việc tạo interface theo behavior. nên có thể developer khi code 1 package thì ko cần tạo interface. ng sử dụng package đó muốn mock hoặc 1 việc gì đó khác cần tạo interface thì tự tạo ở project đó, ko cần yêu cầu quay lại project gốc để sửa.



