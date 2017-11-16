pragma solidity ^0.4.15;

// 用于复合结构event解析

contract RingMinedEvt {

    struct OrderFilled {
        bytes32 _orderHash;
        bytes32 _nextOrderHash;
        uint    _amountS;
        uint    _amountB;
        uint    _lrcReward;
        uint    _lrcFee;
    }

    event RingMined(
        uint                _ringIndex,
        bytes32     indexed _ringhash,
        address     indexed _miner,
        address     indexed _feeRecipient,
        bool                _isRinghashReserved,
        OrderFilled[]       _fills
    );

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

        RingMined(
        _ringIndex++,
        _ringhash,
        _miner,
        _feeRecipient,
        _isRinghashReserved,
        list
        );
    }

}