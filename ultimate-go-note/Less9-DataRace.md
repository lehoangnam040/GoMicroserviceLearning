
- data race xảy ra khi nhiều Goroutine đang access memory location at the same time, và ít nhất 1 goroutine đang ghi dữ liệu. 
- nó nghiêm trọng vì khó dự đoán output, cũng như random lỗi, do chạy goroutine concurrency

## Detection

- Cách 1, dùng ``` -race ``` khi chạy go build hoặc go test. tuy nhiên program sẽ chậm hơn

## Cách fix 1: Atomics

- dùng package sync/atomic để access dữ liệu cần share giữa các goroutine
- package này có rất nhiều hàm, tuy nhiên chỉ thực hiện trên kiểu dữ liệu int32, int64


## Cách fix 2: Mutexes

- dùng mutex.Lock() và Unlock() để khóa 1 đoạn code lại, chỉ có 1 goroutine được thực hiện tại 1 thời điểm. 
- sẽ làm tăng latency nếu đoạn code trong khi lock cần phải switch context, hoặc gọi systemcall nhiều lần

- dùng read/write mutex sẽ cho phép nhiều goroutine cùng đọc tại 1 thời điểm, tuy nhiên nếu việc ghi xảy ra thì việc đọc có thể phải thực hiện lại


