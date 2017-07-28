# 以太坊demo

## 功能
 . 安装
 . 创世块
 . 智能合约
 
## 安装
 下载go-eth项目,使用源码安装,参考：<br>
 https://ethereum.github.io/go-ethereum/install/#build-it-from-source-code
 安装solc,使用binary package安装,参考： <br>
 http://solidity.readthedocs.io/en/latest/installing-solidity.html
 
## 创世块
 新建build/seed0/genesis.json <br>
 建立5个初始账号
 ```json
    {
       "config": {
         "chainId": 150,
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
 vagrant@vagrant:~/gohome/src/eth-demo/build/seed0$ geth --datadir data --networkid 150 --rpc --rpccorsdomain "*" --nodiscover console
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
 
 上述命令中：<br>
 datadir 指定数据目录 <br>
 rpc     启动rpc通讯,可以进行调试和部署智能合约 <br>
 
 控制台运行相关命令,这里以挖矿为例子：
```cmd
 > personal.newAccount("1")
 INFO [07-28|15:37:31] New wallet appeared                      url=keystore:///home/vagrant/gohome… status=Locked
 "0x888aa85221873f7f180c8fd9d88d4f462a40da6c"
 > miner.start(1)
 INFO [07-28|15:37:40] Updated mining threads                   threads=1
 INFO [07-28|15:37:40] Transaction pool price threshold updated price=18000000000
 null
 INFO [07-28|15:37:40] Starting mining operation 
 > INFO [07-28|15:37:40] Commit new mining work                   number=1 txs=0 uncles=0 elapsed=136.516µs
 INFO [07-28|15:37:46] Successfully sealed new block            number=1 hash=a3d5e8…c7bcf6
 INFO [07-28|15:37:46] 🔨 mined potential block                  number=1 hash=a3d5e8…c7bcf6
 INFO [07-28|15:37:46] Commit new mining work                   number=2 txs=0 uncles=0 elapsed=205.795µs
 INFO [07-28|15:37:47] Successfully sealed new block            number=2 hash=86869f…b6b4ff
 INFO [07-28|15:37:47] 🔨 mined potential block                  number=2 hash=86869f…b6b4ff
 INFO [07-28|15:37:47] Commit new mining work                   number=3 txs=0 uncles=0 elapsed=175.83µs
 INFO [07-28|15:37:48] Successfully sealed new block            number=3 hash=89bf98…db5baa
 INFO [07-28|15:37:48] 🔨 mined potential block                  number=3 hash=89bf98…db5baa
 INFO [07-28|15:37:48] Commit new mining work                   number=4 txs=0 uncles=0 elapsed=145.061µs
 INFO [07-28|15:37:52] Successfully sealed new block            number=4 hash=0177f9…b36645
 INFO [07-28|15:37:52] 🔨 mined potential block                  number=4 hash=0177f9…b36645
 INFO [07-28|15:37:52] Commit new mining work                   number=5 txs=0 uncles=0 elapsed=154.19µs
 INFO [07-28|15:37:53] Successfully sealed new block            number=5 hash=97dbde…e6621d
 INFO [07-28|15:37:53] 🔨 mined potential block                  number=5 hash=97dbde…e6621d
 INFO [07-28|15:37:53] Commit new mining work                   number=6 txs=0 uncles=0 elapsed=173.534µs
 INFO [07-28|15:37:56] Successfully sealed new block            number=6 hash=596c60…e2c6b4
 INFO [07-28|15:37:56] 🔗 block reached canonical chain          number=1 hash=a3d5e8…c7bcf6
 INFO [07-28|15:37:56] 🔨 mined potential block                  number=6 hash=596c60…e2c6b4
 INFO [07-28|15:37:56] Commit new mining work                   number=7 txs=0 uncles=0 elapsed=168.144µs
 INFO [07-28|15:37:57] Successfully sealed new block            number=7 hash=15268a…fef71c
 INFO [07-28|15:37:57] 🔗 block reached canonical chain          number=2 hash=86869f…b6b4ff
 INFO [07-28|15:37:57] 🔨 mined potential block                  number=7 hash=15268a…fef71c
 INFO [07-28|15:37:57] Commit new mining work                   number=8 txs=0 uncles=0 elapsed=167.951µs
 > mINFO [07-28|15:37:58] Successfully sealed new block            number=8 hash=d8e9f4…092e4b
 INFO [07-28|15:37:58] 🔗 block reached canonical chain          number=3 hash=89bf98…db5baa
 INFO [07-28|15:37:58] 🔨 mined potential block                  number=8 hash=d8e9f4…092e4b
 INFO [07-28|15:37:58] Commit new mining work                   number=9 txs=0 uncles=0 elapsed=225.274µs
 > minINFO [07-28|15:37:58] Successfully sealed new block            number=9 hash=b85147…eb18d2
 INFO [07-28|15:37:58] 🔗 block reached canonical chain          number=4 hash=0177f9…b36645
 INFO [07-28|15:37:58] 🔨 mined potential block                  number=9 hash=b85147…eb18d2
 INFO [07-28|15:37:58] Commit new mining work                   number=10 txs=0 uncles=0 elapsed=167.023µs
 > minerINFO [07-28|15:37:59] Successfully sealed new block            number=10 hash=fa34a4…626d07
 INFO [07-28|15:37:59] 🔗 block reached canonical chain          number=5  hash=97dbde…e6621d
 INFO [07-28|15:37:59] 🔨 mined potential block                  number=10 hash=fa34a4…626d07
 INFO [07-28|15:37:59] Commit new mining work                   number=11 txs=0 uncles=0 elapsed=140.509µs
 > miner.stINFO [07-28|15:37:59] Successfully sealed new block            number=11 hash=7a06de…5305d9
 INFO [07-28|15:37:59] 🔗 block reached canonical chain          number=6  hash=596c60…e2c6b4
 INFO [07-28|15:37:59] 🔨 mined potential block                  number=11 hash=7a06de…5305d9
 INFO [07-28|15:37:59] Mining too far in the future             wait=2s
 > miner.stop()
 INFO [07-28|15:38:01] Commit new mining work                   number=12 txs=0 uncles=0 elapsed=2.000s
 true

```
 

 