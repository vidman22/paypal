package paypal

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) CreatePartnerReferral(ctx context.Context, params BusinessEntity) (*CreatePartnerReferralResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s%s", c.APIBase, "/v2/customer/partner-referrals"), params)
	response := &CreatePartnerReferralResponse{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}

func (c *Client) ShowReferralData(ctx context.Context, partnerReferralID string) (*ShowReferralDataResponse, error) {
	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s%s", c.APIBase, fmt.Sprintf("/v2/customer/partner-referrals/%s", partnerReferralID)), nil)
	response := &ShowReferralDataResponse{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}
