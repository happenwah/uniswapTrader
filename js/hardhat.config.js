require("@nomiclabs/hardhat-waffle");
require("dotenv").config();

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: "0.8.7",
  networks: {
    mainnet: {
      url: process.env.PROVIDER_MAINNET,
      chainId: 1,
      accounts: {
        mnemonic: process.env.MNEMONIC
      },
    },
    goerli: {
      url: process.env.PROVIDER_GOERLI,
      chainId: 5,
      accounts: {
        mnemonic: process.env.MNEMONIC
      },
    },
  },
};
