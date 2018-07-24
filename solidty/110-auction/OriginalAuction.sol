pragma solidity ^0.4.15;
import "OriginalToken.sol";
import "OriginalAsset.sol";

contract OriginalAuction {
    OriginalToken public Token;
    OriginalAsset public Asset;
    address[] public Bidders;
    address public Exhibitor;
    uint256 highestBid;
    address highestBidder;
    enum stages {
        PutUp,
        Registration,
        Bid
    }
    stages public Stage = stages.PutUp;
    event RegisterEvent(address bidder);
    event BidEvent(address bidder, uint256 value);
    event PutUpEvent(address exhibitor, uint256 value);
    event HiestBidderEvent(address bidder, uint256 value);

    modifier AtStage(stages _stage) {
        require(Stage == _stage);
        _;
    }
    modifier IsExhibitor() {
        require (msg.sender == Exhibitor);
        _;
    }
    modifier IsBidder() {
        bool isBidder = false;
        for (uint i = 0; i < Bidders.length; i++) {
            if (msg.sender == Bidders[i]) isBidder = true;
        }
        require(isBidder);
        _;
    }
}