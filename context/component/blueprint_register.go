package component

import "strconv"

type BlueprintRegister map[string]Blueprint

func (register BlueprintRegister) GetBlueprint(bpName string) Blueprint {
	return register[bpName]
}

func (register BlueprintRegister) HasBlueprint(bpName string) bool {
	_, ok := register[bpName]
	return ok
}

func (register BlueprintRegister) AddBlueprint(bp Blueprint) {
	register[bp.GetItem()] = bp
}

func (register BlueprintRegister) GetAllRecipes() []string {
	bps := make([]string, 0, len(register))
	for _, bp := range register {
		aux := bp.GetItem() + " ="
		length := len(bp.GetAllCosts())
		for item, cost := range bp.GetAllCosts() {
			aux += " " + strconv.Itoa(cost) + "x " + item
			if length > 1 {
				aux += " +"
			}
			length--
		}
		bps = append(bps, aux)
	}
	return bps
}
