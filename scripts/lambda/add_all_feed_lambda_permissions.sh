#!/usr/bin/env bash

echo "adding lambda permission for attach_session"
add_feed_lambda_permission.sh attach_session $1 $2 attach-session

echo "adding lambda permission for deactivate_actor"
add_feed_lambda_permission.sh deactivate_actor $1 $2 deactivate-actor

echo "adding lambda permission for feed_post"
add_feed_lambda_permission.sh feed_post $1 $2 feed-post

echo "adding lambda permission for feed_upvote"
add_feed_lambda_permission.sh feed_upvote $1 $2 feed-upvote

echo "adding lambda permission for get_batch_posts"
add_feed_lambda_permission.sh get_batch_posts $1 $2 get-batch-posts

echo "adding lambda permission for get_feed_post"
add_feed_lambda_permission.sh get_feed_post $1 $2 get-feed-post

echo "adding lambda permission for get_profile"
add_feed_lambda_permission.sh get_profile $1 $2 get-profile

echo "adding lambda permission for get_recent_posts"
add_feed_lambda_permission.sh get_recent_posts $1 $2 get-recent-posts

echo "adding lambda permission for get_recent_votes"
add_feed_lambda_permission.sh get_recent_votes $1 $2 get-recent-votes

echo "adding lambda permission for get_session"
add_feed_lambda_permission.sh get_session $1 $2 get-session

echo "adding lambda permission for profile"
add_feed_lambda_permission.sh profile $1 $2 profile

echo "adding lambda permission for refuel"
add_feed_lambda_permission.sh refuel $1 $2 refuel

echo "adding lambda permission for dev_refuel"
add_feed_lambda_permission.sh dev_refuel $1 $2 dev-refuel
