package commands

import (
	"gravity-simulation/models"
	"gravity-simulation/utils"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func clearCommands(buffer *models.Buffer) {
	buffer.Commands = []string{}
	buffer.Text = []string{}
}

func executeCommand(buffer *models.Buffer, bodies *[]models.Body, central *models.Body, options *models.Options) {
	switch strings.Join(buffer.Commands, "") {
	// BODIES
	case "br":
		filtered := (*bodies)[:0]

		for _, body := range *bodies {
			minPos := utils.Vector3MinValue(body.Position)
			maxPos := utils.Vector3MaxValue(body.Position)

			if minPos >= -1500 && maxPos <= 1500 {
				filtered = append(filtered, body)
			}
		}

		*bodies = filtered

	// CENTRAL
	case "cs":
		central.Mass = float32(30000.0)
		central.Color = rl.Yellow
		central.Radius = 30.0
	case "cb":
		central.Mass = float32(3000000.0)
		central.Color = rl.Black
		central.Radius = 5.0
	case "ct":
		central.Mass = float32(50000.0)
		central.Color = rl.Blue
		central.Radius = 20.0
	case "cg":
		central.Mass = float32(30000.0)
		central.Color = rl.Red
		central.Radius = 50.0

	// OPTIONS
	case "ob":
		options.DrawBounds = !options.DrawBounds

	// SPAW
	case "sd":
		*bodies = append(*bodies, SpawDust(*central, 1000)...)

	case "sp":
		*bodies = append(*bodies, SpawPlanet(*central))
	}

	clearCommands(buffer)
}

func GetInput(buffer *models.Buffer, bodies *[]models.Body, central *models.Body, options *models.Options) {
	switch len(buffer.Commands) {
	case 0:
		MainMenu(buffer)
	case 1:
		switch buffer.Commands[0] {
		case "b":
			BodiesMenu(buffer)
		case "c":
			CentralMenu(buffer)
		case "o":
			OptionsMenu(buffer)
		case "s":
			SpawMenu(buffer)
		}
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		clearCommands(buffer)
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		executeCommand(buffer, bodies, central, options)
	}
}

func MainMenu(buffer *models.Buffer) {
	buffer.MenuOptions = "Main Menu: [B]odies - [C]entral - [O]ptions - [S]paw"

	if rl.IsKeyPressed(rl.KeyC) {
		buffer.Commands = append(buffer.Commands, "c")
		buffer.Text = append(buffer.Text, "(C)entral")
	}
	if rl.IsKeyPressed(rl.KeyS) {
		buffer.Commands = append(buffer.Commands, "s")
		buffer.Text = append(buffer.Text, "(S)pawn")
	}
	if rl.IsKeyPressed(rl.KeyB) {
		buffer.Commands = append(buffer.Commands, "b")
		buffer.Text = append(buffer.Text, "(B)odies")
	}
	if rl.IsKeyPressed(rl.KeyO) {
		buffer.Commands = append(buffer.Commands, "o")
		buffer.Text = append(buffer.Text, "(O)ptions")
	}
}

func SpawMenu(buffer *models.Buffer) {
	buffer.MenuOptions = "Spaw Menu: [D]ust - [P]lanet"

	if rl.IsKeyPressed(rl.KeyD) {
		buffer.Commands = append(buffer.Commands, "d")
		buffer.Text = append(buffer.Text, "(D)ust")
	}
	if rl.IsKeyPressed(rl.KeyP) {
		buffer.Commands = append(buffer.Commands, "p")
		buffer.Text = append(buffer.Text, "(P)lanet")
	}
}

func OptionsMenu(buffer *models.Buffer) {
	buffer.MenuOptions = "Options Menu: Draw [B]ounds"

	if rl.IsKeyPressed(rl.KeyB) {
		buffer.Commands = append(buffer.Commands, "b")
		buffer.Text = append(buffer.Text, "Draw (B)ounds")
	}
}

func BodiesMenu(buffer *models.Buffer) {
	buffer.MenuOptions = "Bodies Menu: [R]remove Aways bodies"

	if rl.IsKeyPressed(rl.KeyR) {
		buffer.Commands = append(buffer.Commands, "r")
		buffer.Text = append(buffer.Text, "(R)emove Away Bodies")
	}
}

func CentralMenu(buffer *models.Buffer) {
	buffer.MenuOptions = "Central Menu: [B]lack Hole - [G]iant Red - [S]un - S[t]ar"

	if rl.IsKeyPressed(rl.KeyS) {
		buffer.Commands = append(buffer.Commands, "s")
		buffer.Text = append(buffer.Text, "(S)un")
	}

	if rl.IsKeyPressed(rl.KeyB) {
		buffer.Commands = append(buffer.Commands, "b")
		buffer.Text = append(buffer.Text, "(B)lack Hole")
	}

	if rl.IsKeyPressed(rl.KeyT) {
		buffer.Commands = append(buffer.Commands, "t")
		buffer.Text = append(buffer.Text, "S(t)ar")
	}

	if rl.IsKeyPressed(rl.KeyG) {
		buffer.Commands = append(buffer.Commands, "g")
		buffer.Text = append(buffer.Text, "(G)iant Red")
	}
}
