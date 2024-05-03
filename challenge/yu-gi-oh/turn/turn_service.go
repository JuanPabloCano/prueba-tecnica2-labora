package turn

import (
	"ejercicio-poo/challenge/yu-gi-oh/card"
	"ejercicio-poo/challenge/yu-gi-oh/world"
)

func Start() {
	for {

	}
}

func options() {
	selectedCard := card.GetRandomCard()
	selectedWorld := world.GetRandomWorld(world.Storage)
	increment := selectedWorld.BonusIncrement(selectedCard.Attribute)
	selectedCard.IncrementAtk(increment)
}
