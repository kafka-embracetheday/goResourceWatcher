### 系统资源指标

#### 1、CPU

##### a.定义及解释

- 中央处理器是一块超大规模的集成电路，是一台计算机的运算核心（Core）和控制核心（ Control Unit）。它的功能主要是解释计算机指令以及处理计算机软件中的数据。CPU Load：系统正在干活的多少的度量，队列长度。系统平均负载。

##### b.监控指标

CPU指标主要指的CPU使用率、利用率，包括用户态（user）、系统态（sys）、等待态（wait）、空闲态（idle）。CPU使用率、利用率要低于业界警戒值范围之内，即小于或者等于75%、CPU sys%小于或者等于30%，CPU wait%小于或者等于5%。单核CPU也需遵循上述指标要求。CPU Load要小于CPU核数。CPU 使用率反映了 CPU 的工作强度,是一个百分比指标。

- CPU使用率：CPU 使用率指的是 CPU 在特定时间内被实际使用的百分比。它反映了 CPU 的工作强度或繁忙程度,是评估 CPU 性能的一个重要指标。CPU 使用率可以从 0% 到 100% 不等,100% 表示 CPU 满负荷运行。CPU 利用率则包括了用户模式和内核模式下 CPU 的使用情况。

- CPU利用率：CPU 利用率则是指系统中所有进程（包括系统进程）占用 CPU 的总百分比。它包括了 CPU 在用户模式和内核模式下的使用情况。CPU 利用率可能超过 100%,因为多个进程可能同时占用 CPU。

- CPU负载：CPU 负载是指在特定时间内系统队列中等待 CPU 资源的进程数。它反映了 CPU 当前的工作压力情况,包括正在运行的进程以及正在等待运行的进程。CPU 负载通常用 1 分钟、5 分钟和 15 分钟的平均值来表示。CPU 负载可以大于 CPU 核心数,表示系统存在处理能力不足的情况。