# iriscli bank token-stats

## 描述

查询token统计数据，包括全部流通的token，全部绑定的token，以及销毁的token。

## 使用方式

```
 iriscli bank token-stats <tokenId> <flags>
```

## 标志

| 命令，速记   | 类型   | 是否必须 | 默认值                | 描述                                      |
| ------------ | ------ | -------- | --------------------- | ----------------------------------------- |
| -h, --help   |        | 否       |                       | 帮助                                      |
| --chain-id   | String | 否       |                       | tendermint 节点Chain ID                     |
| --height     | Int    | 否       |                       | 查询的区块高度用于获取最新的区块。        |
| --indent     | String | 否       |                       | 在JSON响应中增加缩进                      |
| --ledger     | String | 否       |                       | 使用ledger设备                    |
| --node       | String | 否       | tcp://localhost:26657 | `<主机>:<端口>` 链上的tendermint rpc 接口。 |
| --trust-node | String | 否       | True                  | 不验证响应的证明                          |

## 

## 例子

### 查询token统计数据

```
iriscli bank token-stats iris
```

示例返回：

```
TokenStats:
  Loose Token:  
    denom:iris
    amount:1864477.596384156921391687
  Burned Token:
    denom:iris
    amount:7177.596384156921391687
  Bonded Token:  
    denom:iris
    amount:1857300.596384156921391687
```
