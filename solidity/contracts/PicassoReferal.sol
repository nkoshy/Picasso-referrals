// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/** 
 * @title PicassoReferal
 * @dev Implements referal process. users are refered by Influencers 
 */
contract PicassoReferal {
   
    struct Influencer {
        string name; // name of the influencer
        bool referalEnabled;  // if true, then only referal works
        address influencerAddress; // influencer address
        string referalCode;   // referal code for Influencer
    }

    address public picassoadmin;

    mapping(string => Influencer) public referalCodeMapping;
    mapping(address => address) public feeReceipients;
    
    string[] public referalCodes;


    modifier onlyPicassoAdmin() {
        require(msg.sender == picassoadmin, "Caller is not picassoadmin");
        _;
    }

    modifier onlyIfReferalEnabled(string memory referalCode) {
        require(referalCodeMapping[referalCode].referalEnabled == true, "Referal code is not Enabled");
        _;
    }

    constructor() {
        picassoadmin = msg.sender;
    }


    function createInfluencer(string memory referalCode, string memory influencerName, address influencerAddress) public onlyPicassoAdmin {
     
        referalCodeMapping[referalCode] = Influencer({
                name: influencerName,
                influencerAddress: influencerAddress,
                referalEnabled: false,
                referalCode:referalCode
            });
            
        referalCodes.push(referalCode);
    }
    

    function enableRefereals(string memory referalCode) public onlyPicassoAdmin {
        referalCodeMapping[referalCode].referalEnabled = true;            
    }
    

    function referUser(string memory referalCode, address userAddress) public onlyPicassoAdmin onlyIfReferalEnabled(referalCode){
        require(
            msg.sender == picassoadmin,
            "Only picassoadmin can add influencers."
        );

        require(
            feeReceipients[userAddress] == address(0), "The user is already refered."
        );

        Influencer memory influencer = referalCodeMapping[referalCode];
        feeReceipients[userAddress] = influencer.influencerAddress;
    }


    function getInfluencerDetails(string memory referalCode) public view returns (string memory, bool, address, string memory) {
        Influencer memory influencer = referalCodeMapping[referalCode];
        return (influencer.name, influencer.referalEnabled, influencer.influencerAddress, influencer.referalCode);
    }
    
    
    function getReferalCodes() public view returns(string[] memory){
        return referalCodes;
    }


    function getFeeRecepient(address userAddress) public view returns (address) {
        return feeReceipients[userAddress];
    } 

}