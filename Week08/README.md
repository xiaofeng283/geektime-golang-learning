
# 第八周作业

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

测试机器： MacBook Pro 2015，四核16G

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379  -t get,set -q -n 100000 -c 100 -d 10
SET: 75700.23 requests per second, p50=0.663 msec
GET: 76511.09 requests per second, p50=0.663 msec

$ redis-benchmark -h 127.0.0.1 -p 6379  -t get,set -q -n 100000 -c 100 -d 20
SET: 76569.68 requests per second, p50=0.655 msec
GET: 77821.02 requests per second, p50=0.647 msec

$ redis-benchmark -h 127.0.0.1 -p 6379  -t get,set -q -n 100000 -c 100 -d 50
SET: 74074.07 requests per second, p50=0.679 msec
GET: 74571.22 requests per second, p50=0.671 msec

$ redis-benchmark -h 127.0.0.1 -p 6379  -t get,set -q -n 100000 -c 100 -d 100
SET: 74682.60 requests per second, p50=0.671 msec
GET: 75471.70 requests per second, p50=0.671 msec

$ redis-benchmark -h 127.0.0.1 -p 6379  -t get,set -q -n 100000 -c 100 -d 200
SET: 77101.00 requests per second, p50=0.655 msec
GET: 76923.08 requests per second, p50=0.655 msec

$ redis-benchmark -h 127.0.0.1 -p 6379  -t get,set -q -n 100000 -c 100 -d 1000
SET: 77519.38 requests per second, p50=0.663 msec
GET: 76687.12 requests per second, p50=0.663 msec

$ redis-benchmark -h 127.0.0.1 -p 6379  -t get,set -q -n 100000 -c 100 -d 5000
SET: 73099.41 requests per second, p50=0.695 msec
GET: 72150.07 requests per second, p50=0.703 msec
```

在10万的请求量、100并发的情况下，随着value大小的增加，set、get的性能有些许下降。


2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。




info memory：
```shell
$ redis-cli
127.0.0.1:6379> info memory
# Memory
used_memory:1064080
used_memory_human:1.01M
used_memory_rss:3801088
used_memory_rss_human:3.62M
used_memory_peak:12871264
used_memory_peak_human:12.27M
used_memory_peak_perc:8.27%
used_memory_overhead:1026672
used_memory_startup:1009232
used_memory_dataset:37408
used_memory_dataset_perc:68.20%
allocator_allocated:1027152
allocator_active:3763200
allocator_resident:3763200
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:3.66
allocator_frag_bytes:2736048
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.01
rss_overhead_bytes:37888
mem_fragmentation_ratio:3.70
mem_fragmentation_bytes:2773936
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

使用shell脚本写入测试数据：20字节
```shell
#!/bin/bash
for i in {1..10000}
do
  echo "key${i}"
  redis-cli -c -h 127.0.0.1 -p 6379 set key${i} aaaaaaaaaaaaaaaaaaa
done
```

info memory:
```shell
127.0.0.1:6379> info memory
# Memory
used_memory:2155152
used_memory_human:2.06M
used_memory_rss:3977216
used_memory_rss_human:3.79M
used_memory_peak:12871264
used_memory_peak_human:12.27M
used_memory_peak_perc:16.74%
used_memory_overhead:1557744
used_memory_startup:1009232
used_memory_dataset:597408
used_memory_dataset_perc:52.13%
allocator_allocated:2118224
allocator_active:3939328
allocator_resident:3939328
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.86
allocator_frag_bytes:1821104
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.01
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.88
mem_fragmentation_bytes:1858992
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

寻找较大的key：
```shell
redis-cli -h 127.0.0.1 -p 6379 --bigkeys

# Scanning the entire keyspace to find biggest keys as well as
# average sizes per key type.  You can use -i 0.1 to sleep 0.1 sec
# per 100 SCAN commands (not usually needed).

[00.00%] Biggest string found so far '"key9135"' with 19 bytes

-------- summary -------

Sampled 10000 keys in the keyspace!
Total key length in bytes is 68894 (avg len 6.89)

Biggest string found '"key9135"' has 19 bytes

0 lists with 0 items (00.00% of keys, avg size 0.00)
0 hashs with 0 fields (00.00% of keys, avg size 0.00)
10000 strings with 190000 bytes (100.00% of keys, avg size 19.00)
0 streams with 0 entries (00.00% of keys, avg size 0.00)
0 sets with 0 members (00.00% of keys, avg size 0.00)
0 zsets with 0 members (00.00% of keys, avg size 0.00)
```

还有个比较笨的方式，用rdbtools工具来查看key的大小，这样比较麻烦，数量少比较好操作，数量多的话就不好操作了。这里测试的是1万的测试数据，勉强可以用这种方式查看。



