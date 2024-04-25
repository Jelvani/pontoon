types: miner, node, client


M1: "send"

action client TRANSACTION
{
	
	node n = _2
	client payee = _1
	// save required data in variable
	// rand() produces a random value of some type
	var satoshi = rand(type int)
	var payload = (M1,satoshi,payee,)
	payload => n
	
}
