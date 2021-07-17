pragma solidity ^0.4.21;

contract HelloWorld {
    string private stateVariable = "Hello, world";

    function GetHelloWorld() public view returns (string) {
        return stateVariable;
    }
}
