## Khái niệm
- khi program chạy, Go runtime hỏi máy tính số system thread can run parallel. dựa vào số core 

- Với mỗi thread, runtime tạo ra 1 OS thread (M) và đính nó vào 1 logical processor (P). 

- 1 số định nghĩa:
	- Work: A set of instructions to be executed for a running application. This is
accomplished by threads and an application can have 1 to many threads.
	- Thread: A path of execution that is scheduled and performed. Threads are
responsible for the execution of instructions on the hardware.
	- Thread States: A thread can be in one of three states: Running, Runnable, or
Waiting. Running means the thread is executing its assigned instructions on the
hardware by having a G placed on the M. Runnable means the thread wants time on
the hardware to execute its assigned instructions and is sitting in a run queue.
Waiting means the thread is waiting for something before it can resume its work.
Waiting threads are not a concern of the scheduler.
	- Concurrency: la kha nang 1 chuong trinh co the dieu phoi nhieu tac vu trong 1 khoang tgian va chi cho phep 1 tac vu trong 1 thoi diem. Concurrency is about "dealing" with lots of things at once
	- Parallelism: parallelism la kha nang 1 chuong trinh co the thuc hien 2 hoac nhieu tac vu trong cung 1 thoi diem. Parallelism is about "doing" lots of things at once
	- CPU Bound Work: This is work that does not cause the thread to naturally move
into a waiting state. Calculating fibonacci numbers would be considered CPU-Bound
work.
	- I/O Bound Work: This is work that does cause the thread to naturally move into a
waiting state. Fetching data from different URLs would be considered I/O-Bound
work.
	- Synchronization: When two or more Goroutines will need to access the same
memory location potentially at the same time, they need to be synchronized and
take turns. If this synchronization doesn’t take place, and at least one Goroutine is
performing a write, I can end up with a data race. Data races are a cause of data
corruption bugs that can be difficult to find.
	- Orchestration: When two or more Goroutines need to signal each other, with or
without data, orchestration is the mechanic required. If orchestration does not take
place, guarantees about concurrent work being performed and completed will be missed. This can cause all sorts of data corruption bugs.


## Concurrency Basics
``` runtime.GOMAXPROCS(x)  ``` 
- nếu x > 0 thì tạo ra x cặp M/P . Function trên override env variable
- x = 0 thì tự xác định 


```
func main() {
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	go func() {
		f1()
		wg.Done()
	}()
	
	go func() {
		f2()
		wg.Done()
	}()
	
	fmt.Println("Waiting To Finish")
	wg.Wait()
	
	fmt.Println("\nTerminating Program")
}
```

- Cách quản lý này sử dụng WaitGroup, ```wg.Add(2)``` -> có 2 goroutine sẽ đc tạo
- wg.Wait() block đợi 2 goroutine chạy xong hết thì mới chạy câu tiếp theo để kết thúc ctrinh
- wg.Done() để xác định goroutine đó chạy xong

- Logic f1 và f2 sẽ chạy concurrency, nghĩa là f1 có thể chạy 1 vài logic -> waiting -> f2 chạy -> f1 chạy -> ...




