import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from 'chai';
import { ethers } from 'hardhat';

import { CreditData } from '..';

const { loadFixture } = require("@nomicfoundation/hardhat-network-helpers");

const utils = ethers.utils;

describe('CreditData', function () {
    async function deployContracts() {
        const CreditData = await ethers.getContractFactory('CreditData');
        const Disclosure = await ethers.getContractFactory('Disclosure');
        const Sbt = await ethers.getContractFactory('CreditSBTv1');
        const Access = await ethers.getContractFactory('AccessControl');

        const block = await ethers.provider.getBlock("latest");
        // 3 contract deployments, 3 variable sets
        let nonce = block.nonce + 6;

        const signer = ethers.provider.getSigner();
        const signerAddr = await signer.getAddress();
        const hostAddr = ethers.utils.getContractAddress({ from: signerAddr, nonce });

        const creditData = await CreditData.deploy();
        await creditData.deployed();

        const disclosure = await Disclosure.deploy(creditData.address, hostAddr);
        const sbt = await Sbt.deploy(creditData.address);
        await sbt.deployed();

        await creditData.setDisclosure(disclosure.address);
        await creditData.setSbt(sbt.address);

        const [owner, writer, user1, user2, user3] = await ethers.getSigners();
        await creditData.setWriter(writer.address);

        const access = await Access.deploy(disclosure.address);
        await access.deployed();

        return { creditData, disclosure, sbt, access, owner, writer, user1, user2, user3 };
    }

    it('Deploy', async function () {
        const { creditData, disclosure, sbt, access, owner } = await loadFixture(deployContracts);

        expect(await creditData.owner()).to.equal(owner.address);
    });

    it('RequestCreditScoreOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);
        await expect(await creditData.connect(user1).requestCreditScore(user1.address))
            .to.emit(creditData, 'RequestScore')
            .withArgs(user1.address);
    });

    it('RequestCombinedCreditScoreOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);
        await expect(await creditData.connect(user1).requestCombinedCreditScore([user1.address, user2.address]))
            .to.emit(creditData, 'RequestCombinedScore')
            .withArgs([user1.address, user2.address]);
    });

    it('RequestCreditReportOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);
        await expect(await creditData.connect(user1).requestCreditReport(user1.address))
            .to.emit(creditData, 'RequestReport')
            .withArgs(user1.address);
    });

    it('WriterStoreCreditScoreOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);
        await expect(await creditData.connect(writer).storeCreditScore(user1.address, 600, 0))
            .to.emit(creditData, 'ScoreReady')
            .withArgs(user1.address);
    });

    it('NotWriterStoreCreditScoreNotOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);
        await expect(creditData.connect(user1).storeCreditScore(user1.address, 600, 0))
            .to.be.revertedWith('invalid sender');
    });

    it('WriterStoreCombinedCreditScoreOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);
        await expect(await creditData.connect(writer).storeCombinedCreditScore([user1.address, user2.address], 600, 0))
            .to.emit(creditData, 'CombinedScoreReady')
            .withArgs([user1.address, user2.address]);
    });

    it('NotWriterStoreCombinedCreditScoreNotOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);
        await expect(creditData.connect(user1).storeCombinedCreditScore([user1.address, user2.address], 600, 0))
            .to.be.revertedWith('invalid sender');
    });

    it('WriterStoreCreditReportOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);
        await expect(await creditData.connect(writer).storeCreditReport(user1.address, utils.formatBytes32String('600'), 0))
            .to.emit(creditData, 'ReportReady')
            .withArgs(user1.address);
    });

    it('NotWriterStoreCreditReportNotOk', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);
        await expect(creditData.connect(user1).storeCreditReport(user1.address, utils.formatBytes32String('600'), 0))
            .to.be.revertedWith('invalid sender');
    });

    it('CannotAccessScoreWithoutPermission', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);

        await expect(disclosure.connect(user1).accessCreditScore(user2.address))
            .to.be.revertedWith('permission denied');
        await expect(disclosure.connect(user1).accessFullCreditReport(user2.address))
            .to.be.revertedWith('permission denied');
    });

    it('CanAccessScoreWithPermission', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);

        await creditData.connect(writer).storeCreditScore(user2.address, 600, 2**32 - 1);
        await disclosure.connect(user2).grantAccess(user1.address, 0x01);
        await disclosure.connect(user1).accessCreditScore(user2.address);

        await creditData.connect(writer).storeCreditReport(user2.address, utils.formatBytes32String('600'), 2**32 - 1);
        await disclosure.connect(user2).grantAccess(user1.address, 0x02);
        await disclosure.connect(user1).accessFullCreditReport(user2.address);
    });

    it('CannotAccessScoreAfterRevocation', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);

        await creditData.connect(writer).storeCreditScore(user2.address, 600, 2**32 - 1);
        await disclosure.connect(user2).grantAccess(user1.address, 0x01);
        await disclosure.connect(user1).accessCreditScore(user2.address);

        await creditData.connect(writer).storeCreditReport(user2.address, utils.formatBytes32String('600'), 2**32 - 1);
        await disclosure.connect(user2).grantAccess(user1.address, 0x02);
        await disclosure.connect(user1).accessFullCreditReport(user2.address);

        await disclosure.connect(user2).revokeAccess(user1.address);
        await expect(disclosure.connect(user1).accessCreditScore(user2.address))
            .to.be.revertedWith('permission denied');
        await expect(disclosure.connect(user1).accessFullCreditReport(user2.address))
            .to.be.revertedWith('permission denied');
    });

    it('CanMintSbt', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);

        await creditData.connect(writer).storeCreditScore(user1.address, 600, 2**32 - 1);
        await expect(await sbt.connect(user1).safeMint(user1.address))
            .to.emit(sbt, 'Minted')
            .withArgs(user1.address, 0);
        await sbt.hasValid(user1.address);
    });

    it('CanRevokeSbt', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);

        await creditData.connect(writer).storeCreditScore(user1.address, 600, 2**32 - 1);
        await sbt.connect(user1).safeMint(user1.address);
        await expect(await sbt.revoke(user1.address, 0))
            .to.emit(sbt, 'Revoked')
            .withArgs(user1.address, 0);
    });

    it('CanSelfBurnSbt', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1 } = await loadFixture(deployContracts);

        await creditData.connect(writer).storeCreditScore(user1.address, 600, 2**32 - 1);
        await sbt.connect(user1).safeMint(user1.address);
        await sbt.connect(user1).selfBurn(0);
    });

    it('AccessGrantGivesPermission', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);

        await creditData.connect(writer).storeCreditScore(user2.address, 600, 2**32 - 1);
        await access.connect(user2).grantAccess(user1.address, 0x01);
        expect(await disclosure.connect(user1).hasAccess(user2.address, 0x01)).to.equal(true);
        await disclosure.connect(user1).accessCreditScore(user2.address);

        await creditData.connect(writer).storeCreditReport(user2.address, utils.formatBytes32String('600'), 2**32 - 1);
        await access.connect(user2).grantAccess(user1.address, 0x02);
        await disclosure.connect(user1).accessFullCreditReport(user2.address);
    });

    it('AccessRevokeRemovesPermission', async function () {
        const { creditData, disclosure, sbt, access, owner, writer, user1, user2 } = await loadFixture(deployContracts);

        await creditData.connect(writer).storeCreditScore(user2.address, 600, 2**32 - 1);
        await access.connect(user2).grantAccess(user1.address, 0x01);
        await disclosure.connect(user1).accessCreditScore(user2.address);

        await creditData.connect(writer).storeCreditReport(user2.address, utils.formatBytes32String('600'), 2**32 - 1);
        await access.connect(user2).grantAccess(user1.address, 0x02);
        await disclosure.connect(user1).accessFullCreditReport(user2.address);

        await access.connect(user2).revokeAccess(user1.address);
        await expect(disclosure.connect(user1).accessCreditScore(user2.address))
            .to.be.revertedWith('permission denied');
        await expect(disclosure.connect(user1).accessFullCreditReport(user2.address))
            .to.be.revertedWith('permission denied');
    });
});
