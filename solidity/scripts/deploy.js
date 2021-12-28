async function main() {
    const [deployer] = await ethers.getSigners();
  
    console.log("Deploying contracts with the account:", deployer.address);
  
    console.log("Account balance:", (await deployer.getBalance()).toString());
  
    const PicassoReferalContract = await ethers.getContractFactory("PicassoReferal");
    const picassoReferal = await PicassoReferalContract.deploy();
    console.log("Picasso Referal contract address:", picassoReferal.address);
    
  }
  
  main()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });