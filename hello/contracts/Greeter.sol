pragma solidity ^0.4.0;

contract Greeter
{
    address creator;
    string greeting;

    function Greeter(string _greeting) public
    {
        creator = msg.sender;
        greeting = _greeting;
    }

    function greet() constant returns (string)
    {
        return greeting;
    }

    function setGreeting(string _newgreeting)
    {
        greeting = _newgreeting;
    }

    function kill()
    {
        if (msg.sender == creator)
        suicide(creator);
    }

}