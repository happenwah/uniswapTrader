require("@nomiclabs/hardhat-waffle");
require("dotenv").config();
const ethers = require("ethers");

const PK_MAINNET = process.env.PK_MAINNET;
const PK_GOERLI = process.env.PK_GOERLI;

const GAS_PRICE_MAINNET = parseInt(ethers.utils.parseUnits("100", "gwei").toString());
const GAS_PRICE_GOERLI = parseInt(ethers.utils.parseUnits("15", "gwei").toString());

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: {
    compilers: [
      {
        version: "0.8.7",
        settings: {
          optimizer: { enabled: true },
        },
      },
    ],
  },
  networks: {
    mainnet: {
      url: process.env.PROVIDER_MAINNET,
      chainId: 1,
      accounts: PK_MAINNET ? [PK_MAINNET] : [],
      gasPrice: GAS_PRICE_MAINNET,
    },
    goerli: {
      url: process.env.PROVIDER_GOERLI,
      chainId: 5,
      accounts: PK_GOERLI ? [PK_GOERLI] : [],
      gasPrice: GAS_PRICE_GOERLI,
    },
  },
};
