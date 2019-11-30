package server

import (
	"context"
	"time"

	"github.com/Beaxhem/manage/backend/internal/utils"
	"github.com/Beaxhem/manage/backend/pkg/db"
	"github.com/Beaxhem/manage/backend/pkg/financial"
	"github.com/Beaxhem/manage/backend/pkg/logger"
	"github.com/Beaxhem/manage/backend/pkg/quotes"

	"gopkg.in/mgo.v2/bson"
)

func (s *Server) GetIncome(ctx context.Context, params *financial.Params) (*financial.Response, error) {
	dataStore := db.NewDataStore()

	defer dataStore.Close()
	store := db.GetStore(dataStore, "income")

	var payoffs []*financial.Payoff
	var query interface{}
	var sortFields []string
	if params.GetQuery() != nil {
		query, _ = utils.GetQuery(params.GetQuery().Querystring)
		sortFields = params.GetQuery().SortFields
	}

	err := store.GetAll(&payoffs, sortFields, query)
	if err != nil {
		logger.ErrorFunc(err)
		return new(financial.Response), err
	}

	return &financial.Response{Income: payoffs}, nil
}

func NewIncome(quote *quotes.Quote) {

	dataStore := db.NewDataStore()

	defer dataStore.Close()
	store := db.GetStore(dataStore, "income")

	for _, sup := range quote.Supplierids {
		var total float32
		var totalWithProfit float32
		sector := make(map[string]float32)

		for _, p := range quote.Products {
			if p.Product.Userid == sup.Id {
				sector[p.Product.Sector] += float32(p.GetQty()) * float32(p.GetProduct().Sellingprice)
				totalWithProfit += float32(p.GetQty()) * float32(p.GetProduct().Sellingprice)
				total += float32(p.GetQty()) * float32(p.GetProduct().Buyingprice)
			}
		}

		sectorProfit := financial.MapToSectorSlice(sector)

		newPayoff := financial.Payoff{Sectors: sectorProfit, Timestamp: time.Now().Unix(), Toreceive: totalWithProfit, Profitless: total, Quoteid: quote.Id, Supplierid: sup.Id}

		err := store.Insert(newPayoff)
		if err != nil {
			logger.ErrorFunc(err)

		}
	}
}

func DeleteIncome(quote *quotes.Quote) {
	dataStore := db.NewDataStore()

	defer dataStore.Close()

	store := db.GetStore(dataStore, "income")

	_, err := store.C.RemoveAll(bson.M{"quoteid": quote.Id})
	if err != nil {
		logger.ErrorFunc(err)
	}

}

func (s *Server) NewBank(ctx context.Context, bank *financial.Bank) (*financial.EmptyResponse, error) {
	dataStore := db.NewDataStore()

	defer dataStore.Close()
	store := db.GetStore(dataStore, "banks")

	err := store.Insert(bank)
	if err != nil {
		logger.ErrorFunc(err)
	}
	return new(financial.EmptyResponse), nil
}

func (s *Server) GetBanks(ctx context.Context, req *financial.Request) (*financial.Response, error) {
	dataStore := db.NewDataStore()

	defer dataStore.Close()
	store := db.GetStore(dataStore, "banks")

	var banks []*financial.Bank
	err := store.GetAll(&banks, nil, nil)
	if err != nil {
		logger.ErrorFunc(err)
	}

	return &financial.Response{Banks: banks}, nil
}

func (s *Server) ToDestination(ctx context.Context, params *financial.Params) (*financial.Response, error) {

	go func(name string, amount float32) {
		dataStore := db.NewDataStore()

		defer dataStore.Close()
		store := db.GetStore(dataStore, "expenses")

		err := store.Insert(financial.Expense{Timestamp: time.Now().Unix(), Name: name, Amount: amount})
		if err != nil {
			logger.ErrorFunc(err)

		}
	}(params.To, params.Amount)

	dataStore := db.NewDataStore()

	defer dataStore.Close()
	store := db.GetStore(dataStore, "banks")

	var bank financial.Bank

	err := store.GetElement(&bank, bson.M{"name": params.To})
	if err != nil {
		logger.ErrorFunc(err)
		return new(financial.Response), err
	}

	bank.Money += params.Amount

	err = store.C.Update(bson.M{"name": bank.Name}, bank)
	if err != nil {
		logger.ErrorFunc(err)
		return new(financial.Response), err
	}

	return new(financial.Response), nil

}

func (s *Server) GetExpenses(ctx context.Context, params *financial.Params) (*financial.Expenses, error) {
	dataStore := db.NewDataStore()

	defer dataStore.Close()
	store := db.GetStore(dataStore, "expenses")

	var exps []*financial.Expense

	err := store.GetAll(&exps, nil, nil)
	if err != nil {
		logger.ErrorFunc(err)
		return new(financial.Expenses), err
	}

	return &financial.Expenses{Expense: exps}, nil
}

func (s *Server) Pay(ctx context.Context, params *financial.PaymentParams) (*financial.EmptyResponse, error) {
	dataStore := db.NewDataStore()

	defer dataStore.Close()
	store := db.GetStore(dataStore, "income")

	var payoff financial.Payoff
	err := store.GetElement(&payoff, bson.M{"quoteid": params.Id})
	if err != nil {
		logger.ErrorFunc(err)
		return new(financial.EmptyResponse), err
	}

	if payoff.Toreceive-params.Amount < 0 {

		return new(financial.EmptyResponse), nil
	}

	payoff.Toreceive -= params.Amount
	payoff.Paid += params.Amount

	err = store.C.Update(bson.M{"quoteid": params.Id}, payoff)
	if err != nil {
		logger.ErrorFunc(err)
		return new(financial.EmptyResponse), err
	}

	return new(financial.EmptyResponse), nil
}
