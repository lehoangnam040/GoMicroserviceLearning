## 3.1 CPU Caches

- mỗi core trong processor có cache L1 và L2. Các core có 1 cache chung L3 và main memory
- Latency để access data từ ít (tốt nhất) đến nhiều nhất: L1 -> L2 -> L3 -> main
- performance dựa vào how efficiently data flows through hardware. data nằm trong L1 và L2 sẽ làm ctrinh chạy nhanh hơn việc data chỉ nằm trong main memory.
- Làm sao để data luôn được access tại L1, L2. Dựa vào prefetch, prefetch sẽ xác định memory để đưa vào cache. 
- Để làm đc, 1 cách đó là construct a contiguous block of memory and then iterate over that.
- Do đó, array là cấu trúc cực kỳ tốt bởi vì các element cách đều nhau, và khi iterate over array thì lượng memory đưa vào cache line được xác định trước -> dễ dàng đẩy vào L1, L2.

## 3.2 Translation Lookaside Buffer (TLB)
TLB là 1 small cache inside processor nhằm mapping virtual address to physical address. 


## 3.3 Declaring and Initializing Values of Array

khi khai báo 1 mảng string, mỗi phần tử string là 2-word với word đầu tiên là 1 pointer trỏ đến 1 array, word còn lại là độ dài của array đó.

## 3.4 String Assignments
khi string "Apple" được assign cho string x, bản chất là nó copy 2-word của "Apple" cho 2-word x, do đó sẽ có 2 2-word string cùng trỏ đến 1 backing array là "Apple"

## 3.5 Iterating Over Collections (Array)
có 2 cách để iterate over array
- Value semantic: for i, fruit := range fruits


- Pointer semantic: for i := range fruits

## 3.6 Value Semantic Iteration
- for i, fruit := range fruits
- Sử dụng value semantic thì sẽ cost hơn do nó sẽ copy toàn bộ array sang vùng nhớ khác và việc truy cập tại vùng nhớ mới (ko phải array fruits ban đầu)

- Biến fruit hay việc thao tác với biến fruit sẽ đều tạo ra các copy của 2-word string cùng trỏ đến 1 backing array, backing array đó được tạo trong heap.

## 3.7 Pointer Semantic Iteration
- dùng pointer sẽ được thao tác ở original collection

## 3.8 Data Semantic Guideline For Built-In Types
- sử dụng value semantic nếu làm việc với built-in types, để làm việc với copy tránh gây bug do động vào original

## 3.9 Different Type Arrays
[4]int khác với [5]int do size cũng hình thành nên type array, size của array phải được xác định ở compile time


## 3.10 Contiguous Memory Construction
khi in địa chỉ của các phần tử trong array sẽ thấy địa chỉ nằm liên tiếp cạnh nhau, cách nhau các khoảng tương ứng với type dữ liệu

## 3.11 Constructing Slices

- Slice is 3-word data structure. 1st-word is pointer to backing array, 2nd-word is length = number of elements can be accessed in backing array. 3rd-word is capacity = total number of element existed in backing array.

- Các cách khởi tạo
```
// Slice of string set to its zero value state.
var slice []string
// Slice of string set to its empty state.
slice := []string{}
// Slice of string set with a length and capacity of 5.
slice := make([]string, 5)
// Slice of string set with a length of 5 and capacity of 8.
slice := make([]string, 5, 8)
// Slice of string set with values with a length and capacity of 5.
slice := []string{"A", "B", "C", "D", "E"}
```

## 3.12 Slice Length vs Capacity
- đã giải thích bên trên


## 3.13 Data Semantic Guideline For Slices
- recommend sử dụng value semantic khi thao tác vs slide, tránh các bugs nếu dùng pointer semantic


## 3.14 Contiguous Memory Layout
- slice sử dụng array với memory liên tục -> có performance tốt 

## 3.15 Appending With Slices
- hàm append được sử dụng để thêm values vào slice có sẵn
- slice đưa vào hàm append là value semantic, nên nếu muốn thay đổi slice đang có thì cần gán lại return value của hàm append vào slice đang có. VD: ``` data = append(data, "ABC") ```

- khi gọi append, hàm sẽ check nếu length = capacity sẽ thực hiện copies backing array thành 1 backing array mới, và tăng capacity thêm 25% hoặc 100%. nếu length < capacity tức là vẫn còn chỗ trống thì chỉ cần thêm giá trị vào chỗ trống đó. Do đó vẫn giữ được contiguous memory của array -> làm cho hàm append hoạt động rất tốt.

## 3.16 Slicing Slices

- việc này tạo 1 slice mới là sub-array của slice cũ mà ko phải copies hay tốn thêm vùng nhớ, mà chỉ đơn giản tạo 1 3-word slice trỏ vào address của cùng 1 backing array.

```
slice1 := []string{"A", "B", "C", "D", "E"}
slice2 := slice1[2:4]
```

- ở VD trên, slice2 có 2 phần tử là C và D. tuy nhiên vẫn dùng chung 1 backing array, slice1 trỏ vào address của A còn slice2 trỏ vào address của C 


## 3.17 Mutations To The Backing Array
- vẫn ở VD mục 3.16, nếu thay đổi value ở slice2 thì slice1 cũng bị thay đổi vì trỏ trung 1 backing array ở heap

- nếu có mục đích ko dùng chung 1 backing array, khi slice thì cần tạo ra 1 backing array mới để slice2 trỏ đến.

- C1: có thể slice bằng cách slice2=slice1[a:b:c] trong đó len(slice2) = |a-b| và cap(slice2) = |a-c| và b=c -> len = cap, khi đó nếu như dùng hàm append thì sẽ tạo ra 1 backing array mới. 
- C2: Có thể dùng hàm copy, mục dưới

## 3.18 Copying Slices Manually
```
slice1 := []string{"A", "B", "C", "D", "E"}
slice3 := make([]string, len(slice1))
copy(slice3, slice1)
```
- sau khi copy, slice3 có backing array khác vs slice1


## 3.19 Slices Use Pointer Semantic Mutation

- sử dụng pointer trỏ đến 1 phần tử của slice backing array, sau đấy lại dùng hàm copy hoặc append làm backing array của slice bị thay đổi có thể dẫn đến các side effect ko mong muốn. Do vậy cần cẩn thận, double check khi code.


## 3.20 Linear Traversal Efficiency
khi slicing slice thì ko tạo ra vùng nhớ thêm -> minimize heap allocation

## 3.21 UTF-8

- Go’s compiler expects all code to be encoded in the UTF-8 character set

- UTF-8 is a character set , được thể hiện bởi code point và trong Golang là kiểu dữ liệu rune

- s := " 世界 means world" thì s có 18 bytes, trong đó mỗi chữ TQ là 3 bytes kiểu rune.

- khi dùng vòng for ``` for i, r := range s {  ``` thì r sẽ là kiểu rune chứ ko phải for từng bytes trong array. do số lượng bytes của rune khác nhau nên trong vòng for thì biến i sẽ ko tăng liên tục mà có thể nhảy cóc, VD: 0 -> 3 -> 6 -> 7 -> ... tùy theo UTF character
- rl := utf8.RuneLen(r) hàm này dùng để tính số bytes mà biến rune đó cost. VD với chữ TQ thì rl = 3

- với array var buf [4]byte khi dùng buf[:] thì sẽ biến thành 1 slice với cap = len = 4


## 3.22 Declaring And Constructing Maps

```
type user struct {
	name string
	username string
}
// Construct a map set to its zero value,
// that can store user values based on a key of type string.
// Trying to use this map will result in a runtime error (panic).
// vì zero value của map là nil nên khai báo như này và dùng thì sẽ lỗi
var users map[string]user

// Construct a map initialized using make,
// that can store user values based on a key of type string.
users := make(map[string]user)

// Construct a map initialized using empty literal construction,
// that can store user values based on a key of type string.
users := map[string]user{}
```

- dùng for loop để iterate over map sẽ random thứ tự key mỗi lần dùng for. 

## 3.23 Lookups and Deleting Map Keys

```
value, exist := map_variable[key]
```
- exist là 1 biến bool


- để delete cặp K-V dùng hàm ``` delete(users, "Roy") ```


## 3.24 Key Map Restrictions
- ko phải kiểu dữ liệu nào cũng có thể làm key trong map 
- VD: slice ko thể dùng làm key vì ko có cách để compare 2 slice. thường là kiểu dữ liệu có thể compare được thì mới đc dùng làm key



