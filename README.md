# Meepo
[![Telegram](https://img.shields.io/badge/Telegram-online-brightgreen.svg)](https://t.me/meepoDiscussion)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/PeerXu/meepo/pulls)

Meepo的目标是提供一个去中心化的, 便捷的连接工具.

**本项目还处于初期版本, 接口变动会相对频繁, 请留意.**


## 起因

在传统的客户端-服务端架构的网络中, 服务端所在的网络需要能够被客户端访问, 才能完成既定功能.

但是, 由于各种原因, 导致服务端没有足够的资源去暴露端口.

因此, 作者提供了一个工具, 使得客户端可以访问无法暴露端口的服务端.


## 快速入门

### 访问未暴露公有IP的SSH服务

1. 在SSH Server(无公网IP节点, 下文称之为S1)上, 初始化并且运行Meepo服务

`<CustomID>`替换成自己想叫的名字, 建议足够复杂防止与他人冲突.

```bash
$ meepo config init --id <CustomID>
$ meepo serve
```

通过whoami子命令可以校验Meepo服务启动是否成功.

```bash
$ meepo whoami
<CustomID>
```

2. 在SSH Client(下文称之为C1)上, 运行Meepo服务

由于C1不需要被其他Meepo服务访问, 所以ID不重要, 采用默认配置即可.

```bash
$ meepo serve
```

 通过whoami子命令可以校验Meepo服务启动是否成功.

```bash
$ meepo whoami
<RandomID>
```

因为在初始化时没有指定ID, 所以Meepo服务启动时, 会随机指定一个ID.

3. 在C1通过SSH连接到S1

```bash
$ eval $(meepo ssh -- <Username>@<CustomID>)
```

等待片刻, SSH Client就会连接上SSH Server.

4. 关闭Meepo服务

在SSH断开连接后, 可以关闭Meepo服务.

```bash
$ meepo shutdown
```


### 访问未暴露共有IP的HTTP服务

1. 在HTTP Server(下文称之为H1)上, 初始化并且运行Meepo服务.

同上, 请参考上文过程.

假设H1的监听地址为 `0.0.0.0:8080`

2. 在HTTP Client(下文称之为 C1)上, 运行Meepo服务.

同上, 请参考上文过程.

3. 在C1上建立连接

```bash
$ meepo teleport --id <CustomID> --name http --remote-address 127.0.0.1:8080
# sometime later...
# output:
Teleport SUCCESS
Enjoy your teleportation with 127.0.0.1:65311
```

这时候已经成功建立连接, 可以通过 `http://127.0.0.1:65311` 访问HTTP服务.

4. 查看连接情况

```bash
$ meepo teleportation list
# output:
Name    Transport       Portal  Source  Sink    Channels
http    b       source  tcp:127.0.0.1:65311     tcp:127.0.0.1:8080      0
```

5. 关闭连接

```bash
$ meepo teleportation close --name http
# output:
Teleportation closing
```


## 计划

如果有不错的想法, 不妨通过[Telegram](https://t.me/meepoDiscussion)或[issues](https://github.com/PeerXu/meepo/issues)留言.

### 短期计划

- [x] SSH连接端口复用
- [ ] 缩短gather时间
- [ ] 工作原理文档的补全
- [ ] 中英文档的补全

### 长期计划

- [x] 连接变得可管理
- [ ] 支持socks5 proxy
- [ ] 支持http proxy
- [x] 支持port forward


## 为Meepo做贡献

Meepo是一个免费且开源的项目, 欢迎任何人为其开发和进步贡献力量.

* 在使用过程中出现任何问题, 可以通过[issues](https://github.com/PeerXu/meepo/issues)来反馈.
* 在使用过程中出现任何问题, 也可以通过[Telegram](https://t.me/meepoDiscussion)来沟通使用心得.
* 如果还有其他方面的问题与合作, 欢迎联系 pppeerxu@gmail.com .


### 代码提交

* main分支仅作用于稳定版本的发布, [PRs](https://github.com/PeerXu/meepo/pulls)请提交到dev分支.
* Bug修复可以直接提交PR到dev分支.
* 如果有新增功能的想法, 可以先到[issues](https://github.com/PeerXu/meepo/issues)描述想法与对应的实现, 然后fork修改, 最后提交PR到dev分支进行合并.


## 捐赠

如果觉得Meepo能够帮助到你, 欢迎提供适当的捐助来维持项目的长期发展.

### Telegram

[https://t.me/meepoDiscussion](https://t.me/meepoDiscussion)

### BTC

![BTC](./donations/btc.png)

36PnaXCMCtKLbkzVyfrkudhU6u8vjbfax4

### ETH

![ETH](./donations/eth.png)

0xa4f00EdD5fA66EEC124ab0529cF35a64Ee94BFDE


## 贡献者

[PeerXu](https://github.com/PeerXu) (pppeerxu@gmail.com)


## License

MIT
