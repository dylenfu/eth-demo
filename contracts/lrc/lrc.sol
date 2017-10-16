pragma solidity ^0.4.15;

// 仅用于lrc/miner调试

contract LoopringProtocolImpl {


    ////////////////////////////////////////////////////////////////////////////
    /// Variables                                                            ///
    ////////////////////////////////////////////////////////////////////////////

    address public  lrcTokenAddress             = address(0);
    address public  tokenRegistryAddress        = address(0);
    address public  ringhashRegistryAddress     = address(0);
    address public  delegateAddress             = address(0);

    uint    public  maxRingSize                 = 0;
    uint    public  ringIndex                   = 0;
    bool    private entered                     = false;

    // Exchange rate (rate) is the amount to sell or sold divided by the amount
    // to buy or bought.
    //
    // Rate ratio is the ratio between executed rate and an order's original
    // rate.
    //
    // To require all orders' rate ratios to have coefficient ofvariation (CV)
    // smaller than 2.5%, for an example , rateRatioCVSThreshold should be:
    //     `(0.025 * RATE_RATIO_SCALE)^2` or 62500.
    uint    public  rateRatioCVSThreshold       = 0;

    uint    public constant RATE_RATIO_SCALE    = 10000;

    // The following two maps are used to keep trace of order fill and
    // cancellation history.
    mapping (bytes32 => uint) public filled;
    mapping (bytes32 => uint) public cancelled;

    // A map from address to its cutoff timestamp.
    mapping (address => uint) public cutoffs;


    ////////////////////////////////////////////////////////////////////////////
    /// Structs                                                              ///
    ////////////////////////////////////////////////////////////////////////////

    struct Rate {
    uint amountS;
    uint amountB;
    }

    struct Order {
    address owner;
    address tokenS;
    address tokenB;
    uint    amountS;
    uint    amountB;
    uint    timestamp;
    uint    ttl;
    uint    salt;
    uint    lrcFee;
    bool    buyNoMoreThanAmountB;
    uint8   marginSplitPercentage;
    uint8   v;
    bytes32 r;
    bytes32 s;
    }


    /// @param order        The original order
    /// @param orderHash    The order's hash
    /// @param feeSelection -
    ///                     A miner-supplied value indicating if LRC (value = 0)
    ///                     or margin split is choosen by the miner (value = 1).
    ///                     We may support more fee model in the future.
    /// @param rate         Exchange rate provided by miner.
    /// @param availableAmountS -
    ///                     The actual spendable amountS.
    /// @param fillAmountS  Amount of tokenS to sell, calculated by protocol.
    /// @param lrcReward    The amount of LRC paid by miner to order owner in
    ///                     exchange for margin split.
    /// @param lrcFee       The amount of LR paid by order owner to miner.
    /// @param splitS      TokenS paid to miner.
    /// @param splitB      TokenB paid to miner.
    struct OrderState {
    Order   order;
    bytes32 orderHash;
    uint8   feeSelection;
    Rate    rate;
    uint    availableAmountS;
    uint    fillAmountS;
    uint    lrcReward;
    uint    lrcFee;
    uint    splitS;
    uint    splitB;
    }

    struct Ring {
    bytes32      ringhash;
    OrderState[] orders;
    address      miner;
    address      feeRecepient;
    bool         throwIfLRCIsInsuffcient;
    }


    ////////////////////////////////////////////////////////////////////////////
    /// Events                                                               ///
    ////////////////////////////////////////////////////////////////////////////

    event RingMined(
    uint                _ringIndex,
    uint                _time,
    uint                _blocknumber,
    bytes32     indexed _ringhash,
    address     indexed _miner,
    address     indexed _feeRecepient,
    bool                _ringhashFound);

    event OrderFilled(
    uint                _ringIndex,
    uint                _time,
    uint                _blocknumber,
    bytes32     indexed _ringhash,
    bytes32             _prevOrderHash,
    bytes32     indexed _orderHash,
    bytes32              _nextOrderHash,
    uint                _amountS,
    uint                _amountB,
    uint                _lrcReward,
    uint                _lrcFee);

    event OrderCancelled(
    uint                _time,
    uint                _blocknumber,
    bytes32     indexed _orderHash,
    uint                _amountCancelled);

    event CutoffTimestampChanged(
    uint                _time,
    uint                _blocknumber,
    address     indexed _address,
    uint                _cutoff);


    ////////////////////////////////////////////////////////////////////////////
    /// Constructor                                                          ///
    ////////////////////////////////////////////////////////////////////////////

    function LoopringProtocolImpl(
    address _lrcTokenAddress,
    address _tokenRegistryAddress,
    address _ringhashRegistryAddress,
    address _delegateAddress,
    uint    _maxRingSize,
    uint    _rateRatioCVSThreshold
    )
    public {
        lrcTokenAddress             = _lrcTokenAddress;
        tokenRegistryAddress        = _tokenRegistryAddress;
        ringhashRegistryAddress     = _ringhashRegistryAddress;
        delegateAddress             = _delegateAddress;
        maxRingSize                 = _maxRingSize;
        rateRatioCVSThreshold       = _rateRatioCVSThreshold;
    }


    ////////////////////////////////////////////////////////////////////////////
    /// Public Functions                                                     ///
    ////////////////////////////////////////////////////////////////////////////

    /// @dev Disable default function.
    function () payable {
        revert();
    }

    /// @dev Submit a order-ring for validation and settlement.
    /// @param addressList  List of each order's tokenS. Note that next order's
    ///                     `tokenS` equals this order's `tokenB`.
    /// @param uintArgsList List of uint-type arguments in this order:
    ///                     amountS, amountB, timestamp, ttl, salt, lrcFee,
    ///                     rateAmountS.
    /// @param uint8ArgsList -
    ///                     List of unit8-type arguments, in this order:
    ///                     marginSplitPercentageList,feeSelectionList.
    /// @param vList        List of v for each order. This list is 1-larger than
    ///                     the previous lists, with the last element being the
    ///                     v value of the ring signature.
    /// @param rList        List of r for each order. This list is 1-larger than
    ///                     the previous lists, with the last element being the
    ///                     r value of the ring signature.
    /// @param sList        List of s for each order. This list is 1-larger than
    ///                     the previous lists, with the last element being the
    ///                     s value of the ring signature.
    /// @param ringminer    The address that signed this tx.
    /// @param feeRecepient The recepient address for fee collection. If this is
    ///                     '0x0', all fees will be paid to the address who had
    ///                     signed this transaction, not `msg.sender`. Noted if
    ///                     LRC need to be paid back to order owner as the result
    ///                     of fee selection model, LRC will also be sent from
    ///                     this address.
    /// @param throwIfLRCIsInsuffcient -
    ///                     If true, throw exception if any order's spendable
    ///                     LRC amount is smaller than requried; if false, ring-
    ///                     minor will give up collection the LRC fee.
    function submitRing(
    address[2][]    addressList,
    uint[7][]       uintArgsList,
    uint8[2][]      uint8ArgsList,
    bool[]          buyNoMoreThanAmountBList,
    uint8[]         vList,
    bytes32[]       rList,
    bytes32[]       sList,
    address         ringminer,
    address         feeRecepient,
    bool            throwIfLRCIsInsuffcient
    )
    public {
        uint ringSize = addressList.length;

        bytes32 ringhash = calculateRinghash(
        ringminer,
        ringSize,
        vList,
        rList,
        sList
        );

        // Assemble input data into a struct so we can pass it to functions.
        var orders = assembleOrders(
        ringSize,
        addressList,
        uintArgsList,
        uint8ArgsList,
        buyNoMoreThanAmountBList,
        vList,
        rList,
        sList);

        var ring = Ring(
        ringhash,
        orders,
        ringminer,
        feeRecepient,
        throwIfLRCIsInsuffcient);

        RingMined(
        ringIndex++,
        block.timestamp,
        block.number,
        ring.ringhash,
        ring.miner,
        ring.feeRecepient,
        true
        );
    }

    function calculateRinghash(
    address     ringminer,
    uint        ringSize,
    uint8[]     vList,
    bytes32[]   rList,
    bytes32[]   sList
    )
    public
    constant
    returns (bytes32) {

        return keccak256(
        ringminer,
        vList,//vList.xorReduce(ringSize),
        rList,//rList.xorReduce(ringSize),
        sList);//sList.xorReduce(ringSize));
    }

    /// @dev Cancel a order. Amount (amountS or amountB) to cancel can be
    ///                           specified using orderValues.
    /// @param addresses          owner, tokenS, tokenB
    /// @param orderValues        amountS, amountB, timestamp, ttl, salt,
    ///                           lrcFee, and cancelAmount
    /// @param buyNoMoreThanAmountB -
    ///                           If true, this order does not accept buying
    ///                           more than `amountB`.
    /// @param marginSplitPercentage -
    ///                           The percentage of margin paid to miner.
    /// @param v                  Order ECDSA signature parameter v.
    /// @param r                  Order ECDSA signature parameters r.
    /// @param s                  Order ECDSA signature parameters s.
    function cancelOrder(
    address[3] addresses,
    uint[7]    orderValues,
    bool       buyNoMoreThanAmountB,
    uint8      marginSplitPercentage,
    uint8      v,
    bytes32    r,
    bytes32    s
    )
    public {

        uint cancelAmount = orderValues[6];

        var order = Order(
        addresses[0],
        addresses[1],
        addresses[2],
        orderValues[0],
        orderValues[1],
        orderValues[2],
        orderValues[3],
        orderValues[4],
        orderValues[5],
        buyNoMoreThanAmountB,
        marginSplitPercentage,
        v,
        r,
        s
        );

        bytes32 orderHash = calculateOrderHash(order);

        OrderCancelled(
        block.timestamp,
        block.number,
        orderHash,
        cancelAmount
        );
    }

    function getOrderFilled(bytes32 orderHash)
    public
    constant
    returns (uint) {
        return filled[orderHash];
    }

    function getOrderCancelled(bytes32 orderHash)
    public
    constant
    returns (uint) {
        return cancelled[orderHash];
    }


    ////////////////////////////////////////////////////////////////////////////
    /// Internal & Private Functions                                         ///
    ////////////////////////////////////////////////////////////////////////////

    /// @dev        assmble order parameters into Order struct.
    /// @return     A list of orders.
    function assembleOrders(
    uint            ringSize,
    address[2][]    addressList,
    uint[7][]       uintArgsList,
    uint8[2][]      uint8ArgsList,
    bool[]          buyNoMoreThanAmountBList,
    uint8[]         vList,
    bytes32[]       rList,
    bytes32[]       sList
    )
    internal
    constant
    returns (OrderState[]) {

        var orders = new OrderState[](ringSize);

        for (uint i = 0; i < ringSize; i++) {
            uint j = (i + 1) % ringSize;

            var order = Order(
            addressList[i][0],
            addressList[i][1],
            addressList[j][1],
            uintArgsList[i][0],
            uintArgsList[i][1],
            uintArgsList[i][2],
            uintArgsList[i][3],
            uintArgsList[i][4],
            uintArgsList[i][5],
            buyNoMoreThanAmountBList[i],
            uint8ArgsList[i][0],
            vList[i],
            rList[i],
            sList[i]);

            bytes32 orderHash = calculateOrderHash(order);

            orders[i] = OrderState(
            order,
            orderHash,
            uint8ArgsList[i][1],  // feeSelection
            Rate(uintArgsList[i][6], order.amountB),
            0,   // availableAmountS
            0,   // fillAmountS
            0,   // lrcReward
            0,   // lrcFee
            0,   // splitS
            0    // splitB
            );

        }

        return orders;
    }

    /// @dev Get the Keccak-256 hash of order with specified parameters.
    function calculateOrderHash(Order order)
    internal
    constant
    returns (bytes32) {

        return keccak256(
        address(this),
        order.owner,
        order.tokenS,
        order.tokenB,
        order.amountS,
        order.amountB,
        order.timestamp,
        order.ttl,
        order.salt,
        order.lrcFee,
        order.buyNoMoreThanAmountB,
        order.marginSplitPercentage);
    }
}