- Developer cần có trách nhiệm handler đầy đủ, chính xác error để trả về cho user 
- Handling error ở 3 thứ: logging the error, not propagating the error any further, and
determining if the Goroutine/program needs to be terminated

- Trong golang, ``` error``` là 1 builtin interface, cho nên nó ko cần viết hoa 
## Khai báo error
 Khai báo error thường dùng unexported struct và unexported field, sử dụng pointer receiver trong hàm Error -> chỉ có address của struct đó được share 

```
// http://golang.org/src/pkg/errors/errors.go
type errorString struct {
	s string
}
// http://golang.org/src/pkg/errors/errors.go
func (e *errorString) Error() string {
	return e.s
}

// http://golang.org/src/pkg/errors/errors.go
func New(text string) error {
	return &errorString{text}
}
```


## Check error
- Khi check error trong code, có thể khai báo các error khác nhau và dung switch case để check

```
var (
	ErrBadRequest = errors.New("Bad Request")
	ErrPageMoved = errors.New("Page Moved")
)

func main() {
	if err := webCall(true); err != nil {
		switch err {
			case ErrBadRequest:
				fmt.Println("Bad Request Occurred")
				return
			case ErrPageMoved:
				fmt.Println("The Page moved")
				return
			default:
				fmt.Println(err)
				return
		}
	}
	fmt.Println("Life is good")
}
```

## Custom Error 
- tự khai báo error, thường đặt suffix cuối là Error, dùng pointer receiver trong hàm Error

```
type UnmarshalTypeError struct {
	Value string
	Type reflect.Type
}
func (e *UnmarshalTypeError) Error() string {
	return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}
```

## Generic type assertion
- check error tự khai báo sử dụng ``` err.(type)  ```

```
func main() {
 	var u user
 	err := Unmarshal([]byte(`{"name":"bill"}`), u)
 	if err != nil {
 		switch e := err.(type) {
 			case \*UnmarshalTypeError:
 				fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
 			case \*InvalidUnmarshalError:
 				fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
 			default:
 				fmt.Println(err)
 		}
 		return
 	}
 	fmt.Println("Name:", u.Name)
 }
```


- có thể switch case để check xem error có phải thuộc 1 interface nào đó hay ko, giúp giảm các case và tâp trung vào handle 1 lỗi nào đó

```
type temporary interface {
	Temporary() bool
}

func (c \*client) BehaviorAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
				case temporary:
					if !e.Temporary() {
						log.Println("Temporary: Client leaving chat")
						return
					}
				default:
					if err == io.EOF {
						log.Println("EOF: Client leaving chat")
						return
					}
					log.Println("read-routine", err)
			}
		}
		fmt.Println(line)
	}
}
```

## Always Use The Error Interface

```
type customError struct{}

func (c *customError) Error() string {
	return "Find the bug."
}

func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error
	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail?")
	}
	log.Println("No Error")
}

Output:
Why did this fail?
```

- đoạn code trên mặc dù ko có lỗi nhưng vẫn nhảy vào trong error handle vì developer trong hàm fail khi khai báo parameter sử dụng customError, và trong main thì class err lại là interface error -> err sẽ không nil nữa mà có 1 kiểu dữ liệu. 
- Do đó, các hàm khi return luôn return interface ``` error ``` để xử lý lỗi


## Handling Errors
- sử dụng Dave Cheney’s errors package "github.com/pkg/errors"
- mục đích, wrap error lại và trace dễ dàng hơn
