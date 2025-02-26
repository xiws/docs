# MySql事务隔离级别与锁机制

## 事务的特点

**ACID**：
- Atomicity（原子性）：一个事务（transaction）中的所有操作，要么全部完成，要么全部回退，不会结束在中间某个环节。
- Consistency（一致性）：在事务开始之前和事务结束以后，数据库的完整性没有被破坏。比如：破会了字段的唯一索引，字段类型不一致等
- Isolation（隔离性）：隔离性还有其他的称呼，如并发控制（concurrency control）、可串行化（serializability）、锁（locking）等。数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。
- Durability（持久性）：事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。

**隔离的等级:**
读未提交（Read uncommitted）：脏读指的就是在不同的事务下，当前事务可以读到另外事务未提交的数据，简单来说就是可以读到脏数据

读提交（read committed）：读到已经提交的数据，但是其违反了数据库事务一致性的要求
可重复读（repeatable read）：
串行化（Serializable）：

## 事务的类型
❑扁平事务（Flat Transactions）
❑带有保存点的扁平事务（Flat Transactions with Savepoints）
❑链事务（Chained Transactions）
❑嵌套事务（Nested Transactions）
❑分布式事务（Distributed Transactions）

## 锁的类型
通过锁定机制可以实现事务的隔离性要求，使得事务可以并发地工作。

共享锁（S Lock），允许事务读一行数据。
排他锁（X Lock），允许事务删除或更新一行数据。
意向锁 ，为了支持在不同粒度上进行加锁操作，InnoDB存储引擎支持一种额外的锁方式。包括：意向共享锁，意向排它锁

只有事务T1和事务T2都获取共享锁时才都能获取到锁，这个称为锁兼容。能不能兼容以第一个获取锁的来判断，第一个是排他锁，那就获取不到。

##  一致性非锁定读
一致性的非锁定读（consistent nonlocking read）是指InnoDB存储引擎
通过行多版本控制（multi versioning）的方式来读取当前执行时间数据库中
行的数据。如果读取的行正在执行DELETE或UPDATE操作，这时读取操作不会
因此去等待行上锁的释放。相反地，InnoDB存储引擎会去读取行的一个快照数据
![](./attachments/Pasted%20image%2020240528165540.png)

1. 事务隔离级别（读提交）READ COMMITTED 时，非锁定读读取最新的一份快照
2. 事务隔离级别（可重复读） REPEATABLE READ 时，非锁定读 读取事务开始时的行数据版本

## 一致性锁定读
在默认配置下，即事务的隔离级别为REPEATABLE READ模式下，InnoDB存储引擎的SELECT操作使用一致性非锁定读。

用户需要显式地对数据库读取操作进行加锁以保证数据逻辑的一致性。而这要求数据库支持加锁语句，即使是对于SELECT的只读操作。

InnoDB存储引擎对于SELECT语句支持两种一致性的锁定读（locking read）
操作：
SELECT…FOR UPDATE
SELECT…LOCK IN SHARE MODE

SELECT…FOR UPDATE对读取的行记录加一个X锁，其他事务不能对已锁定的行加上任何锁。
SELECT…LOCK IN SHARE MODE对读取的行记录加一个S锁，其他事务可以向被锁定的行加S锁，但是如果加X锁，则会被阻塞。


## 外键的特殊
外键主要用于引用完整性的约束检查。
对于外键值的插入或更新，首先需要查询父表中的记录，即SELECT父表。
但是对于父表的SELECT操作，不是使用一致性非锁定读的方式，因为这样会发
生数据不一致的问题，因此这时使用的是SELECT…LOCK IN SHARE MODE方
式，即主动对父表加一个S锁。如果这时父表上已经这样加X锁，子表上的操作
会被阻塞

## 行锁的3种算法
InnoDB存储引擎有3种行锁的算法，其分别是：
1. Record Lock：单个行记录上的锁。 
2. Gap Lock：间隙锁，锁定一个范围，但不包含记录本身
3. Next-Key Lock∶Gap Lock+Record Lock，锁定一个范围，并且锁定记录本身（是结合了Gap Lock和Record Lock的一种锁定算法）

当查询的索引含有唯一属性时，将其降级为Record Lock

InnoDB存储引擎采用Next-Key Locking 避免幻读： （READ COMMITTED事务级别）
SELECT*FROM t WHERE a＞2 FOR UPDATE 锁定的行时（2，正无穷], 所以其他没有插入的数据满足这个条件也是无法执行。