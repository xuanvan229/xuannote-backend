package constant

// MsgFlags message content
var MsgFlags = map[string]string{
	SUCCESS :                         	"ok",
	ERROR :                           	"fail",
	WRONG_PASSWORD: 										"invalid password",
	INVALID_PARAMS :                  	"invalid params",
	COULD_NOT_CONNECT_TO_DATABASE: 			"Could not connect to database server",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL : 			"ERROR_UPLOAD_SAVE_IMAGE_FAIL",
	USER_DOES_NOT_EXIST: 			  				"Use does not exist",
	USER_NAME_IS_EXIST: 								"Username already exist",
	PERMISSION_ALREADY_EXIST: 					"Permission already exist",
	MISSING_SOME_FIELD:									"Missing some field",
}

// GetMsg from code to message
func GetMsg(code string) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return code
}