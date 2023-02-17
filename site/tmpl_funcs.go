package site

import (
	"fmt"
)

func formatFloatToCurrency(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
