pragma solidity ^0.4.0;

contract Bank
{
    enum DepositStatus { PartialFinished, FullFinished }

    mapping(bytes32 => OrderState)  orderStates;
    mapping(address => uint)  accounts;

    struct Deposit {
        bytes32          hash;
        address         account;
        uint            amount;
    }

    struct Order {
        bytes32         hash;
        address         accountS;
        address         accountB;
        uint            amountS;
        uint            amountB;
    }

    struct OrderState {
        Order           order;
        bool            ok;
    }

    ///////////////////////////////////////////////////////////////////
    //
    // events
    //
    ///////////////////////////////////////////////////////////////////
    event DepositFilled(
        bytes32         hash,
        address         account,
        uint            amount,
        bool            ok
    );

    event OrderFilled(
        bytes32         hash,
        address         accountS,
        address         accountB,
        uint            amountS,
        uint            amountB,
        bool            ok
    );

    event Exception(string message);

    ///////////////////////////////////////////////////////////////////
    //
    // external functions
    //
    ///////////////////////////////////////////////////////////////////
    function submitTransfer(bytes32 hash, address accountS, address accountB, uint amountS, uint amountB) {
        check(balanceOf(accountB) > amountB, "not enough money");

        accounts[accountB] -= amountB;
        accounts[accountS] += amountS;

        var ord = Order(
            hash,
            accountS,
            accountB,
            amountS,
            amountB
        );

        var state = OrderState(
            ord,
            true
        );

        orderStates[hash] = state;
        OrderFilled(hash, accountS, accountB, amountS, amountB, true);
    }

    function submitDeposit(bytes32 _id, address _owner, uint _amount) {
        accounts[_owner] += _amount;
        DepositFilled(_id, _owner, _amount, true);
    }

    function balanceOf(address _owner) constant returns (uint balance) {
        return accounts[_owner];
    }

    function check(bool condition, string message) {
        if (!condition) {
            Exception(message);
            revert();
        }
    }

}