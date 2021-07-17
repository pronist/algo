const Web3 = require('web3');
const web = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));

console.log(web.isConnected());
