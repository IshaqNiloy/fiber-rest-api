package applibs

type CodeObject struct {
	Lang    string
	Message string
}

var RequestSuccess = CodeObject{
	Lang:    "en",
	Message: "success",
}

var RequestFailed = CodeObject{
	Lang:    "en",
	Message: "failed",
}

var InvalidRequestData = CodeObject{
	Lang:    "en",
	Message: "invalid request data",
}
