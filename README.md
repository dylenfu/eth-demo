以太坊demo

功能
 安装
 创世块
 智能合约
 
安装
 下载go-eth项目,使用源码安装,参考：
 https://ethereum.github.io/go-ethereum/install/#build-it-from-source-code
 安装solc,使用binary package安装,参考：
 http://solidity.readthedocs.io/en/latest/installing-solidity.html
 
创世块
 新建build/seed0/genesis.json
 建立5个初始账号
 ```json
    {
       "config": {
         "chainId": 1,
         "homesteadBlock": 0,
         "eip155Block": 0,
         "eip158Block": 0
       },
       "alloc": {
         "3ae88fe370c39384fc16da2c9e768cf5d2495b48": {
           "balance": "10000000000"
         },
         "81063419f13cab5ac090cd8329d8fff9feead4a0": {
           "balance": "10000000000"
         },
         "9da26fc2e1d6ad9fdd46138906b0104ae68a65d8": {
           "balance": "10000000000"
         },
         "bd2d69e3e68e1ab3944a865b3e566ca5c48740da": {
           "balance": "10000000000"
         },
         "ca9f427df31a1f5862968fad1fe98c0a9ee068c4": {
           "balance": "10000000000"
         }
       },
     
       "nonce": "0x0000000000000042",
       "difficulty": "0x020000",
       "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
       "coinbase": "0x0000000000000000000000000000000000000000",
       "timestamp": "0x00",
       "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
       "extraData": "0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa",
       "gasLimit": "0x4c4b40"
     }
```
 
 初始化：
 ```bash
 vagrant@vagrant:~/gohome/src/eth-demo/build/seed0$ mkdir data
 vagrant@vagrant:~/gohome/src/eth-demo/build/seed0$ ll
 总用量 16
 drwxrwxr-x 3 vagrant vagrant 4096 7月  28 15:00 ./
 drwxrwxr-x 3 vagrant vagrant 4096 7月  28 11:35 ../
 drwxrwxr-x 2 vagrant vagrant 4096 7月  28 15:00 data/
 -rw-rw-r-- 1 vagrant vagrant  990 7月  28 14:35 genesis.json
 
 vagrant@vagrant:~/gohome/src/eth-demo/build/seed0$ geth --datadir data init genesis.json
 WARN [07-28|15:00:39] No etherbase set and no accounts found as default 
 INFO [07-28|15:00:39] Allocated cache and file handles         database=/home/vagrant/gohome/src/eth-demo/build/seed0/data/geth/chaindata cache=16 handles=16
 INFO [07-28|15:00:39] Writing custom genesis block 
 INFO [07-28|15:00:39] Successfully wrote genesis state         database=chaindata                                                         hash=c5f7a6…dd680c
 INFO [07-28|15:00:39] Allocated cache and file handles         database=/home/vagrant/gohome/src/eth-demo/build/seed0/data/geth/lightchaindata cache=16 handles=16
 INFO [07-28|15:00:39] Writing custom genesis block 
 INFO [07-28|15:00:39] Successfully wrote genesis state         database=lightchaindata                                                         hash=c5f7a6…dd680c
 ```
 启动节点:
 ```bash
 vagrant@vagrant:~/gohome/src/eth-demo/build/seed0$ geth --datadir data --networkid 1 --rpc --rpccorsdomain "*" --nodiscover console
 WARN [07-28|15:05:30] No etherbase set and no accounts found as default 
 INFO [07-28|15:05:30] Starting peer-to-peer node               instance=Geth/v1.7.0-unstable/linux-amd64/go1.8.1
 INFO [07-28|15:05:30] Allocated cache and file handles         database=/home/vagrant/gohome/src/eth-demo/build/seed0/data/geth/chaindata cache=128 handles=1024
 WARN [07-28|15:05:30] Upgrading database to use lookup entries 
 INFO [07-28|15:05:30] Database deduplication successful        deduped=0
 INFO [07-28|15:05:30] Initialised chain configuration          config="{ChainID: 1 Homestead: 0 DAO: <nil> DAOSupport: false EIP150: <nil> EIP155: 0 EIP158: 0 Metropolis: <nil> Engine: unknown}"
 INFO [07-28|15:05:30] Disk storage enabled for ethash caches   dir=/home/vagrant/gohome/src/eth-demo/build/seed0/data/geth/ethash count=3
 INFO [07-28|15:05:30] Disk storage enabled for ethash DAGs     dir=/home/vagrant/.ethash                                          count=2
 WARN [07-28|15:05:30] Upgrading db log bloom bins 
 INFO [07-28|15:05:30] Bloom-bin upgrade completed              elapsed=60.801µs
 INFO [07-28|15:05:30] Initialising Ethereum protocol           versions="[63 62]" network=1
 INFO [07-28|15:05:30] Loaded most recent local header          number=0 hash=c5f7a6…dd680c td=131072
 INFO [07-28|15:05:30] Loaded most recent local full block      number=0 hash=c5f7a6…dd680c td=131072
 INFO [07-28|15:05:30] Loaded most recent local fast block      number=0 hash=c5f7a6…dd680c td=131072
 INFO [07-28|15:05:30] Starting P2P networking 
 INFO [07-28|15:05:30] HTTP endpoint opened: http://127.0.0.1:8545 
 INFO [07-28|15:05:30] RLPx listener up                         self="enode://46144da688e40dd7a1209147c1085eb2f08ef4efa03553d51676bb857e2e81539292b2df9b426760416b05c81943e59d4397dac72f57e989905de4008d8bd27b@[::]:30303?discport=0"
 INFO [07-28|15:05:30] IPC endpoint opened: /home/vagrant/gohome/src/eth-demo/build/seed0/data/geth.ipc 
 Welcome to the Geth JavaScript console!
 
 instance: Geth/v1.7.0-unstable/linux-amd64/go1.8.1
  modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
 
```
 
 上述命令中：
 datadir 指定数据目录
 rpc     启动rpc通讯,可以进行调试和部署智能合约
 
 控制台运行相关命令,这里以挖矿为例子：
```cmd
 > personal.newAccount("1")
 INFO [07-28|15:11:00] New wallet appeared                      url=keystore:///home/vagrant/gohome… status=Locked
 "0x577e2b2ae9d3bce1eaf60f735425a2e5fc67093e"
 > miner.start(1)
 INFO [07-28|15:11:04] Updated mining threads                   threads=1
 INFO [07-28|15:11:04] Transaction pool price threshold updated price=18000000000
 null
 > INFO [07-28|15:11:04] Starting mining operation 
 INFO [07-28|15:11:04] Commit new mining work                   number=1 txs=0 uncles=0 elapsed=303.736µs
 INFO [07-28|15:11:07] Generating DAG in progress               epoch=0 percentage=0 elapsed=1.530s
 INFO [07-28|15:11:08] Generating DAG in progress               epoch=0 percentage=1 elapsed=3.065s
 ......
 > miner.stop()
 true
```
 

 