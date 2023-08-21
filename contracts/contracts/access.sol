// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

import {Host, Result} from "@oasisprotocol/sapphire-contracts/contracts/OPL.sol";

contract AccessControl is Host {
    constructor(address _disclosure) Host(_disclosure) {}

    type Access is uint256;
    Access public constant ACCESS_CREDIT_SCORE = Access.wrap(0x01);
    Access public constant ACCESS_FULL_CREDIT_REPORT = Access.wrap(0x02);

    function grantAccess(address _whom, Access _access) external payable {
        postMessage("grantAccess", abi.encode(_msgSender(), _whom, _access));
    }

    function revokeAccess(address _whom) external payable {
        postMessage("revokeAccess", abi.encode(_msgSender(), _whom));
    }
}
