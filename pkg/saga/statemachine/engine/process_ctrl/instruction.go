package process_ctrl

import (
	"github.com/seata/seata-go/pkg/saga/statemachine/statelang"
)

type Instruction interface {
}

type StateInstruction struct {
	StateName        string
	StateMachineName string
	TenantId         string
	End              bool
}

func (s StateInstruction) GetState(context ProcessContext) (statelang.State, error) {
	//TODO implement me
	panic("implement me")
}

func NewStateInstruction(stateMachineName string, tenantId string) *StateInstruction {
	return &StateInstruction{StateMachineName: stateMachineName, TenantId: tenantId}
}