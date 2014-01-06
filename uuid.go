package main

import (
	"github.com/nu7hatch/gouuid"
)

type UUID *uuid.UUID

func MakeUUID() (UUID, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return (UUID)(uuid), nil
}
