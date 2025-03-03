package helper

import (
	"bytes"
	"encoding/json"
	"go-xendit-rmt003/entity"
	"net/http"
)

func CreateInvoice(product entity.ProductRequest, customer entity.CustomerRequest) (*entity.Invoice, error) {
	apiKey := "xnd_development_EV8LLejJV0hNNoeeNjFIfAr5uTXg99uq2hmvJfw1KKcvoMwYGH0JOyGPq38AP4s"
	apiUrl := "https://api.xendit.co/v2/invoices"

	bodyRequest := map[string]interface{}{
		"external_id":      "1",
		"amount":           product.Price,
		"description":      "Dummy Invoice RMT003",
		"invoice_duration": 86400,
		"customer": map[string]interface{}{
			"name":  customer.Name,
			"email": customer.Email,
		},
		"currency": "IDR",
		"items": []interface{}{
			map[string]interface{}{
				"name":     product.Name,
				"quantity": 1,
				"price":    15000,
			},
		},
	}

	reqBody, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(apiKey, "")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var resInvoice entity.Invoice
	if err := json.NewDecoder(response.Body).Decode(&resInvoice); err != nil {
		return nil, err
	}

	return &resInvoice, nil

}

// {"id":"65a738de263bd04f8f4b7194","external_id":"payment-link-example","user_id":"57fdbb445eec38910d3a4c47","status":"PENDING","merchant_name":"Your Company","merchant_profile_picture_url":"https://xnd-companies.s3.amazonaws.com/prod/1476344224287_930.png","amount":100000,"description":"Invoice Demo #123","expiry_date":"2024-01-18T02:18:06.783Z","invoice_url":"https://checkout-staging.xendit.co/latest/65a738de263bd04f8f4b7194","available_banks":[{"bank_code":"MANDIRI","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0},{"bank_code":"BRI","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0},{"bank_code":"BNI","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0},{"bank_code":"PERMATA","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0}],"available_retail_outlets":[],"available_ewallets":[],"available_qr_codes":[{"qr_code_type":"QRIS"}],"available_direct_debits":[],"available_paylaters":[{"paylater_type":"KREDIVO"},{"paylater_type":"UANGME"},{"paylater_type":"AKULAKU"},{"paylater_type":"ATOME"}],"should_exclude_credit_card":false,"should_send_email":false,"success_redirect_url":"https://www.google.com","failure_redirect_url":"https://www.google.com","created":"2024-01-17T02:18:06.932Z","updated":"2024-01-17T02:18:06.932Z","currency":"IDR","items":[{"name":"Air Conditioner","quantity":1,"price":100000,"category":"Electronic","url":"https://yourcompany.com/example_item"}],"fees":[{"type":"ADMIN","value":5000}],"customer":{"given_names":"John","surname":"Doe","email":"johndoe@example.com","mobile_number":"+6287774441111","addresses":[{"city":"Jakarta Selatan","country":"Indonesia","postal_code":"12345","state":"Daerah Khusus Ibukota Jakarta","street_line1":"Jalan Makan","street_line2":"Kecamatan Kebayoran Baru"}]},"customer_notification_preference":{"invoice_created":["whatsapp","email","viber"],"invoice_reminder":["whatsapp","email","viber"],"invoice_paid":["whatsapp","email","viber"]}}%
