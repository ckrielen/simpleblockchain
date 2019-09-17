# Simple Block Chain Go Project

My first Go Project trying to make a simple Block chain 

The Chain allows you to add previous/current block to the Chain and get a new block via the method ___CreateNewBlock___

The block are linked with the Previous Hash Value and the Index (ToDo check on index on validate and Add)

The Block has methods
*  ___createHash___ to creat the hash of the block
* ___Validate___ to validate is correct with prevHash
* ___ToString___ to help with printing to console

The first block is created with a fixed prevHash


## TODO
[ ]  Update Add to check the Index
[ ] Update Validate to check the Index