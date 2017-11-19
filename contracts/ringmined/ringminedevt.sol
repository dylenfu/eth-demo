pragma solidity ^0.4.15;

// 用于复合结构event解析

contract RingMinedEvt {

    struct OrderFilled {
        bytes32 orderHash;
        bytes32 nextOrderHash;
        uint    amountS;
        uint    amountB;
        uint    lrcReward;
        uint    lrcFee;
    }

    event RingMinedEvent(
        uint                ringIndex,
        bytes32             ringhash,
        address             miner,
        address             feeRecipient,
        bool                isRinghashReserved,
        OrderFilled[]       fills
    );

    event RingEvent (
        uint                ringIndex,
        bytes32             ringhash,
        address             miner,
        address             feeRecipient,
        bool                isRinghashReserved
    );

    event MinEvent(
        address             miner,
        uint[]              amounts
    );

    struct SimpleFill {
        address owner;
        uint    amount;
    }

    event SimpleRingEvent(
        address protocol,
        uint    res,
        SimpleFill[] fills
    );

    function simpleRing(
        address protocol,
        address miner,
        uint    amount,
        uint    res
    ) public {

        var list = new SimpleFill[](2);
        var fill = SimpleFill(
            miner,
            amount
        );
        var fillSec = SimpleFill(
            miner,
            amount++
        );
        list[0] = fill;
        list[1] = fillSec;

        SimpleRingEvent(
            protocol,
            res,
            list
        );
    }

    function submitRing(
        uint    _ringIndex,
        bytes32 _ringhash,
        address _miner,
        address _feeRecipient,
        bool    _isRinghashReserved,
        bytes32 _orderHash,
        uint    _amount,
        uint    _lrcReward,
        uint    _lrcFee
    )
    public {

        var order = OrderFilled(
        _orderHash,
        _orderHash,
        _amount,
        _amount,
        _lrcReward,
        _lrcFee
        );

        var list = new OrderFilled[](2);
        list[0] = order;
        list[1] = order;

        RingMinedEvent(
        _ringIndex++,
        _ringhash,
        _miner,
        _feeRecipient,
        _isRinghashReserved,
        list
        );
    }

    function justRing(
    uint    _ringIndex,
    bytes32 _ringhash,
    address _miner,
    address _feeRecipient,
    bool    _isRinghashReserved
    )
    public {
        RingEvent(
        _ringIndex++,
        _ringhash,
        _miner,
        _feeRecipient,
        _isRinghashReserved
        );
    }

    function min(address miner, uint number) public {
        var list = new uint[](2);
        list[0] = number++;
        list[1] = number++;

        MinEvent(
            miner,
            list
        );
    }

}