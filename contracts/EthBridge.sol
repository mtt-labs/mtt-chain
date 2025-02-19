pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {IERC20} from "@openzeppelin/contracts/interfaces/IERC20.sol";

contract EthBridge is Ownable {
    bool public isPause;
    uint256 public numConfirmationsRequired;
    uint256 public minterNum;
    uint256 public fee;

    mapping(bytes32 => bool) public executed;
    mapping(string => bool) public chains;
    mapping(bytes32 => log) public logs;

    mapping(address => bool) public isMinter;
    mapping(bytes32 => mapping(address => bool)) public isConfirmed;

    address public token;

    struct log {
        bytes32 txHash;
        address addr;
        uint256 amount;
        uint8 approveNum;
    }

    event MintEvent(
        bytes32 indexed hash,
        address indexed account,
        uint256 value
    );

    event SetChainEvent(string indexed chain, bool status);
    event SetMinterEvent(address indexed account, bool status);
    event SetFeeEvent(uint256 fee);
    event SetTokenEvent(address token);
    event SetConfirmationsRequiredEvent(uint256 confirmationsRequired);
    event PauseEvent(bool status);
    event ConfirmEvent(bytes32 indexed hash,bytes32 txHash, address indexed account);
    event BridgeEvent(address indexed account, uint256 value, string chain);

    constructor(
        address initialOwner,
        address[] memory _minters,
        uint256 _numConfirmationsRequired,
        address _token
    ) Ownable(initialOwner) {
        require(_minters.length > 0, "minters required");
        require(
            _numConfirmationsRequired > 0 &&
                _numConfirmationsRequired <= _minters.length,
            "invalid number of required confirmations"
        );

        for (uint256 i = 0; i < _minters.length; i++) {
            address minter = _minters[i];

            require(minter != address(0), "invalid owner");
            require(!isMinter[minter], "owner not unique");

            isMinter[minter] = true;
        }

        minterNum = _minters.length;
        numConfirmationsRequired = _numConfirmationsRequired;
        token = _token;
    }

    function setChain(string calldata name, bool status) external onlyOwner {
        chains[name] = status;
        emit SetChainEvent(name, status);
    }

    function setMinter(address addr, bool value) external onlyOwner {
        if (value) {
            if (!isMinter[addr]) {
                minterNum = minterNum + 1;
            }
        } else {
            if (isMinter[addr]) {
                minterNum = minterNum - 1;
            }
        }
        isMinter[addr] = value;
        require(minterNum >= numConfirmationsRequired, "minter not enough");
        emit SetMinterEvent(addr, value);
    }

    function setFee(uint256 _fee) external onlyOwner {
        fee = _fee;
        emit SetFeeEvent(_fee);
    }

    function setToken(address _token) external onlyOwner {
        token = _token;
        emit SetTokenEvent(_token);
    }

    function pauseBridge(bool _val) external onlyOwner {
        isPause = _val;
        emit PauseEvent(_val);
    }

    function setConfirmationsRequired(
        uint256 _confirmationsRequired
    ) external onlyOwner {
        numConfirmationsRequired = _confirmationsRequired;
        require(numConfirmationsRequired>0, "invalid number of required confirmations");
        require(minterNum >= numConfirmationsRequired, "minter not enough");
        emit SetConfirmationsRequiredEvent(_confirmationsRequired);
    }

    function withdraw(address to, uint256 value) external onlyOwner {
        require(IERC20(token).transfer(to, value), "transfer fail");
    }

    function bridgeWithPermit(
        uint256 value,
        string calldata chain,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external {
        require(value>fee, "fee not enough");
        require(!isPause, "bridge Paused");
        require(chains[chain], "chain not support");
        IERC20Permit(token).permit(
            msg.sender,
            address(this),
            value,
            deadline,
            v,
            r,
            s
        );
        require(
            IERC20(token).transferFrom(msg.sender, address(this), value),
            "token transfer from fail"
        );
        emit BridgeEvent(msg.sender, value, chain);
    }

    function bridge(uint256 value, string calldata chain) external {
        require(value>fee, "fee not enough");
        require(!isPause, "bridge Paused");
        require(chains[chain], "chain not support");
        require(
            IERC20(token).transferFrom(msg.sender, address(this), value),
            "token transfer from fail"
        );
        emit BridgeEvent(msg.sender, value, chain);
    }

    modifier onlyMinter() {
        require(isMinter[msg.sender], "not minter");
        _;
    }

    modifier notExecuted(bytes32 txHash) {
        require(!executed[txHash], "tx already executed");
        _;
    }

    function mint(
        bytes32 txHash,
        address addr,
        uint256 amount
    ) external onlyMinter notExecuted(txHash) {
        bytes32 hash = keccak256(abi.encodePacked(txHash, addr, amount));
        require(!isConfirmed[hash][msg.sender], "tx already confirmed");
        log storage txLog = logs[hash];

        if (txLog.addr != address(0)) {
            txLog.approveNum = txLog.approveNum + 1;
        } else {
            txLog.addr = addr;
            txLog.amount = amount;
            txLog.approveNum = 1;
            txLog.txHash = txHash;
        }
        isConfirmed[hash][msg.sender] = true;

        emit ConfirmEvent(hash,txHash, msg.sender);

        if (txLog.approveNum >= numConfirmationsRequired) {
            if (txLog.amount > fee) {
                amount = txLog.amount - fee;
                require(IERC20(token).transfer(txLog.addr, amount), "transfer fail");
                emit MintEvent(txHash, txLog.addr, amount);
            } else {
                emit MintEvent(txHash, txLog.addr, 0);
            }
            executed[txHash] = true;
        }
    }
}
