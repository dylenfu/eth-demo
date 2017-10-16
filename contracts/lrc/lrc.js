var _lrcTokenAddress = /* var of type address here */ ;
var _tokenRegistryAddress = /* var of type address here */ ;
var _ringhashRegistryAddress = /* var of type address here */ ;
var _delegateAddress = /* var of type address here */ ;
var _maxRingSize = /* var of type uint256 here */ ;
var _rateRatioCVSThreshold = /* var of type uint256 here */ ;
var browser_untitled1_sol_loopringprotocolimplContract = web3.eth.contract([{"constant":true,"inputs":[{"name":"orderHash","type":"bytes32"}],"name":"getOrderCancelled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"filled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"cancelled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ringIndex","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"addresses","type":"address[3]"},{"name":"orderValues","type":"uint256[7]"},{"name":"buyNoMoreThanAmountB","type":"bool"},{"name":"marginSplitPercentage","type":"uint8"},{"name":"v","type":"uint8"},{"name":"r","type":"bytes32"},{"name":"s","type":"bytes32"}],"name":"cancelOrder","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"RATE_RATIO_SCALE","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"ringminer","type":"address"},{"name":"ringSize","type":"uint256"},{"name":"vList","type":"uint8[]"},{"name":"rList","type":"bytes32[]"},{"name":"sList","type":"bytes32[]"}],"name":"calculateRinghash","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"lrcTokenAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"tokenRegistryAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"addressList","type":"address[2][]"},{"name":"uintArgsList","type":"uint256[7][]"},{"name":"uint8ArgsList","type":"uint8[2][]"},{"name":"buyNoMoreThanAmountBList","type":"bool[]"},{"name":"vList","type":"uint8[]"},{"name":"rList","type":"bytes32[]"},{"name":"sList","type":"bytes32[]"},{"name":"ringminer","type":"address"},{"name":"feeRecepient","type":"address"},{"name":"throwIfLRCIsInsuffcient","type":"bool"}],"name":"submitRing","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"delegateAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"orderHash","type":"bytes32"}],"name":"getOrderFilled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"maxRingSize","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ringhashRegistryAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"cutoffs","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"rateRatioCVSThreshold","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_lrcTokenAddress","type":"address"},{"name":"_tokenRegistryAddress","type":"address"},{"name":"_ringhashRegistryAddress","type":"address"},{"name":"_delegateAddress","type":"address"},{"name":"_maxRingSize","type":"uint256"},{"name":"_rateRatioCVSThreshold","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_ringIndex","type":"uint256"},{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_ringhash","type":"bytes32"},{"indexed":true,"name":"_miner","type":"address"},{"indexed":true,"name":"_feeRecepient","type":"address"},{"indexed":false,"name":"_ringhashFound","type":"bool"}],"name":"RingMined","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_ringIndex","type":"uint256"},{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_ringhash","type":"bytes32"},{"indexed":false,"name":"_prevOrderHash","type":"bytes32"},{"indexed":true,"name":"_orderHash","type":"bytes32"},{"indexed":false,"name":"_nextOrderHash","type":"bytes32"},{"indexed":false,"name":"_amountS","type":"uint256"},{"indexed":false,"name":"_amountB","type":"uint256"},{"indexed":false,"name":"_lrcReward","type":"uint256"},{"indexed":false,"name":"_lrcFee","type":"uint256"}],"name":"OrderFilled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_orderHash","type":"bytes32"},{"indexed":false,"name":"_amountCancelled","type":"uint256"}],"name":"OrderCancelled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_address","type":"address"},{"indexed":false,"name":"_cutoff","type":"uint256"}],"name":"CutoffTimestampChanged","type":"event"}]);
var browser_untitled1_sol_loopringprotocolimpl = browser_untitled1_sol_loopringprotocolimplContract.new(
    _lrcTokenAddress,
    _tokenRegistryAddress,
    _ringhashRegistryAddress,
    _delegateAddress,
    _maxRingSize,
    _rateRatioCVSThreshold,
    {
        from: web3.eth.accounts[0],
        data: '0x606060405260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060045560006005556000600660006101000a81548160ff0219169083151502179055506000600755341561014057600080fd5b60405160c0806119bd833981016040528080519060200190919080519060200190919080519060200190919080519060200190919080519060200190919080519060200190919050505b856000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600481905550806007819055505b5050505050505b61170b806102b26000396000f300606060405236156100e4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806312e0c430146100ec578063288cdc91146101275780632ac126221461016257806341ffbc1f1461019d57806347a99e43146101c65780634a63864b146102755780634beb7e7c1461029e5780634c0a6532146103bc5780635be2aca01461041157806364c86dda146104665780636d96a2aa1461073d578063798591bb14610792578063cbc12d13146107cd578063d21b96ab146107f6578063de794c1e1461084b578063df565ca214610898575b5b600080fd5b005b34156100f757600080fd5b6101116004808035600019169060200190919050506108c1565b6040518082815260200191505060405180910390f35b341561013257600080fd5b61014c6004808035600019169060200190919050506108e7565b6040518082815260200191505060405180910390f35b341561016d57600080fd5b6101876004808035600019169060200190919050506108ff565b6040518082815260200191505060405180910390f35b34156101a857600080fd5b6101b0610917565b6040518082815260200191505060405180910390f35b34156101d157600080fd5b61027360048080606001906003806020026040519081016040528092919082600360200280828437820191505050505091908060e001906007806020026040519081016040528092919082600760200280828437820191505050505091908035151590602001909190803560ff1690602001909190803560ff16906020019091908035600019169060200190919080356000191690602001909190505061091d565b005b341561028057600080fd5b610288610b01565b6040518082815260200191505060405180910390f35b34156102a957600080fd5b61039e600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091905050610b07565b60405180826000191660001916815260200191505060405180910390f35b34156103c757600080fd5b6103cf610c08565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561041c57600080fd5b610424610c2d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561047157600080fd5b61073b60048080359060200190820180359060200190808060200260200160405190810160405280939291908181526020016000905b828210156104ec5784848390506040020160028060200260405190810160405280929190826002602002808284378201915050505050815260200190600101906104a7565b5050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020016000905b8282101561056857848483905060e002016007806020026040519081016040528092919082600760200280828437820191505050505081526020019060010190610523565b5050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020016000905b828210156105e457848483905060400201600280602002604051908101604052809291908260026020028082843782019150505050508152602001906001019061059f565b5050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080351515906020019091905050610c53565b005b341561074857600080fd5b610750610d9b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561079d57600080fd5b6107b7600480803560001916906020019091905050610dc1565b6040518082815260200191505060405180910390f35b34156107d857600080fd5b6107e0610de7565b6040518082815260200191505060405180910390f35b341561080157600080fd5b610809610ded565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561085657600080fd5b610882600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610e13565b6040518082815260200191505060405180910390f35b34156108a357600080fd5b6108ab610e2b565b6040518082815260200191505060405180910390f35b60006009600083600019166000191681526020019081526020016000205490505b919050565b60086020528060005260406000206000915090505481565b60096020528060005260406000206000915090505481565b60055481565b6000610927611450565b600088600660078110151561093857fe5b602002015192506101c0604051908101604052808b600060038110151561095b57fe5b602002015173ffffffffffffffffffffffffffffffffffffffff1681526020018b600160038110151561098a57fe5b602002015173ffffffffffffffffffffffffffffffffffffffff1681526020018b60026003811015156109b957fe5b602002015173ffffffffffffffffffffffffffffffffffffffff1681526020018a60006007811015156109e857fe5b602002015181526020018a6001600781101515610a0157fe5b602002015181526020018a6002600781101515610a1a57fe5b602002015181526020018a6003600781101515610a3357fe5b602002015181526020018a6004600781101515610a4c57fe5b602002015181526020018a6005600781101515610a6557fe5b6020020151815260200189151581526020018860ff1681526020018760ff1681526020018660001916815260200185600019168152509150610aa682610e31565b905080600019167fdb6b2087bc479399a384085a4ff023f244b32682c168d31c93b82194caa7be8842438660405180848152602001838152602001828152602001935050505060405180910390a25b50505050505050505050565b61271081565b600085848484604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401848051906020019060200280838360005b83811015610b7f5780820151818401525b602081019050610b63565b50505050905001838051906020019060200280838360005b83811015610bb35780820151818401525b602081019050610b97565b50505050905001828051906020019060200280838360005b83811015610be75780820151818401525b602081019050610bcb565b50505050905001945050505050604051809103902090505b95945050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080610c5e611510565b610c66611524565b8d519350610c7787858c8c8c610b07565b9250610c89848f8f8f8f8f8f8f611012565b915060a060405190810160405280846000191681526020018381526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018615158152509050806060015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff168260000151600019167f4108ba28a2df317b9d4c4a2a42b33aa556bfebc994f9b4ba275fe337bf589898600560008154809291906001019190505542436001604051808581526020018481526020018381526020018215151515815260200194505050505060405180910390a45b5050505050505050505050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006008600083600019166000191681526020019081526020016000205490505b919050565b60045481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600a6020528060005260406000206000915090505481565b60075481565b600030826000015183602001518460400151856060015186608001518760a001518860c001518960e001518a61010001518b61012001518c6101400151604051808d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140189815260200188815260200187815260200186815260200185815260200184815260200183151515157f01000000000000000000000000000000000000000000000000000000000000000281526001018260ff1660ff167f01000000000000000000000000000000000000000000000000000000000000000281526001019c50505050505050505050505050604051809103902090505b919050565b61101a611510565b611022611510565b60008061102d611450565b60008d60405180591061103d5750595b90808252806020026020018201604052801561107357816020015b61106061158b565b8152602001906001900390816110585790505b509450600093505b8d84101561143b578d6001850181151561109157fe5b0692506101c0604051908101604052808e868151811015156110af57fe5b9060200190602002015160006002811015156110c757fe5b602002015173ffffffffffffffffffffffffffffffffffffffff1681526020018e868151811015156110f557fe5b90602001906020020151600160028110151561110d57fe5b602002015173ffffffffffffffffffffffffffffffffffffffff1681526020018e8581518110151561113b57fe5b90602001906020020151600160028110151561115357fe5b602002015173ffffffffffffffffffffffffffffffffffffffff1681526020018d8681518110151561118157fe5b90602001906020020151600060078110151561119957fe5b602002015181526020018d868151811015156111b157fe5b9060200190602002015160016007811015156111c957fe5b602002015181526020018d868151811015156111e157fe5b9060200190602002015160026007811015156111f957fe5b602002015181526020018d8681518110151561121157fe5b90602001906020020151600360078110151561122957fe5b602002015181526020018d8681518110151561124157fe5b90602001906020020151600460078110151561125957fe5b602002015181526020018d8681518110151561127157fe5b90602001906020020151600560078110151561128957fe5b602002015181526020018b868151811015156112a157fe5b90602001906020020151151581526020018c868151811015156112c057fe5b9060200190602002015160006002811015156112d857fe5b602002015160ff1681526020018a868151811015156112f357fe5b9060200190602002015160ff168152602001898681518110151561131357fe5b90602001906020020151600019168152602001888681518110151561133457fe5b9060200190602002015160001916815250915061135082610e31565b905061014060405190810160405280838152602001826000191681526020018c8681518110151561137d57fe5b90602001906020020151600160028110151561139557fe5b602002015160ff16815260200160408051908101604052808f888151811015156113bb57fe5b9060200190602002015160066007811015156113d357fe5b602002015181526020018560800151815250815260200160008152602001600081526020016000815260200160008152602001600081526020016000815250858581518110151561142057fe5b906020019060200201819052505b838060010194505061107b565b8495505b505050505098975050505050505050565b6101c060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600015158152602001600060ff168152602001600060ff16815260200160008019168152602001600080191681525090565b602060405190810160405280600081525090565b60a060405190810160405280600080191681526020016115426115f1565b8152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581525090565b610300604051908101604052806115a0611605565b815260200160008019168152602001600060ff1681526020016115c16116c5565b81526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b602060405190810160405280600081525090565b6101c060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600015158152602001600060ff168152602001600060ff16815260200160008019168152602001600080191681525090565b6040805190810160405280600081526020016000815250905600a165627a7a723058201f8b05b3eadf4199f563df521e6414aecc762c7edf2586a9624530b60e7226640029',
        gas: '4700000'
    }, function (e, contract){
        console.log(e, contract);
        if (typeof contract.address !== 'undefined') {
            console.log('Contract mined! address: ' + contract.address + ' transactionHash: ' + contract.transactionHash);
        }
    })