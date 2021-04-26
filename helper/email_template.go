package helper

import (
	"fmt"

	"github.com/hpazk/go-ticketing/database/model"
)

func PaymentOrderTemplate(event model.Event) string {
	emailBody := fmt.Sprintf(`
	<h2>Order Detail</h2>
		<table bordered>
		<tbody>
			<tr>
			<td>Event Title</td>
			<td>%s</td>
			</tr>
			<tr>
			<td>Descrption</td>
			<td>%s</td>
			</tr>
			<tr>
			<td>Price</td>
			<td>%f</td>
			</tr>
		</tbody>
		</table>
		<h2>Please make payment to 1909999100</h2>
		`, event.TitleEvent, event.Description, event.Price)

	return emailBody
}
