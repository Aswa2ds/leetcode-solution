## 内存

### 内存逃逸

#### 引发变量逃逸到堆情况

1. 方法内将局部变量指针返回

   ```go
   package main
   
   import "fmt"
   
   func foo() *int {
   	a := new(int)
   	*a = 0
   	return a
   }
   
   func main() {
   	escape := foo()
   	fmt.Println(*escape)
   }
   ```
   
## container
### heap
1. 堆第n个节点的父节点：n/2-1 左子节点：2n 右子节点：2n+1\
   堆节点n的父节点：(n-1)/2 左子节点：2n+1 右子节点：2n+2
2. 操作依赖接口，不依赖底层数据结构，对底层数据结构的操作都是通过继承接口实现的方法 \
3. 
