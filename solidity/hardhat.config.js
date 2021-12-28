require("@nomiclabs/hardhat-waffle");

/**
 * @type import('hardhat/config').HardhatUserConfig
 */

const PRIVATE_KEY = "48c2341e77218d5a094d50b2b7ee49bd3e4a189fda134cec6501d5f222e037ca";

module.exports = {
  solidity: "0.8.2",
  networks: {
    ropsten: {
      url: `https://ropsten.infura.io/v3/d1f944a7e6b8418caaa1c504a17aa7bb`,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    matic: {
      url: `https://rpc-mumbai.maticvigil.com`,
      accounts: [`0x${PRIVATE_KEY}`]
    }
  }
};
