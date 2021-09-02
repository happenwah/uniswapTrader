const { ethers } = require("hardhat");
const hre = require("hardhat");

async function main() {    
    const [deployer] = await ethers.getSigners();
    
    console.log("Deploying with address: ", deployer.address);

    let WETHAddress = "";

    switch (hre.network.name) {
        case "mainnet":
            WETHAddress = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2";
            break;
        case "goerli":
            WETHAddress = "0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6";
            break;
        default:
            throw new Error(
                "network: Flashbots only supported on Ethereum mainnet and goerli"
            );
    }

    if (WETHAddress === "")
        throw new Error(
            "Failed to set constructor arg: WETHAddress"
        );

    const FlashbotsTrader = await hre.ethers.getContractFactory("FlashbotsTrader");
    const flashbotsTrader = await Trader.deploy(WETHAddress);
    
    await flashbotsTrader.deployed();

    console.log("FlashbotsTrader deployed at: ", flashbotsTrader.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.log(error);
        process.exit(1);
    });