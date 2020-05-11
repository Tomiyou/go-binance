package binance

import (
	"context"
	"encoding/json"
)

// ListDepositsService list deposits
type ListDepositsService struct {
	c         *Client
	asset     *string
	status    *int
	startTime *int64
	endTime   *int64
}

// Asset set asset
func (s *ListDepositsService) Asset(asset string) *ListDepositsService {
	s.asset = &asset
	return s
}

// Status set status
func (s *ListDepositsService) Status(status int) *ListDepositsService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *ListDepositsService) StartTime(startTime int64) *ListDepositsService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListDepositsService) EndTime(endTime int64) *ListDepositsService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListDepositsService) Do(ctx context.Context, opts ...RequestOption) (deposits []*Deposit, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/wapi/v1/getDepositHistory.html",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != nil {
		m["asset"] = *s.asset
	}
	if s.status != nil {
		m["status"] = *s.status
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res := new(DepositHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res.Deposits, nil
}

// DepositHistoryResponse define deposit history
type DepositHistoryResponse struct {
	Success  bool       `json:"success"`
	Deposits []*Deposit `json:"depositList"`
}

// Deposit define deposit info
type Deposit struct {
	InsertTime int64   `json:"insertTime"`
	Amount     float64 `json:"amount"`
	Asset      string  `json:"asset"`
	Status     int     `json:"status"`
	TxID       string  `json:"txId"`
}

// GetDepositAddressService list deposits
type GetDepositAddressService struct {
	c          *Client
	asset      *string
	status     *bool
	recvWindow *int64
	timestamp  *int64
}

// Asset set asset
func (s *GetDepositAddressService) Asset(asset string) *GetDepositAddressService {
	s.asset = &asset
	return s
}

// Status set status
func (s *GetDepositAddressService) Status(status bool) *GetDepositAddressService {
	s.status = &status
	return s
}

// ReceiveWindow set recvWindow
func (s *GetDepositAddressService) ReceiveWindow(recvWindow int64) *GetDepositAddressService {
	s.recvWindow = &recvWindow
	return s
}

// Timestamp set timestamp
func (s *GetDepositAddressService) Timestamp(timestamp int64) *GetDepositAddressService {
	s.timestamp = &timestamp
	return s
}

// Do send request
func (s *GetDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (address, addressTag string, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/depositAddress.html",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != nil {
		m["asset"] = *s.asset
	}
	if s.status != nil {
		m["status"] = *s.status
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	if s.timestamp != nil {
		m["timestamp"] = *s.timestamp
	}
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res := new(DepositAddressResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res.Address, res.AddressTag, nil
}

// DepositHistoryResponse define deposit address
type DepositAddressResponse struct {
	Address    string `json:"address"`
	Success    bool   `json:"success"`
	AddressTag string `json:"addressTag"`
	Asset      string `json:"asset"`
}
