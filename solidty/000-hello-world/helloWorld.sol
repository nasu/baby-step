pragma solidity ^0.4.24;
contract HelloWorld {
    string public text;
    constructor() public { text = "Hello World!"; }
    function say() view public returns (string) { return text; }
}
