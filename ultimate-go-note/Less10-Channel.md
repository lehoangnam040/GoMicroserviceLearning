## Semantics

- think channel not a data structure, but as a mechanic for signaling.

- channel có 2 loại:
	- unbuffered channel: guarantee phải có receive ở 1 goroutine khi 1 goroutine send signal.
	- buffered channel: có 1 buffer gồm X phần tử, nếu chưa full thì có thể send vào buffer này, còn ko sẽ bị block. Nếu buffer có data thì có thể receive còn ko sẽ bị block

- channel có 3 state:
	- nil: khi construct với zero values. send và receive sẽ bị block 
	- open: dùng hàm ``` make```
	- close: dùng hàm ``` close```. send vào 1 closed channel gây panic

# Các basic patterns

## Wait For Result
- goroutine A tạo ra 1 goroutine B, sau đó A đợi kết quả từ B trả về.


## Wait For Task
- Goroutine A tạo 1 goroutine B, B nhận data từ A để làm việc

## Wait for Finish
- Goroutine A tạo 1 goroutine B, B có thể close để A biết 


# Các patterns đc build từ 3 pattern cơ bản trên

## Fan Out/In
- được build từ Wait For Result 
- 1 goroutine cha có thể tạo X goroutine con, sau đó cha sẽ đợi kết quả từ các con dựa vào buffered channel với cap = X

```
func fanOut() {
	children := 2000
	ch := make(chan string, children)
	
	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}
}
```

## Pooling
- build từ Wait For Task
- tạo ra 1 pooling gồm X goroutine con. Gorountine cha sẽ có 1 list công việc để truyền vào và X con sẽ random nhận data và xử lý công việc. 

```
func pooling() {
	ch := make(chan string)
	g := runtime.GOMAXPROCS(0)
	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, d)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent : sent signal :", w)
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
}
```

## Drop
- build từ Wait For Task
- tạo 1 buffered channel gồm X phần tử, goroutine cha liên tục gửi data để goroutine con xử lý, nếu quá buffer thì sẽ ko block send mà drop send, sử dụng switch 
- phù hợp cho những bài toán xử lý trên network mà đôi khi phải drop request khi quá nhiều


```
func drop() {

	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
			case ch <- "data":
				fmt.Println("parent : sent signal :", w)
			default:
				fmt.Println("parent : dropped data :", w)
		}
	}
	close(ch)

	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

```

## Cancellation
- build từ Wait for result 
- tạo 1 channel để làm timeout, khi timeout hết thì goroutine thực thi nhiệm vụ cũng sẽ hết

```
func cancellation() {

	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)

	defer cancel()

	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()

	select {
		case d := <-ch:
			fmt.Println("work complete", d)
		case <-ctx.Done():
			fmt.Println("work cancelled")
	}
}
```

## Fan Out/In Semaphore
- Mở rộng của fan out/in, tuy nhiên thêm 1 buffered channel để xác định số lượng goroutine con được phép chạy đồng thời

```
func fanOutSem() {
	children := 2000
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- "data"
				fmt.Println("child : sent signal :", child)
			}
			<-sem
		}(c)
	}
	
	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

```


## Bounded Work Pooling

```
func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}
	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)
	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()
}

```

## Retry Timeout
- retry sau khi đã timeout

```
func retryTimeout(ctx context.Context, retryInterval time.Duration, check func(ctx context.Context) error) {

	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}
		fmt.Println("check if timeout has expired")
		
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}

		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)
	
		select {
			case <-ctx.Done():
				fmt.Println("timed expired 2 :", ctx.Err())
				t.Stop()
				return
			case <-t.C:
				fmt.Println("retry again")
		}
	}
}

```


## Channel Cancellation

```
func channelCancellation(stop <-chan struct{}) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
			case <-stop:
				cancel()
			case <-ctx.Done():
		}
	}()

	func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			"https://www.ardanlabs.com/blog/index.xml",
			nil,
		)
		if err != nil {
			return err
		}
		
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		
		return nil
	}(ctx)
}
```
