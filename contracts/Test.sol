// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract Test {
    address public owner;

    string public word;
    uint16 public number;
    bytes32 public obj;
    address public addr;

    constructor(
        uint16 _number,
        string memory _word,
        bytes32 _obj,
        address _addr
    ) {
        owner = msg.sender;
        number = _number;
        word = _word;
        obj = _obj;
        addr = _addr;
    }
}
