package kokots

import "github.com/google/uuid"

type Kokot struct {
	ID   string
	Name string
}

type KokotController struct {
	Kokots map[string]Kokot
}

func (k KokotController) Add(name string) *Kokot {
	// Generate UUID for new Kokot
	id := uuid.New()
	newKokot := Kokot{Name: name, ID: id.String()}
	k.Kokots[id.String()] = newKokot

	return &newKokot
}
func (k KokotController) Delete(id string) {
	delete(k.Kokots, id)
}
func (k KokotController) GetOne(id string) *Kokot {
	if ko, ok := k.Kokots[id]; !ok {
		return nil
	} else {
		return &ko
	}
}
func (k KokotController) GetAll() []Kokot {
	kokots := make([]Kokot, 0)
	for id, v := range k.Kokots {
		v.ID = id
		kokots = append(kokots, v)
	}
	return kokots
}
func (k KokotController) Update(ko Kokot) bool {
	if ko.ID == "" {
		return false
	}
	if _, ok := k.Kokots[ko.ID]; !ok {
		return false
	}
	k.Kokots[ko.ID] = ko
	return true
}

func NewKokotController() KokotController {
	k := KokotController{}
	k.Kokots = make(map[string]Kokot)
	return k
}
