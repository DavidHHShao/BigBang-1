package eth_config

import (
  "github.com/ethereum/go-ethereum/common"
  "math/big"
  "BigBang/internal/platform/postgres_config/feed/post_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
  "BigBang/internal/platform/postgres_config/feed/purchase_mps_record_config"
)

type Event interface{}

var NullHashString string = common.HexToHash("0x0").String()

type PostEventResult struct {
  Poster common.Address  // indexed
  BoardId common.Hash    // indexed
  ParentHash common.Hash
  PostHash common.Hash   // indexed
  IpfsPath common.Hash
  TypeHash [4]byte
  Timestamp *big.Int
}

type UpvoteEventResult struct {
  Actor   common.Address // indexed
  BoardId   common.Hash    // indexed
  PostHash  common.Hash    // indexed
  Value     *big.Int
  Timestamp *big.Int
}

type PurchaseReputationsEventResult struct {
  MsgSender   common.Address // indexed
  Purchaser   common.Address // indexed
  NumVetX     *big.Int
  NumReputation *big.Int
  Timestamp *big.Int
}

func (postEventResult *PostEventResult) ToPostRecord() *post_config.PostRecord {
  return &post_config.PostRecord {
    Actor:      postEventResult.Poster.String(),
    BoardId:    postEventResult.BoardId.String(),
    ParentHash: postEventResult.ParentHash.String(),
    PostHash:   postEventResult.PostHash.String(),
    PostType:   feed_attributes.CreatePostTypeFromHashStr("0x" + common.Bytes2Hex(postEventResult.TypeHash[:])).Value(),
  }
}

func (upvoteEventResult *UpvoteEventResult) ToPostVotesRecord() *post_votes_record_config.PostVotesRecord{
  return &post_votes_record_config.PostVotesRecord {
    Actor:   upvoteEventResult.Actor.String(),
    PostHash:  upvoteEventResult.PostHash.String(),
    VoteType: feed_attributes.CreateVoteTypeFromValue(upvoteEventResult.Value.Int64()),
  }
}

func (purchaseReputationsEventResult *PurchaseReputationsEventResult) ToPurchaseReputationsRecord() *purchase_mps_record_config.PurchaseMPsRecord {
  return &purchase_mps_record_config.PurchaseMPsRecord{
    Purchaser:  purchaseReputationsEventResult.Purchaser.String(),
    VetX: purchaseReputationsEventResult.NumVetX.Int64(),
    MPs: purchaseReputationsEventResult.NumReputation.Int64(),
  }
}