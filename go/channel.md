# GO Channel Note

##### **C*ontext*** 主要用來在 goroutine 之間傳遞上下文訊息, 包括取消訊號, 超時, 截止時間, key-value 等等

- 看看 context 的結構
    ```
    type Context interface {

        Deadline() (deadline time.Time, ok bool)

        Done() <-chan struct{}

        Err() error

        Value(key interface{}) interface{}
    }
    ```

- background 通常用在 main function 中, 作為所有 context 的根節點
    ```
    type emptyCtx int

    func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
        return
    }

    func (*emptyCtx) Done() <-chan struct{} {
        return nil
    }

    func (*emptyCtx) Err() error {
        return nil
    }

    func (*emptyCtx) Value(key interface{}) interface{} {
        return nil
    }

    var (
    background = new(emptyCtx)
    todo       = new(emptyCtx)
    )

    func Background() Context {
        return background
    }

    func TODO() Context {
        return todo
    }
    ```