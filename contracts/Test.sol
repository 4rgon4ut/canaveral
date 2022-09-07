// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract Test {
    address public owner;
    uint16 public number;
    string public word;
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

    function getFields()
        external
        view
        returns (
            uint16,
            string memory,
            bytes32,
            address
        )
    {
        return (number, word, obj, addr);
    }

    function setNumber(uint16 _number) external {
        number = _number;
    }

    function setWord(string memory _word) external {
        word = _word;
    }

    function setObj(bytes32 _obj) external {
        obj = _obj;
    }

    function setAddr(address _addr) external {
        addr = _addr;
    }
}
