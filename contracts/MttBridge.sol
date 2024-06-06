pragma solidity ^0.8.20;

import {IToken} from "./interface/IToken.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

contract MttBridge is Ownable {
    bool public isPause;
    address public minter;
    address public poster;

    mapping(bytes32 => bool) public isMint;
    mapping(bytes32 => bool) public posted;
    mapping(bytes32 => Log) public logs;
    mapping(string => bool) public chains;

    struct Log {
        address account;
        uint256 value;
    }

    event MintEvent(bytes32 indexed hash);
    event PostEvent(
        bytes32 indexed hash,
        address indexed account,
        uint256 value
    );
    event BridgeEvent(address indexed account, uint256 value, string chain);

    modifier onlyMinter() {
        require(msg.sender == minter, "caller is not minter");
        _;
    }

    modifier onlyPoster() {
        require(msg.sender == poster, "caller is not poster");
        _;
    }

    constructor() Ownable(msg.sender) {}

    function setChain(string calldata name, bool status) external onlyOwner {
        chains[name] = status;
    }

    function setMinter(address _minter) external onlyOwner {
        minter = _minter;
    }

    function setPoster(address _poster) external onlyOwner {
        poster = _poster;
    }

    function pauseBridge(bool _val) external onlyOwner {
        isPause = _val;
    }

    function withdraw(address payable to, uint256 value) external onlyOwner {
        bool sent = to.send(value);
        require(sent, "Failed to send MTT");
    }

    function bridge(uint256 value, string calldata chain) external payable {
        require(!isPause, "bridge Paused");
        require(chains[chain], "chain not support");
        require(msg.value == value, "value not equal");
        emit BridgeEvent(msg.sender, value, chain);
    }

    function mint(bytes32 txHash) external onlyMinter {
        require(posted[txHash], "tx not posted");
        require(isMint[txHash] == false, "dupilicate hash");
        emit MintEvent(txHash);
    }

    function post(
        bytes32 txHash,
        address account,
        uint256 value
    ) external onlyPoster {
        posted[txHash] = true;
        Log storage log = logs[txHash];
        log.account = account;
        log.value = value;
        emit PostEvent(txHash, account, value);
    }
}
