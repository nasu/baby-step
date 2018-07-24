pragma solidity ^0.4.15;
contract OriginalToken {
    // Token amount which an Address has.
    mapping (address => uint256) balances;
    // Token amount to be able to transfer from former Address to latter Address.
    mapping (address => mapping (address => uint256)) allowed;
    //TODO: What's "indexed" ?
    event TransferEvent(address indexed _from, address indexed _to, uint256 _amount);
    event ApprovalEvent(address indexed _owner, address indexed _spender, uint256 indexed _amount);
    modifier ValidAmount(uint256 amount) {
        require (0 < amount);
        _;
    }
    modifier ValidBalance(address from, uint256 amount) {
        require (amount <= balances[from]);
        _;
    }
    modifier ValidAllowed(address from, address to, uint256 amount) {
        require (amount <= allowed[from][to]);
        _;
    }

    //TODO: Check kinds of data types.
    constructor(uint256 _supply) public {
        balances[msg.sender] = _supply;
    }

    // Transfer _amount from me to _to.
    function Transfer(address _to, uint256 _amount)
        ValidAmount(_amount) ValidBalance(msg.sender, _amount)
        public returns (bool _success)
    {
        balances[msg.sender] -= _amount;
        balances[_to] += _amount;
        emit TransferEvent(msg.sender, _to, _amount);
        return true;
    }

    function TransferFrom(address _from, address _to, uint256 _amount)
        ValidAmount(_amount) ValidBalance(_from, _amount) ValidAllowed(_from, _to, _amount)
        public returns (bool success)
    {
        balances[_from] -= _amount;
        balances[_to] += _amount;
        allowed[_from][_to] += _amount;
        emit TransferEvent(_from, _to, _amount);
        return true;
    }

    function Approve(address _spender, uint256 _amount)
        public returns (bool success)
    {
        allowed[msg.sender][_spender] = _amount;
        emit ApprovalEvent(msg.sender, _spender, _amount);
        return true;
    }

    function BalanceOf(address _addr)
        view public returns (uint256 _balance)
    {
        return (balances[_addr]);
    }

    function Allowance(address _owner, address _spender)
        view public returns (uint256 remaining)
    {
        return allowed[_owner][_spender];
    }
}