package commands

import (
    "github.com/google/uuid"
)

type DeleteContactCommand struct {
    ID uuid.UUID `validate:"required,uuid"`
}

func NewDeleteContactCommand(id uuid.UUID) DeleteContactCommand {
    return DeleteContactCommand{ID: id}
}