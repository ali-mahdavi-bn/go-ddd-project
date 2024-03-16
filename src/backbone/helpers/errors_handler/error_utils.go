package errors_handler

import (
	"fmt"
)

func ErrorExist(err error) {
	if err != nil {
		fmt.Printf("Error :%v", err)
	}
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorExistOrSuccess(data any, err error) (result any, erro error) {
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}
