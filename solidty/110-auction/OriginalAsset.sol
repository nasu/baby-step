pragma solidity ^0.4.15;
contract OriginalAsset {
    address owner;
    string public name;
    mapping (address => bool) allowed;

    event TransferEvent(string _name, address _from, address _to);
    modifier isOwner() {
        require (owner == msg.sender);
        _; //TODO: What's this?
    }

    constructor(string _name) public {
        owner = msg.sender;
        name = _name;
    }

    function TransferAsset(address _to)
        isOwner
        public returns (bool _success)
    {
        emit TransferEvent(name, owner, _to);
        owner = _to;
        return true;
    }

    function TransferAssetApproved(address _to)
        public returns (bool _success)
    {
        if (allowed[_to] != true) return false;
        allowed[_to] = false;
        emit TransferEvent(name, owner, _to);
        owner = _to;
        return true;
    }

    function Approve(address _to)
        isOwner
        public returns (bool _success)
    {
        allowed[_to] = true;
        return true;
    }

    function Allowance(address _to)
        view public returns (bool _allowed)
    {
        return allowed[_to];
    }

    function GetOwner()
        view public returns (address _owner)
    {
        return owner;
    }

    function GetName()
        view public returns (string _name)
    {
        return name;
    }
}