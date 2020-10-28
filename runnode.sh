#!/bin/bash


hsd init hschain --chain-id=one
hsd add-genesis-account $(hscli keys show validator -a) 100000000000syscoin,500000000000000uhst

hsd gentx --name validator


hsd collect-gentxs
