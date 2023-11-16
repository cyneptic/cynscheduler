package styles

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/cyneptic/cynscheduler/task"
)

var colors = map[string]lipgloss.Color{
	"Yellow":    lipgloss.Color("#FFFF00"),
	"Red":       lipgloss.Color("#FF0000"),
	"Orange":    lipgloss.Color("#FF5500"),
	"Teal":      lipgloss.Color("#00AAFF"),
	"Bubblegum": lipgloss.Color("#FF77BC"),
}

func GetMainStyle(str string) string {
	baseStyle := lipgloss.NewStyle().Padding(1, 1).Border(lipgloss.RoundedBorder()).BorderForeground(colors["Bubblegum"])
	return baseStyle.Render(str)
}

func GetTitleString(timer string) string {
	baseStyle := lipgloss.NewStyle().Bold(true).Foreground(colors["Teal"])
	s := baseStyle.SetString("Remaining Time of the Day: ").String()
	v := baseStyle.Copy().Foreground(colors["Yellow"]).SetString(timer).String()
	return lipgloss.NewStyle().Align(lipgloss.Center).Width(80).Render(s + v)
}

func GetLegendString(curTask *task.Task) string {
	var s string

	s += "(CTRL + F) - Finish Task"

	if curTask.Urgent && !curTask.Important {
		s += ", (CTRL + D) - Delegate Task"
	}

	baseStyle := lipgloss.NewStyle().Width(80).Align(lipgloss.Center).Foreground(lipgloss.Color("248"))

	return baseStyle.Render(s)
}

func GetStyledTable(tasks []*task.Task) string {
	baseStyle := lipgloss.NewStyle().Padding(0, 1).Bold(true)
	selectedStyle := baseStyle.Copy().Foreground(lipgloss.Color("#01BE85")).Background(lipgloss.Color("#00fF2F"))

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
		data = append(data, []string{fmt.Sprintf("%d", i+1), t.Name, t.Description, t.Timer.Timer().String(), action})
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
