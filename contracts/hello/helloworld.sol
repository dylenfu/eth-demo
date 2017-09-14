pragma solidity ^0.4.0;

contract HelloWorld
{
    address creator;
    string greeting;

    function HelloWorld() public
    {
        creator = msg.sender;
    }

    function greet() constant returns (string)
    {
        return "hello solidity";
    }

    function setGreeting(string _newgreeting)
    {
        greeting = _newgreeting;
    }

    function getGreeting() constant returns (string)
    {
        return greeting;
    }

    function kill()
    {
        if (msg.sender == creator)
            suicide(creator);
    }
}