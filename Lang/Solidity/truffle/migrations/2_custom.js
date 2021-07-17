const hw = artifacts.require('First');
const hw1 = artifacts.require('Second');

module.exports = (deployer) => {
  deployer.deploy(hw).then(() => deployer.deploy(hw1, hw.address));
};
