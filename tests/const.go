package tests

const(
	TestServer = "http://localhost:8080"
	Version = "/v1"
	Createdemo = Version + "/demo"
	Updatedemo = Version + "/demo/1"
	Getdemoinfo = Updatedemo
	Postdemolist = Version + "/demo-all"
	Getdemolist = Postdemolist
)