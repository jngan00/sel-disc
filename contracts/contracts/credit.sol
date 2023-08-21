// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

// A repository of credit scores and credit reports for each wallet.
contract CreditData is Ownable {
    // The address of the contract to provide disclosure. External access to any credit data is
    // possible only through this contract.
    address public _disclosure;

    // The only user with permission to directly write credit data.
    address public _writer;

    // The address of the contract to create SBT.
    address public _sbt;

    event RequestScore(address whom);
    event RequestCombinedScore(address[] whom);
    event RequestReport(address whom);

    event ScoreReady(address whom);
    event CombinedScoreReady(address[] whom);
    event ReportReady(address whom);

    // Per-wallet data.
    struct WalletData {
        uint16 score;
	uint scoreExpiry;
        // If the score is a combined score for multiple linked addresses, this is the list of all
        // addresses used.
        address[] linked;

	string report;
	uint reportExpiry;
    }

    mapping(address => WalletData) private _data;

    function setDisclosure(address whom) external onlyOwner {
        _disclosure = whom;
    }

    function setWriter(address whom) external onlyOwner {
        _writer = whom;
    }

    function setSbt(address whom) external onlyOwner {
        _sbt = whom;
    }

    modifier fromWalletOwner(address whom) {
	require(_msgSender() == whom, "invalid requester");
	_;
    }

    modifier fromWriter() {
        require(_msgSender() == _writer, "invalid sender");
        _;
    }

    // Request a credit score for a single address.
    function requestCreditScore(address whom) external fromWalletOwner(whom) {
	emit RequestScore(whom);
    }

    // Request a combined credit score from multiple addresses. The resulted score is assigned to only
    // the first address.
    function requestCombinedCreditScore(address[] calldata whom) external fromWalletOwner(whom[0]) {
        require(whom.length > 1, "need multiple addresses");

	emit RequestCombinedScore(whom);
    }

    // Request a credit report for a single address.
    function requestCreditReport(address whom) external fromWalletOwner(whom) {
	emit RequestReport(whom);
    }

    // Write a credit score for a single address.
    function storeCreditScore(address whom, uint16 score, uint expiry) external fromWriter() {
	_data[whom].score = score;
	_data[whom].scoreExpiry = expiry;
	delete _data[whom].linked;

	emit ScoreReady(whom);
    }

    // Write a credit score for a multiple addresses.
    function storeCombinedCreditScore(address[] calldata whom, uint16 score, uint expiry) external fromWriter() {
	_data[whom[0]].score = score;
	_data[whom[0]].scoreExpiry = expiry;
	_data[whom[0]].linked = whom;

	emit CombinedScoreReady(whom);
    }

    // Write a credit report for a single address.
    function storeCreditReport(address whom, string calldata report, uint expiry) external fromWriter() {
	_data[whom].report = report;
	_data[whom].reportExpiry = expiry;

	emit ReportReady(whom);
    }

    function clearCreditScore() external {
	_data[_msgSender()].score = 0;
    }

    function clearCreditReport() external {
        _data[_msgSender()].report = '';
    }

    // Below are access functions for disclosure contract to get credit information.

    // Returns (credit score, expiry, list of wallet addresses that generated this score)
    function getCreditScore(address whom) public view returns (uint16, uint, address[] memory) {
        require(_msgSender() == _disclosure || _msgSender() == _sbt,
                "no permission");
        WalletData memory data = _data[whom];
        require(data.score > 0, "no credit score available");
        require(data.scoreExpiry > block.timestamp, "credit score has expired");
        return (data.score, data.scoreExpiry, data.linked);
    }

    // Returns (report, expiry)
    function getCreditReport(address whom) public view returns (string memory, uint) {
        require(_msgSender() == _disclosure, "no permission");
        WalletData memory data = _data[whom];
        require(bytes(data.report).length > 0, "no credit report available");
        require(data.reportExpiry > block.timestamp, "credit report has expired");
        return (data.report, data.reportExpiry);
    }
}
