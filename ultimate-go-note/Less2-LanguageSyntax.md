
## 2.1 Built-in types

type cung cấp 2 thông tin:
- memory sử dụng là bao nhiêu (1, 2, 4, 8 bytes)
- kiểu dữ liệu đó thể hiện thông tin gì (int, uint, bool, ...)

type có thể chỉ rõ lượng memory ví dụ như int32 hoặc float64, hoặc không chỉ rõ (non-precision based type) như uint hay int. Khi đó size phụ thuộc vào architecture của máy:
- 32 bit arch: int = 4 bytes 
- 64 bit arch: int = 8 bytes 

## 2.2 word size

word size = size của int = phụ thuộc vào arch (32 bit or 64 bit)


## 2.3 Zero value

khi khai báo biến -> go compiler set tất cả các bit = 0 của kiểu dữ liệu đó. VD:
- bool = false
- int = 0
- string = "" (empty string)

## 2.4 Declare & Initialize
Dùng từ khóa var để khai báo biến -> được set giá trị zero value 
Với kiểu dữ liệu string là kiểu 2-word:
- word đầu = pointer đến array of bytes
- word thứ 2 = số lượng bytes của array đó
- zero value của string: 1st-word = nil, 2nd-word = 0

## 2.5 Conversion over Casting

Go không có casting từ kiểu dữ liệu này sang kiểu khác, mà là Conversion. Khi conversion, các bytes được copy sang vùng nhớ khác tương ứng với kiểu dữ liệu mới


## 2.6 Struct
khai báo anynomous struct (unnamed literal type)

e := struct {
	counter int
} {
	counter: 10
}


## 2.7 Padding & Alignment
```
type example struct {
	flag bool
	counter int16
	pi float32
}
```
tốn tổng cộng 8 bytes chứ ko phải 7 bytes (bool + int16 + float32) do được cộng 1 padding byte vào giữa flag và counter. đó là cơ chế alignment của compiler.

Để optimize bộ nhớ, cần sắp xếp các field với type từ nhiều byte đến ít byte nhất (float32 > int16 > bool). Tuy nhiên chỉ optimize khi thật sự cần thiết, phải ưu tiên sự chính xác, dễ đọc của code.


## 2.8 Assigning values

có 2 struct với cấu trúc giống hệt nhau, khai báo 2 biến e1 và e2 thì không thể assign ``` e1 = e2 ``` mà bắt buộc phải sử dụng conversion để compiler hiểu được.


## 2.9 Pointer

- Mỗi khi chạy ctrinh Go, Go runtime sinh ra các goroutines. Luôn có ít nhất 1 goroutine là main goroutine. mỗi goroutine đc cấp cho 1 vùng nhớ là stack, mỗi stack đc bắt đầu với 2048 bytes (2k) và có thể tăng lên theo tgian. 

- Khi có 1 function đc gọi, 1 vùng con của stack được gán cho Goroutine sử dụng, được gọi là frame. Size của frame được tính toán khi runtime, nếu như có value mà compiler ko biết trước size, nó sẽ được khởi tạo trong heap. 

- Pointer với mục đích để share value accross program boudaries. Giữa các frame ở trên là 1 dạng boundary.

## 2.10 Pass by value

- dữ liệu được truyền giữa các boundaries qua value của nó.
- Có 2 loại value là value của biến và address's value.

## 2.11 Escape Analysis
- là algorithm mà compiler dùng để xác định xem khởi tạo 1 biến trong stack hay trong heap
- Khởi tạo trong heap được gọi là allocation trong Go
- Nếu như 1 biến khởi tạo trong function vẫn tồn tại khi hết function -> heap, ngược lại là stack
- VD: tạo 1 biến và return giá trị biến đó thì là stack, còn nếu return address -> biến đó vẫn tồn tại -> heap

## 2.12 Stack Growth

- size của frame được tính toán trong khi compile, do đó nếu compiler ko biết chính xác size của biến -> sẽ được khởi tạo trong heap. 
- Khi hàm đc gọi, Go tính toán xem có đủ memory trong stack cho frame đó ko. Nếu ko đủ, 1 stack mới lớn hơn sẽ được tạo và các thông tin sẽ được copy sang.

## 2.13 Garbage Collection
- khi 1 value được khởi tạo trong heap, thì GC sẽ đc sử dụng. 


