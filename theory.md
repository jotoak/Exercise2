Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> *An atomic operation is a operation that can not be blocked, it is beeing executed in full befor beeing interupted. *

### What is a critical section?
> *A critical section is a proteced region where only one process can be executet at once, nothing happens in parallel. *

### What is the difference between race conditions and data races?
> *A race condition is when a operation depends on interleaving of other operations. A data races is when to threads is trying to change the same variable.*

### What are the differences between semaphores, binary semaphores, and mutexes?
> *A mutex is locking mechanism that makse it so that only one thread at a time can acess varible. A binary semaphores is a count that can be either 1 or 0, this puts a thread on wait to until a other thread signals it done. Semaphores is used to keep a thread waiting until another thread is done with somthing. *

### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> *Channels makes it posible to communicate between threads, the flow of information i only one way from sender to receiver. A  mailbox makes i posible to sende more messenges and the can be prioritized afther they are sendt. Queues uses the first in first out prinsiple. All this types of communication makes i safer to sende messenges between threads. *

### List some advantages of using message passing over lock-based synchronization primitives.
> *Your answer here*

### List some advantages of using lock-based synchronization primitives over message passing.
> *Your answer here*