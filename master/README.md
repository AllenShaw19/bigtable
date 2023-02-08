# Master

控制平面

Master 控制了 Tablet 主动切主、被动选主、迁移、分裂与合并的具体流程控制，

此外 Master 向 Client 暴露 Tablet 分布的位置信息。

# 选主流程
当 Master 在指定时间内没有收到某 TabletServer 的心跳消息，则将此 TabletServer 判定为 Offline，

然后对此 TabletServer 上为 Leader 的 Tablet 发起一轮选主。

选主的过程与 Raft 类似仅仅是将这部分逻辑上移到 Master 统一管理，

首先收集 Tablet 各个副本的最新日志号，

然后选取足够新的副本作为选主目标，向 Tablet 各个副本发起投票，当投票成功后向选主目标发送当主请求。

