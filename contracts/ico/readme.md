使用在线编辑器编辑solidity代码：<br>
https://ethereum.github.io/browser-solidity <br>
编译后再contract/web3 deploy中拷贝js代码,复制并放入到ico项目token.js中<br>

合约合服erc20标准，就是说部署上之后代币即为erc20代币.<br>
合约中的构造函数，需要在编译合约后获得js文件，在js文件中填充两个地址.<br>
js文件中填充的是params/chain.go中account1&account2<br>

对同一个合约文件同时部署两次就可以得到两种代币A，B<br>

合约A地址
0x359bbea6ade5155bce1e95918879903d3e93365f
合约B地址
0xc85819398e4043f3d951367d6d97bb3257b862e0