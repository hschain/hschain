### A Next Generation Blockchain

For Testnet
# You can run all of these commands from your home directory
cd $HOME

# Initialize the genesis.json file that will help you to bootstrap the network
hsd init --chain-id=testing testing

# Create a key to hold your validator account
hscli keys add validator

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
hsd add-genesis-account $(hscli keys show validator -a) 10000hst

# Generate the transaction that creates your validator
hsd gentx --name validator

# Add the generated bonding transaction to the genesis file
hsd collect-gentxs

# Now its safe to start `hsd`
hsd start
