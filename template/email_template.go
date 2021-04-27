package template

import (
	"fmt"

	"github.com/hpazk/go-ticketing/apps/report"
	"github.com/hpazk/go-ticketing/database/model"
)

func InvoiceTemplate(event model.Event) string {
	emailBody := fmt.Sprintf(`
	<h2>Order Detail</h2>
		<table border="1">
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

func PaymentSuccessLayout(event report.Report) string {
	emailBody := fmt.Sprintf(`
	<h2>Webinar Detail</h2>
		<table border="1">
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
			<td>Webinar Link</td>
			<td>%s</td>
			</tr>
		</tbody>
		</table>`, event.TitleEvent, event.Description, event.LinkWebinar)

	return emailBody
}
