
## 解锁时间
. 使用personal.unlockAccount解锁账户,解锁时间是有限定的<br>
```cmd
> personal.listWallets
[{
    accounts: [{
        address: "0xa3a601635f5f51392c2045d5f5617bde98622b13",
        url: "keystore:///home/vagrant/gohome/src/eth-demo/build/seed0/data/keystore/UTC--2017-07-28T08-09-12.079263977Z--a3a601635f5f51392c2045d5f5617bde98622b13"
    }],
    status: "Locked",
    url: "keystore:///home/vagrant/gohome/src/eth-demo/build/seed0/data/keystore/UTC--2017-07-28T08-09-12.079263977Z--a3a601635f5f51392c2045d5f5617bde98622b13"
}]
> personal.unlockAccount("0xa3a601635f5f51392c2045d5f5617bde98622b13", "1", 864000)
true
> 
> personal.listWallets
[{
    accounts: [{
        address: "0xa3a601635f5f51392c2045d5f5617bde98622b13",
        url: "keystore:///home/vagrant/gohome/src/eth-demo/build/seed0/data/keystore/UTC--2017-07-28T08-09-12.079263977Z--a3a601635f5f51392c2045d5f5617bde98622b13"
    }],
    status: "Unlocked",
    url: "keystore:///home/vagrant/gohome/src/eth-demo/build/seed0/data/keystore/UTC--2017-07-28T08-09-12.079263977Z--a3a601635f5f51392c2045d5f5617bde98622b13"
}]
```

. 被锁定的账户仍然能挖矿,只是不能交易<br>
继续挖矿后查看账户<br>
这里注意eth即web3.eth
```cmd
> eth.getBalance(eth.accounts[0])
335000000000000000000
> web3.eth.getBalance("0xa3a601635f5f51392c2045d5f5617bde98622b13")
335000000000000000000
```
