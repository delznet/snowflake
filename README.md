# snowflake ID生成系统

服务端用golang生成,关于snowflake的算法参考[这里](https://www.cnblogs.com/relucent/p/4955340.html)

## 使用方法

1. 进入server目录

go build -o ../bin/id-server

2. 将server目录下app.conf拷贝到../bin目录

cp app.conf ../bin

3. 修改app.conf

```
{
    "port":":7890",  //服务端口号
    "epoch":1514736000000, //开始时间戳
    "machine_id_bits":4, //机器号位长度
    "service_id_bits":6, //业务位长度
    "step_bits":12, //序列位长度
    "machine_id":0, //机器id
    "service_id":0 //业务id
}
```

说明: 
（1）machine_id_bits+service_id_bits+step_bits 必须等于22
（2）如果machine_id_bits为4，machine_id的取值范围是0-15（2的4次方），service_id也是同理。
（3）不同启动实例的machine_id必须不同，service_id可以重复。
（4）可以通过计算service_id,获取不同id的业务类型，比如有两种编码，一种表示位置编码，一种表示设备编码，通过设置不同的service_id可以区分。

4. 启动server

## 客户端

1. golang客户端可参考client目录
2. [php客户端代码](https://github.com/delznet/snowflake-php-sdk)