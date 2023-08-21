<script setup lang="ts">

import { ref } from 'vue'
import { ethers } from 'ethers'
import * as sapphire from '@oasisprotocol/sapphire-paratime';

import CreditData from '../artifacts/CreditData.json'
import CreditSBTv1 from '../artifacts/CreditSBTv1.json'
import Disclosure from '../artifacts/Disclosure.json'

// contract addresses on Sapphire testnet
const creditDataAdd = "0x2a275b858C413b06051FF5B6f67d7d8B675f8c7B"
const disclosureAdd = "0xdccF4F49fA818Bc3Fd7B7bfb21CE7A5C123dcffd"
const creditSBTv1Add = "0xAf06A8BfABae4A56b0aE994FCAbcd883C137529E"


const network = ref("st")
const userType = ref(0)

var address = ref("connect to show your address")
var targetAddress = ref("")
var result = ref("nothing to see yet")
var tokenId = ref("")
var creditValue = ref("")

function handleAddress(add: any) {
  address.value = add[0];
}

function user() {
  userType.value = 0;
}

function verifier() {
  userType.value = 1;
}

function writer() {
  userType.value = 2;
}

const switchAccount = async () => {
  const chainIdMap = {
    'sm': ['0x5afe', 'Sapphire Mainnet', 'https://sapphire.oasis.io'],
    'st': ['0x5aff', 'Sapphire Testnet', 'https://testnet.sapphire.oasis.dev'],
    'lh': ['0x5afd', 'Local Hardhat', 'http://127.0.0.1:8545'],
  }
  let chain = chainIdMap[network.value as keyof typeof chainIdMap]
  console.log(chain)
  try {
    await window.ethereum.request({
      method: 'wallet_switchEthereumChain',
      params: [{ chainId: chain[0] }],
    })
  } catch (switchError) {
    // This error code indicates that the chain has not been added to MetaMask.
    if (switchError.code === 4902) {
      try {
        await window.ethereum.request({
          method: 'wallet_addEthereumChain',
          params: [
            {
              chainId: chain[0],
              chainName: chain[1],
              rpcUrls: [chain[2]] /* ... */,
            },
          ],
        })
      } catch (addError) {
        result.value = addError as string
      }
    } else {
        result.value = switchError.message
    }
  }
}

const connect = async () => {
  await window.ethereum
    .request({ method: 'eth_requestAccounts' })
    .then(handleAddress);
  await switchAccount()
};

function updateErr(err: any) {
  if (err instanceof Error) {
    console.log(err);
    result.value = err.message;
  }
  else if (typeof err == 'object') {
    result.value = 'Failed. ' + err
  } else {
    console.log(err)
    result.value = err
  }
}

const getScore = async (target?: string) => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const disc = new ethers.Contract(disclosureAdd, Disclosure.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    if (typeof target === 'undefined') {
      target = address.value
    }
    console.log('target is', target)
    const data = await disc.accessCreditScore(target)
    console.log(data)
    result.value = "Credit score is " + data[0]
  } catch (err: any) {
    result.value = "Cannot retrieve credit score"
  }
};

const getReport = async (target?: string) => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const disc = new ethers.Contract(disclosureAdd, Disclosure.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    if (typeof target === 'undefined') {
      target = address.value
    }
    console.log('target is', target)
    const data = await disc.accessFullCreditReport(target)
    console.log(data)
    result.value = "Credit report: " + data[0]
  } catch (err) {
    result.value = "Cannot retrieve credit report"
  }
}

const issueScore = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const cd = new ethers.Contract(creditDataAdd, CreditData.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await cd.requestCreditScore(address.value)
    console.log("transaction: " + transaction)
    const data = await transaction.wait()
    console.log("data: " + data)
    result.value = "Credit score generated"
  } catch (err) {
    updateErr(err)
  }
}

const issueReport = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const cd = new ethers.Contract(creditDataAdd, CreditData.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await cd.requestCreditReport(address.value)
    const data = await transaction.wait()
    result.value = "Credit report generated"
  } catch (err) {
    updateErr(err)
  }
}

const clearScore = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const cd = new ethers.Contract(creditDataAdd, CreditData.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await cd.clearCreditScore()
    console.log("transaction: " + transaction)
    const data = await transaction.wait()
    console.log("data: " + data)
    result.value = "Credit score cleared"
  } catch (err) {
    updateErr(err)
  }
}

const clearReport = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const cd = new ethers.Contract(creditDataAdd, CreditData.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await cd.clearCreditReport()
    console.log("transaction: " + transaction)
    const data = await transaction.wait()
    console.log("data: " + data)
    result.value = "Credit score cleared"
  } catch (err) {
    updateErr(err)
  }
}

const mintSBT = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const sbtv1 = new ethers.Contract(creditSBTv1Add, CreditSBTv1.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    sbtv1.on("Minted", (to, token) => {
      result.value = "SBT minted; token id: " + token;
    });
    const transaction = await sbtv1.safeMint(address.value)
    const data = await transaction.wait()
    //result.value = "SBT minted"
  } catch (err) {
    updateErr(err)
  }
}

const burnSBT = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const sbtv1 = new ethers.Contract(creditSBTv1Add, CreditSBTv1.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await sbtv1.selfBurn(tokenId.value)
    const data = await transaction.wait()
    result.value = "SBT burnt"
  } catch (err) {
    updateErr(err)
  }
}

const hasSBT = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const sbtv1 = new ethers.Contract(creditSBTv1Add, CreditSBTv1.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const data = await sbtv1.hasValid(address.value)
    result.value = data
  } catch (err) {
    updateErr(err)
  }
}

const approveScore = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const disc = new ethers.Contract(disclosureAdd, Disclosure.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await disc.grantAccess(targetAddress.value, 1)
    const data = await transaction.wait()
    console.log(transaction)
    console.log(data)
    result.value = "access granted"
  } catch (err) {
    updateErr(err)
  }
}

const approveReport = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const disc = new ethers.Contract(disclosureAdd, Disclosure.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await disc.grantAccess(targetAddress.value, 2)
    const data = await transaction.wait()
    console.log(transaction)
    console.log(data)
    result.value = "access granted"
  } catch (err) {
    updateErr(err)
  }
}
const revoke = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const disc = new ethers.Contract(disclosureAdd, Disclosure.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const transaction = await disc.revokeAccess(targetAddress.value)
    const data = await transaction.wait()
    console.log(data)
    result.value = "access revoked"
  } catch (err) {
    updateErr(err)
  }
}

const check = async (access: number) => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const disc = new ethers.Contract(disclosureAdd, Disclosure.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    console.log("checking " + targetAddress.value + " by " + address.value)
    const data = await disc.hasAccess(targetAddress.value, access)
    result.value = data
  } catch (err) {
    updateErr(err)
  }
}

function checkTarget() {
  if (targetAddress.value.length != 42) {
    window.alert("put input valid target address first")
    return false
  }
  return true
}

const verifyScore = async () => {
  checkTarget() && getScore(targetAddress.value)
}

const verifyReport = async () => {
  checkTarget() && getReport(targetAddress.value)
}

const writeScore = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const cd = new ethers.Contract(creditDataAdd, CreditData.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const data = await cd.storeCreditScore(targetAddress.value, Number(creditValue.value), Date.now())
    result.value = "wrote ok"
  } catch (err) {
    result.value = "Cannot write credit score"
  }
}

const writeReport = async () => {
  await connect()
  result.value = "Please wait..."
  const provider = sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum))
  const cd = new ethers.Contract(creditDataAdd, CreditData.abi,
    sapphire.wrap(new ethers.providers.Web3Provider(window.ethereum).getSigner()))
  try {
    const data = await cd.storeCreditReport(targetAddress.value, creditValue.value, Date.now())
    result.value = "wrote ok"
  } catch (err) {
    result.value = "Cannot write credit report"
  }
}
</script>

<template>
  <div style="font-size:32pt"><img width="80" src="/oasis.svg"><br>Credit-on-Sapphire Demo</div>
  <p/>

  <div style="margin-top: 20px;">
    <div class="banner">
        Command result:<br/>
        <span class="result">{{ result }}</span>
    </div>
  </div>

  <table style="width:100%;text-align:center;margin:20px;">
    <tr>
      <td colspan="3">
        <button type="button" @click="connect()" style="width:100%">Connnect</button>
      </td>
    </tr>
    <tr>
      <td colspan="3">
        <input class="show" readonly placeholder="connect to show your address" v-model="address"/>
      </td>
    </tr>

    <tr>
      <td width="34%">
        <button type="button" @click="user()" style="width:100%" :class="{ passive: userType == 0 }">User</button>
      </td>
      <td width="33%">
        <button type="button" @click="verifier()" style="width:100%" :class="{ passive: userType == 1 }">Verifier</button>
      </td>
      <td width="33%">
        <button type="button" @click="writer()" style="width:100%" :class="{ passive: userType == 2 }">Writer</button>
      </td>
    </tr>
  </table>

  <div v-if="userType == 0">
    <h2>User Menu</h2>
    <div>
      Credit score
      <button type="button" @click="getScore()">Get</button>
      <button type="button" @click="issueScore()">Issue</button>
      <button type="button" @click="clearScore()">Clear</button>
    </div>
    <div>
      Credit report
      <button type="button" @click="getReport()">Get</button>
      <button type="button" @click="issueReport()">Issue</button>
      <button type="button" @click="clearReport()">Clear</button>
    </div>
    <div>
      SBT
      <button type="button" @click="mintSBT()">Mint</button>
      <button type="button" @click="hasSBT()">Check SBT</button>
      <br>
      <button type="button" @click="burnSBT()">Burn</button>
      <input v-model="tokenId" placeholder="burn token id" />
    </div>
    <div style="margin-top: 12px;">
      <h3>Selective disclosure approval:</h3>
      <input class="target" v-model="targetAddress" placeholder="target address" /> <br/>
      <button type="button" @click="checkTarget() && approveScore()">Approve score</button>
      <button type="button" @click="checkTarget() && approveReport()">Approve report</button>
      <button type="button" @click="checkTarget() && revoke()">Revoke</button>
    </div>
  </div>
  <div v-else-if="userType == 1">
    <h2>Verifier Menu</h2>
    <div>
      <div>Target address: <input v-model="targetAddress" placeholder="target address" /></div>
      Credit score
      <button type="button" @click="checkTarget() && check(1)">Check approval</button>
      <button type="button" @click="verifyScore()">Get</button>
    </div>
    <div>
      Credit report
      <button type="button" @click="checkTarget() && check(2)">Check approval</button>
      <button type="button" @click="verifyReport()">Get</button>
    </div>
  </div>
  <div v-else>
    <h2>Writer Menu</h2>
    <div>
      <div>Target address: <input v-model="targetAddress" placeholder="target address" /></div>
      <input v-model="creditValue" placeholder="credit value" />
      <button type="button" @click="checkTarget() && writeScore()">Write score</button>
      <button type="button" @click="checkTarget() && writeReport()">Write report</button>
    </div>
  </div>
</template>

<style>
.banner {
    color: lightseagreen;
    font-size: larger;
    margin: 3px;
    font-weight: bold;
}

.result {
    color: lightcoral;
}

.box {
    border: double;
    padding: 2px;
    margin: 5px;
}

.passive {
  background-image: linear-gradient(1deg, #1F306A, #0F6C7A 99%) !important;
  color: #FFF !important;
}

button {
  color: #CCC;
  background-image: linear-gradient(1deg, #4F58FD, #149BF3 99%);
  background-color: #3EB2FD;
  margin: 5px;
  border-radius: 100px;
  border-width: 0;
  box-shadow: none;
  box-sizing: border-box;
  font-family: CircularStd,sans-serif;
}

input.show {
  width:100%;
  text-align: center;
  font-size: large;
  border-radius: 100px;
}

input.target {
  font-size: medium;
}

</style>
