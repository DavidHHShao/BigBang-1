#!/usr/bin/env bash
bazel clean
bazel run //:gazelle

./run_unit_test.sh  //test/lambda/feed/feed_events_table_creator:go_default_test
./run_unit_test.sh  //test/lambda/TCR/tcr_table_creator:go_default_test

./run_unit_test.sh  //test/lambda/feed/profile:go_default_test UserAuth

./run_unit_test.sh  //test/lambda/TCR/project:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/add_milestone:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/activate_milestone:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/finalize_milestone:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/objective:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/batch_objectives:go_default_test AdminAuth

./run_unit_test.sh  //test/lambda/TCR/get_project:go_default_test NoAuth
./run_unit_test.sh  //test/lambda/TCR/get_project_list:go_default_test  NoAuth
./run_unit_test.sh  //test/lambda/TCR/get_milestone:go_default_test  NoAuth
./run_unit_test.sh  //test/lambda/TCR/get_objective:go_default_test  NoAuth


./run_unit_test.sh  //test/lambda/TCR/rating_vote:go_default_test AdminAuth

./run_unit_test.sh  //test/lambda/TCR/add_proxy:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/update_available_delegate_votes:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/update_received_delegate_votes:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/update_batch_available_delegate_votes:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/update_batch_received_delegate_votes:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/add_proxy_voting_for_principal:go_default_test AdminAuth

./run_unit_test.sh  //test/lambda/TCR/get_rating_vote_list:go_default_test AdminAuth AdminAuth
./run_unit_test.sh  //test/lambda/TCR/get_batch_rating_vote_list:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/get_proxy_voting_info:go_default_test UserAuth
./run_unit_test.sh  //test/lambda/TCR/get_batch_proxy_voting_info:go_default_test UserAuth

./run_unit_test.sh  //test/lambda/TCR/finalize_validators:go_default_test  AdminAuth
./run_unit_test.sh  //test/lambda/TCR/get_finalized_validators:go_default_test NoAuth
./run_unit_test.sh  //test/lambda/TCR/get_batch_finalized_validators:go_default_test NoAuth

./run_unit_test.sh  //test/lambda/TCR/delete_objective:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/delete_batch_objectives:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/delete_milestone:go_default_test AdminAuth
./run_unit_test.sh  //test/lambda/TCR/delete_project:go_default_test AdminAuth

./run_unit_test.sh  //test/lambda/TCR/get_proxy_list:go_default_test NoAuth
./run_unit_test.sh  //test/lambda/TCR/get_project_id_by_admin:go_default_test NoAuth
./run_unit_test.sh  //test/lambda/TCR/delete_proxy:go_default_test AdminAuth
