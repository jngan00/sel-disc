// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

import {Enclave, Result, autoswitch} from "@oasisprotocol/sapphire-contracts/contracts/OPL.sol";

import { CreditData } from "./credit.sol";

contract Disclosure is Enclave {
    CreditData public immutable creditEndpoint;

    constructor(address endpoint, address remote) Enclave(remote, autoswitch("polygon")) {
        creditEndpoint = CreditData(endpoint);
	registerEndpoint("grantAccess", _oplGrantAccess);
	registerEndpoint("revokeAccess", _oplRevokeAccess);
    }

    type Access is uint256;
    Access public constant ACCESS_CREDIT_SCORE = Access.wrap(0x01);
    Access public constant ACCESS_FULL_CREDIT_REPORT = Access.wrap(0x02);
    mapping(address => mapping(address => Access)) private access; // granter => grantee => access

    modifier hasPermission(address _whom, Access _access) {
        require(hasAccess(_whom, _access), "permission denied");
        _;
    }

    function hasAccess(address _whom, Access _accessType) public view returns (bool)
    {
        return
            _whom == _msgSender() ||
            Access.unwrap(access[_whom][_msgSender()]) & Access.unwrap(_accessType) ==
                Access.unwrap(_accessType);
    }

    function grantAccess(address _whom, Access _accessType) public {
	require(_whom != _msgSender(), "no need to grant access to self");

        access[_msgSender()][_whom] = Access.wrap(
            Access.unwrap(access[_msgSender()][_whom]) | Access.unwrap(_accessType)
        );
    }

    function _oplGrantAccess(bytes calldata args) internal returns (Result) {
        (address _from, address _whom, Access _accessType) = abi.decode(
            args,
            (address, address, Access)
        );

	require(_whom != _from, "no need to grant access to self");
        access[_from][_whom] = Access.wrap(
            Access.unwrap(access[_from][_whom]) | Access.unwrap(_accessType)
        );
        return Result.Success;
    }

    function revokeAccess(address _whom) public {
        require(_whom != _msgSender(), "cannot revoke access to self");

        access[_msgSender()][_whom] = Access.wrap(0);
    }

    function _oplRevokeAccess(bytes calldata args) internal returns (Result) {
        (address _from, address _whom) = abi.decode(args, (address, address));

        require(_whom != _from, "cannot revoke access to self");
        access[_from][_whom] = Access.wrap(0);

        return Result.Success;
    }

    // Returns (score, expiry, list of wallet addresses that generated this score).
    function accessCreditScore(address _whom)
        external
        view
        hasPermission(_whom, ACCESS_CREDIT_SCORE)
        returns (uint16, uint, address[] memory)
    {
        return creditEndpoint.getCreditScore(_whom);
    }

    // Returns (report, expiry).
    function accessFullCreditReport(address _whom)
        external
        view
        hasPermission(_whom, ACCESS_FULL_CREDIT_REPORT)
        returns (string memory, uint)
    {
        return creditEndpoint.getCreditReport(_whom);
    }
}
