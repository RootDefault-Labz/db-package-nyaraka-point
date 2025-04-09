package nyarakadb

import (
    "github.com/RootDefault-Labz/RestCallPackage"
)

func (request *Request)CreateUserRequest() (string,error) {
    response,err:=network.MakePOSTRequest("Create User (DB)",request.Endpoint, request.Body, request.Headers)
    if err != nil {
        return response, err
    }
    return response, nil
}

func (request *Request)GetUserRequest() (string,error) {
    response,err:=network.MakeGETRequest("Get User (DB)",request.Endpoint, request.QueryParams,request.Headers)
    if err != nil {
        return response, err
    }
    return response, nil
}