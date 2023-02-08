# TabletServer
数据平面

TabletServer 承载了多个 Tablet 的数据存储，并负责进行实际数据的读写操作

每一个 Tablet 都会与其他 TabletServer 上面的对应 Tablet 会组成一个复制组，复制组内的多个成员之间使用简化的 Raft 复制协议进行数据复制。

>所谓简化的 Raft 复制协议就是只实现了 Log Replication 的功能，
>将 Leader Election 和 Member Change 的功能移到了 Master 端。
>这样做分离了其数据平面和控制平面，实现了数据平面的逻辑简化。
>与此同时将控制平面操作统一到 Master 中便于定制更复杂的选主策略，
>并借此实现了多组 Raft 的心跳合并到 Master 减小了心跳的额外消耗。

## 数据模型

## 分裂与合并


