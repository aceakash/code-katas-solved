[From here](https://github.com/saltpay/code-katas/discussions/29)

In this kata, we will implement a "Stack" data structure. In later katas, we will _use_ it.

You can visualise a stack as a stack of plates. You can `push` a plate onto an existing stack of plates, and the stack grows. But you can only push to the top of the stack, you cannot "insert" a plate somewhere in between the stack.

Also, you can `pop` (remove) a plate only from the top of a stack. Popping a stack removes the topmost item from the stack and returns it.

You can probably see why a stack is described as a LIFO (last-in-first-out) or a FILO (first-in-last-out) data structure.

![image](https://user-images.githubusercontent.com/1387900/173227897-305756d9-ad05-4ad5-b385-576c3fb62439.png)


### Your task

Implement a `Stack` data structure (type) with the following operations (methods):

* `new(capacity)` : creates a new stack with a max capacity
* `push(item)` : adds the given item to the top of the stack
* `pop() item` : takes the top item off the stack and returns it
* `peek() item` : returns the top item from the stack without removing it
* `isEmpty() bool` : checks if the stack is empty
* `isFull() bool` : checks if the stack is full