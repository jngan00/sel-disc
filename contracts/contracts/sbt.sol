// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ERC721 } from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import { ERC721Burnable } from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import { Counters } from "@openzeppelin/contracts/utils/Counters.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

import { IERC4671 } from "./interfaces/IERC4671.sol";
import { IERC5192 } from "./interfaces/IERC5192.sol";
import { CreditData } from "./credit.sol";

contract CreditSBTv1 is ERC721, ERC721Burnable, Ownable, IERC4671, IERC5192 {
    CreditData public creditEndpoint;

    using Counters for Counters.Counter;
    Counters.Counter private _tokenIdCounter;

    struct TokenData {
        uint16 score;
        uint expiry;
    }

    // From tokenId to score.
    mapping(uint256 => TokenData) private _tokenData;

    constructor(address endpoint) ERC721("CreditSBT", "CreditSBTv1") {
        creditEndpoint = CreditData(endpoint);
    }

    function safeMint(address to) public returns (uint256) {
        require(_msgSender() == to, "only by user");
        require(balanceOf(to) == 0, "already has SBT");

        TokenData memory data;
        (data.score, data.expiry, ) = creditEndpoint.getCreditScore(to);

        require(data.expiry > block.timestamp, "score already expired");

        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
        _tokenData[tokenId] = data;

        emit Minted(to, tokenId);
	return tokenId;
    }

    function revoke(address owner, uint256 tokenId) public onlyOwner {
        require(ownerOf(tokenId) == owner, "token not owned by owner");
        _burn(tokenId);

        emit Revoked(owner, tokenId);
    }

    // Can only be called by token owner.
    function selfBurn(uint256 tokenId) public {
        require(ownerOf(tokenId) == _msgSender(), "token not owned by sender");
        _burn(tokenId);
        delete _tokenData[tokenId];
    }

    // Can be called by anyone but token must already be expired.
    function removeExpired(uint256 tokenId) public {
        require(_tokenData[tokenId].expiry < block.timestamp, "not expired");
        _burn(tokenId);
        delete _tokenData[tokenId];
    }

    function scoreOf(uint256 tokenId) external view returns (uint16) {
        TokenData memory data = _tokenData[tokenId];
        require(data.score > 0, "token not found");
        require(data.expiry > block.timestamp, "credit score expired");
        return data.score;
    }

    /// @dev See {IERC5192-locked}.
    function locked(uint256) public pure override(IERC5192) returns (bool) {
        return true;
    }

    /// @dev See {IERC4671-isValid}.
    function isValid(uint256 tokenId) external view override(IERC4671) returns (bool) {
        return _exists(tokenId);
    }

    /// @dev See {IERC4671-hasValid}.
    function hasValid(address owner) external view override(IERC4671) returns (bool) {
        return balanceOf(owner) > 0;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, IERC165)
        returns (bool)
    {
        return
            super.supportsInterface(interfaceId) ||
            interfaceId == type(IERC4671).interfaceId ||
            interfaceId == type(IERC5192).interfaceId;
    }

    // The following functions are overrides required by Solidity.

    function ownerOf(uint256 tokenId) public view override(ERC721, IERC4671) returns (address) {
        return super.ownerOf(tokenId);
    }

    function balanceOf(address owner) public view override(ERC721, IERC4671) returns (uint256) {
        return super.balanceOf(owner);
    }
}
