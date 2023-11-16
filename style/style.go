package styles

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/cyneptic/cynscheduler/task"
)

func GetStyledTable(tasks []*task.Task) string {
	baseStyle := lipgloss.NewStyle().Padding(0, 1).Bold(true)
	selectedStyle := baseStyle.Copy().Foreground(lipgloss.Color("#01BE85")).Background(lipgloss.Color("#00fF2F"))
	colors := map[string]lipgloss.Color{
		"Yellow": lipgloss.Color("#FFFF00"),
		"Red":    lipgloss.Color("#FF0000"),
		"Orange": lipgloss.Color("#FF5500"),
		"Teal":   lipgloss.Color("#00AAFF"),
	}

	headers := []string{"Task", "Name", "Description", "Remaining", "Suggested_Action"}
	data := [][]string{}
	for i, t := range tasks {
		action := ""
		if t.Important && t.Urgent {
			action = "Do Now!"
		} else if t.Important {
			action = "Try To Do!"
		} else if t.Urgent {
			action = "Try To Delegate!"
		} else {
			action = "Delete or Excess Time."
		}
		data = append(data, []string{fmt.Sprintf("%d", i), t.Name, t.Description, t.Timer.Timer().String(), action})
	}

	CapitalizeHeaders := func(data []string) []string {
		for i := range data {
			data[i] = strings.ToUpper(data[i])
		}
		return data
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers(CapitalizeHeaders(headers)...).
		Rows(data...).
		Width(80).
		StyleFunc(func(row, col int) lipgloss.Style {
			even := row > 0 && row%2 == 0
			if row > 0 && data[row-1][col] == "Do Now!" {
				return baseStyle.Copy().Foreground(colors["Yellow"])
			}
			if row > 0 && data[row-1][col] == "Try To Do!" {
				return baseStyle.Copy().Foreground(colors["Teal"])
			}
			if row > 0 && data[row-1][col] == "Try To Delegate!" {
				return baseStyle.Copy().Foreground(colors["Orange"])
			}
			if row > 0 && data[row-1][col] == "Delete or Excess Time." {
				return baseStyle.Copy().Foreground(colors["Red"])
			}

			if row == 1 {
				return baseStyle.Copy().Foreground(selectedStyle.GetBackground())
			}

			if even {
				return baseStyle.Copy().Foreground(lipgloss.Color("248"))
			}
			_ = selectedStyle
			return baseStyle
		})
	return t.Render()
}
