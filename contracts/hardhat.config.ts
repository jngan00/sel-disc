import { HardhatUserConfig } from "hardhat/config";
import '@oasisprotocol/sapphire-hardhat';
import "@nomicfoundation/hardhat-toolbox";

const accounts = process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [];

task('deploy-sapphire')
  .addParam('hostNetwork')
  .addParam('writer')
  .setAction(async (args, hre) => {
    await hre.run('compile');
    const ethers = hre.ethers;
    const CreditData = await ethers.getContractFactory('CreditData');
    const Disclosure = await ethers.getContractFactory('Disclosure');
    const Sbt = await ethers.getContractFactory('CreditSBTv1');

    const signer = ethers.provider.getSigner();
    const signerAddr = await signer.getAddress();

    // Start by predicting the address of the Access contract.
    const hostConfig = hre.config.networks[args.hostNetwork];
    if (!('url' in hostConfig)) throw new Error(`${args.hostNetwork} not configured`);
    const provider = new ethers.providers.JsonRpcProvider(hostConfig.url);
    let nonce = await provider.getTransactionCount(signerAddr);
    if (args.hostNetwork === 'local')
      // 3 contract deployments, 3 variable sets
      nonce += 6;

    const hostAddr = ethers.utils.getContractAddress({ from: signerAddr, nonce });
    console.log('expected host address', hostAddr);

    const creditData = await CreditData.deploy();
    const disclosure = await Disclosure.deploy(creditData.address, hostAddr);
    const sbt = await Sbt.deploy(creditData.address);

    await creditData.deployed();
    await disclosure.deployed();
    await sbt.deployed();

    // Reference disclosure and sbt contracts
    await creditData.setDisclosure(disclosure.address);
    await creditData.setSbt(sbt.address);
    await creditData.setWriter(args.writer);

    console.log('credit data address', creditData.address);
    console.log('disclosure address', disclosure.address);
    console.log('sbt address', sbt.address);

    return creditData.address;
  });

task('deploy-access')
  .addParam('disclosureAddr')
  .setAction(async (args, hre) => {
    await hre.run('compile');
    const ethers = hre.ethers;
    const Access = await ethers.getContractFactory('AccessControl');

    const access = await Access.deploy(args.disclosureAddr);
    await access.deployed();

    console.log('access', access.address);
  });

task('deploy-local')
  .setAction(async (args, hre) => {
    const sap = await hre.run('deploy-sapphire', {
      hostNetwork: 'local',
      // hardhat Account #0
      writer: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    });
    await hre.run('deploy-access', { disclosureAddr: sap });
  });

const config: HardhatUserConfig = {
  solidity: "0.8.17",
  networks: {
    'local': {
      url: 'http://127.0.0.1:8545',
    },
    'sapphire_mainnet': {
      url: 'https://sapphire.oasis.io',
      chainId: 0x5afe,
      accounts,
    },
    'sapphire_testnet': {
      url: 'https://testnet.sapphire.oasis.dev',
      chainId: 0x5aff,
      accounts,
    },
    'bsc_testnet': {
      url: 'https://data-seed-prebsc-1-s1.binance.org:8545',
      chainId: 97,
      accounts,
    },
    'polygon_testnet': {
      url: 'https://polygon-mumbai-bor.publicnode.com',
      chainId: 80001,
      accounts,
    }
  },
};

export default config;
