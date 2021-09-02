//SPDX-License-Identifier: MIT
pragma solidity =0.8.7;

import { IERC20 } from './interfaces/IERC20.sol';

interface IWETH is IERC20 {
  function deposit() external payable;
  function withdraw(uint) external;
}

/// @notice Flashbots Trader, allows for gas-efficient swaps with 0 slippage
/// @dev Designed specifically for Flashbots integration, DO NOT USE VIA MEMPOOL!
///      Works on UniswapV2 and UniswapV3
///      Can be used as a template for a MEV searcher,
///      or permissioned contract that avoids mempool sandwiches
contract FlashbotsTrader {
  address payable private immutable owner;
  address payable private immutable WETHAddress;

  constructor(address _WETHAddress) 
    payable {
    owner = payable(msg.sender);
    WETHAddress = payable(_WETHAddress);
  }
  
  /// @notice Executes swap between this contract and a target pool
  /// @dev Supports inputToken <-> outputToken swaps
  ///      IMPORTANT: Assumes all data-encoding safety checks have been done off-chain
  /// @param poolAddressWithData - Packed with (blockNumber, fromIsToken0, poolAddress)
  /// @param amounts - Packed with (outputAmount, minerAmountWETH)
  /// @param from - Token to be transfered to the pool
  /// @param inputAmount - Exact amount of token we transfer to poolAddress, assuming no transfer fees are applied
  function executeTrade(
    uint256 poolAddressWithData, 
    uint256 amounts, 
    address from,
    uint256 inputAmount
  )
  external
  payable
  onlyOwner() {
    // Shift 32 - 11 = 21 bytes to the right
    uint256 blockNumber = (poolAddressWithData >> 168);
    // Uncle bandit protection whenever blockNumber > 0
    require(block.number == blockNumber || blockNumber == 0, 'WRONG_BLOCK_NO');
    // Select last 21 bytes + shift 20 bytes to the right
    bool fromIsToken0 = (((poolAddressWithData & 0xffffffffffffffffffffffffffffffffffffffffff) >> 160) == 1) ? true : false;
    // Select last 20 bytes + Pre-prend with 00 (see EIP-55)
    address poolAddress = address(uint160(poolAddressWithData & 0x00ffffffffffffffffffffffffffffffffffffffff));
    // Shift 16 bytes to the right
    uint256 outputAmount = (amounts >> 128);
    // Select last 16 bytes
    uint256 minerAmountWETH = (amounts & 0xffffffffffffffffffffffffffffffff);
    // Transfer inputAmount to poolAddress
    (bool success1,) = from.call(
      abi.encodeWithSelector(0xa9059cbb, poolAddress, inputAmount)
    );
    require(success1, 'TRANSFER_FAILED');
    // Swap
    if (fromIsToken0) {
        (bool success2,) = poolAddress.call(
          abi.encodeWithSelector(0x022c0d9f, 0, outputAmount, address(this), new bytes(0))
        );
        require(success2, 'SWAP_FAILED');
    } else {
        (bool success2,) = poolAddress.call(
          abi.encodeWithSelector(0x022c0d9f, outputAmount, 0, address(this), new bytes(0))
        );
        require(success2, 'SWAP_FAILED');
    }
    // Miner payment
    // If neither is true, we pay tx solely in gas
    if (msg.value > 0) {
      block.coinbase.call{gas: 25000, value: msg.value}("");
    } else if (minerAmountWETH > 0) {
      IWETH(WETHAddress).withdraw(minerAmountWETH);
      block.coinbase.call{gas: 25000, value: minerAmountWETH}("");
    }
  }
  
  /// @notice receive ETH
  receive() external payable {}
  
  /// @dev Fallback must never be called 
  ///      But in case someone tries, 
  ///      we get informed and ensure no external code can be executed
  fallback() external { require(msg.data.length == 0, 'FALLBACK_TRIGGERED'); }

  /// @notice deposits ETH into WETH
  /// @param amount - amount to be deposited
  function convertETHtoWETH(uint256 amount)
    external
    onlyOwner() {
      IWETH(WETHAddress).deposit{value: amount}();
  }
  
  /// @notice Transfers ETH into owner
  /// @dev Withdraws from WETH if necessary
  /// @param amount - maximum amount to withdraw
  function withdrawETH(uint256 amount)
    external
    onlyOwner() {
      uint256 ETHBalance = address(this).balance;
      if (amount > ETHBalance) {
        IWETH WETH = IWETH(WETHAddress);
        uint256 WETHBalance = WETH.balanceOf(address(this));
        if (WETHBalance + ETHBalance >= amount) {
          WETH.withdraw(amount - ETHBalance);
          owner.transfer(amount);
        } else {
          WETH.withdraw(WETHBalance);
          owner.transfer(ETHBalance + WETHBalance);
        }
      } else {
        owner.transfer(amount);
      }
  }
  
  /// @notice Sends token balance back to owner
  /// @param token - address of token to transfer
  function withdrawToken(address token)
    external
    onlyOwner() {
      uint256 balance = IERC20(token).balanceOf(address(this));
      (bool success,) = token.call(abi.encodeWithSelector(0xa9059cbb, owner, balance));
      require(success, 'TRANSFER_FAILED');
  }
  
  /// @notice Destroys contract, sending all its ETH to the owner
  /// @dev IMPORTANT: In case this contract contains other tokens,
  ///      do not forget to call withdrawToken before calling kill
  function kill() 
    external 
    onlyOwner() {
    IWETH WETH = IWETH(WETHAddress);
    uint256 WETHBalance = WETH.balanceOf(address(this));
    if (WETHBalance > 0) WETH.withdraw(WETHBalance);
    selfdestruct(owner);
  }
  
  /// @notice access restricted to owner
  modifier onlyOwner() {
    require(msg.sender == owner, 'ONLY_OWNER');
    _;
  }
}
