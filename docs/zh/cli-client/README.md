# 命令行客户端

## 查询命令的flags

| 名称, 速记       | 类型    | 是否必须  | 默认值                 | 描述                                                    | 
| --------------- | ----   | -------- | --------------------- | ------------------------------------------------------- |
| --chain-id      | string |          |                       | tendermint节点的Chain ID                                 |
| --height        | int    |          | 0                     | 查询某个高度的区块链数据，如果是0，则返回最新的区块链数据        |
| --help, -h      | string |          |                       | 打印帮助信息                                              |
| --output        | string |          | text                  | 指定返回的格式 text或者json                                |
| --indent        | bool   |          | false                 | 格式化json字符串                                          |
| --ledger        | bool   |          | false                 | 使用ledger设备                                            |
| --node          | string |          | tcp://localhost:26657 | tendermint节点的rpc地址                                   |
| --trust-node    | bool   |          | true                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性 |

每个区块链状态查询命令都包含上表中的flags，同时不同查询命令还可能会有自己独有的flags。

### 格式化json返回

`output`字段可以指定查询的返回格式：

不指定，默认以text方式返回：

```
root@ubuntu:~# iriscli stake validators
Validator
  Operator Address:            iva1gfcee5u5f54kfcnufv4ypcfyldw0vu0zpwl52q
  Validator Consensus Pubkey:  icp1zcjduepquednrr0aqw4nkt8jnkhpmg4acfc7vlr0yre4uud4z0ups68hcpfsx4x9ng
  Jailed:                      false
  Status:                      Bonded
  Tokens:                      1361.0004000000246900000000000000
  Delegator Shares:            1361.0004000000246900000000000000
  Description:                 {B-2  3_C a1_}
  Unbonding Height:            0
  Minimum Unbonding Time:      1970-01-01 00:00:00 +0000 UTC
  Commission:                  rate: 0.1001000000, maxRate: 1.0000000000, maxChangeRate: 1.0000000000, updateTime: 2019-05-09 03:13:39.720700953 +0000 UTC
```

指定`output`和`indent`，以格式化后的json方式返回：

```
root@ubuntu:~# iriscli stake validators --output=json --indent
[
  {
    "operator_address": "iva1gfcee5u5f54kfcnufv4ypcfyldw0vu0zpwl52q",
    "consensus_pubkey": "icp1zcjduepquednrr0aqw4nkt8jnkhpmg4acfc7vlr0yre4uud4z0ups68hcpfsx4x9ng",
    "jailed": false,
    "status": 2,
    "tokens": "1361.0004000000246900000000000000",
    "delegator_shares": "1361.0004000000246900000000000000",
    "description": {
      "moniker": "B-2",
      "identity": "",
      "website": "3_C",
      "details": "a1_"
    },
    "bond_height": "0",
    "unbonding_height": "0",
    "unbonding_time": "1970-01-01T00:00:00Z",
    "commission": {
      "rate": "0.1001000000",
      "max_rate": "1.0000000000",
      "max_change_rate": "1.0000000000",
      "update_time": "2019-05-09T03:13:39.720700953Z"
    }
  }
]
```

## 发送交易命令的flags


| 名称, 速记        | 类型    | 是否必须  | 默认值                 | 描述                                                                                                |
| -----------------| -----  | -------- | --------------------- | --------------------------------------------------------------------------------------------------- |
| --account-number | int    |          | 0                     | 发起交易的账户的编号                                                                                   |
| --async          | bool   |          | false                 | 是否异步广播交易（不对交易进行任何验证，立即返回交易的hash，仅当commit为false时有效）                           |
| --commit         | bool   |          | false                 | 广播交易并等到交易被打包再返回                                                                           |
| --chain-id       | string | 是       |                       | tendermint节点的`Chain ID`                                                                           |
| --dry-run        | bool   |          | false                 | 模拟执行交易，并返回消耗的`gas`。`--gas`指定的值会被忽略                                                    |
| --fee            | string | 是       |                       | 交易费（指定交易费的上限）                                                                               |
| --from           | string |          |                       | 发送交易的账户名称                                                                                      |
| --from-addr      | string |          |                       | 签名地址，在`generate-only`为`true`的情况下有效                                                          |
| --gas            | int    |          | 50000                 | 交易的gas上限; 设置为"simulate"将自动计算相应的阈值                                                        |
| --gas-adjustment | int    |          | 1.5                   | gas调整因子，这个值降乘以模拟执行消耗的`gas`，计算的结果返回给用户; 如果`--gas`的值不是`simulate`，这个标志将被忽略 |
| --generate-only  | bool   |          | false                 | 是否仅仅构建一个未签名的交易便返回                                                                         |
| --help, -h       | string |          |                       | 打印帮助信息                                                                                           |
| --indent         | bool   |          | false                 | 格式化json字符串                                                                                       |
| --json           | string |          | false                 | 指定返回结果的格式，`json`或者`text`                                                                     |
| --ledger         | bool   |          | false                 | 使用ledger设备                                                                                        |
| --memo           | string |          |                       | 指定交易的memo字段                                                                                     |
| --node           | string |          | tcp://localhost:26657 | tendermint节点的rpc地址                                                                                |
| --print-response | bool   |          | false                 | 是否打印交易返回结果，仅在`async`为true的情况下有效                                                         |
| --sequence       | int    |          | 0                     | 发起交易的账户的sequence                                                                                |
| --trust-node     | bool   |          | true                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性                                               | 

每个发送交易的命令都包含上表中的flags，同时不同交易的命令还可能会有自己独有的flags。

## 模块命令列表

1. [status 命令](./status/README.md)
2. [tendermint 命令](./tendermint/README.md)
3. [keys 命令](./keys/README.md)
4. [bank 命令](./bank/README.md)
5. [stake 命令](./stake/README.md)
6. [distribution 命令](./distribution/README.md)
7. [gov 命令](./gov/README.md)
8. [upgrade 命令](./upgrade/README.md)
9. [service 命令](./service/README.md)

## 配置命令

`iriscli config` 命令可以交互式地配置一些默认参数，例如chain-id，home，fee 和 node。

例子：

```
root@ubuntu16:~# iriscli config
> Where is your iriscli home directory? (Default: ~/.iriscli)
/root/my_cli_home
> Where is your validator node running? (Default: tcp://localhost:26657)
tcp://192.168.0.1:26657
Do you trust this node? [y/n]:y
> What is your chainID?
irishub
> Please specify default fee
50000

root@ubuntu16:~# iriscli status --home=/root/my_cli_home
```
