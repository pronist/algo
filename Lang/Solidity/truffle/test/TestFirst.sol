pragma solidity ^0.4.21;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/First.sol";

contract TestFirst {
    function testGetDoublePositiveUsingDeployedContract() public {
        First meta = First(DeployedAddresses.First());
        Assert.equal(meta.GetDouble(10), 20, "Positive input gives double value");
    }
}
