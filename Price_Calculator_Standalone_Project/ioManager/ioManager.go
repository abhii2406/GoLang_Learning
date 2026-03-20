package iomanager

type IoManager interface{
		ReadLines()([]string,error)
		WriteJSON(data any)error
	}