package orderbook

import "sort"

type Orderbook struct {
	Bids []*Order
	Asks []*Order
}

func New() *Orderbook {
	orderbook := &Orderbook{}
	orderbook.Bids = []*Order{}
	orderbook.Asks = []*Order{}
	return orderbook
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	switch order.Side {
	case SideBid:
		return orderbook.TradeBids(order)
	case SideAsk:
		return orderbook.TradeAsks(order)
	}

	return nil, nil
}

func (orderbook *Orderbook) AddTo(order *Order) {
	switch order.Side {
	case SideAsk:
		{
			index := sort.Search(len(orderbook.Bids), func(i int) bool { return orderbook.Bids[i].Price > order.Price })
			orderbook.Bids = Append(order, orderbook.Bids, index)
		}
	case SideBid:
		{
			index := sort.Search(len(orderbook.Asks), func(i int) bool { return orderbook.Asks[i].Price < order.Price })
			orderbook.Asks = Append(order, orderbook.Asks, index)
		}
	}
}

func Append(order *Order, Add []*Order, index int) []*Order {
	Add = append(Add, order)
	copy(Add[index+1:], Add[:])
	Add[index] = order
	return Add
}

func (orderbook *Orderbook) TradeAsks(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}

	for i := 0; i < len(orderbook.Asks); i++ {
		a := orderbook.Asks[i]
		if order.Price <= a.Price {
			if a.Volume <= order.Volume {
				trades = append(trades, &Trade{Bid: order, Ask: a, Volume: a.Volume, Price: a.Price})
				order.Volume -= a.Volume
				a.Volume = 0
				orderbook.Asks = append(orderbook.Asks[:i], orderbook.Asks[i+1:]...)
				i--
			} else {
				trades = append(trades, &Trade{Bid: order, Ask: a, Volume: order.Volume, Price: a.Price})
				a.Volume -= order.Volume
				order.Volume = 0
				return trades, nil
			}
		} else {
			break
		}
	}

	if order.Kind == 1 {
		return trades, order
	} else {
		orderbook.AddTo(order)
	}
	return trades, nil

}

func (orderbook *Orderbook) TradeBids(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}

	for i := 0; i < len(orderbook.Bids); i++ {
		bid := orderbook.Bids[i]
		if order.Price == 0 || bid.Price <= order.Price {
			if bid.Volume <= order.Volume {
				trades = append(trades, &Trade{Bid: bid, Ask: order, Volume: bid.Volume, Price: bid.Price})
				order.Volume -= bid.Volume
				bid.Volume = 0
				orderbook.Bids = append(orderbook.Bids[:i], orderbook.Bids[i+1:]...)
				i--
			} else {
				trades = append(trades, &Trade{Bid: bid, Ask: order, Volume: order.Volume, Price: bid.Price})
				bid.Volume -= order.Volume
				order.Volume = 0
			}
			if order.Volume == 0 {
				return trades, nil
			}
		} else {
			break
		}
	}

	if order.Kind == 1 {
		return trades, order
	} else {
		orderbook.AddTo(order)
	}
	return trades, nil
}
