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

	for count := 0; count < len(orderbook.Asks); count++ {
		ask := orderbook.Asks[count]
		if order.Price <= ask.Price {
			if ask.Volume <= order.Volume {
				trades = append(trades, &Trade{Bid: order, Ask: ask, Volume: ask.Volume, Price: ask.Price})
				order.Volume -= ask.Volume
				ask.Volume = 0
				orderbook.Asks = append(orderbook.Asks[:count], orderbook.Asks[count+1:]...)
				count--
			} else {
				trades = append(trades, &Trade{Bid: order, Ask: ask, Volume: order.Volume, Price: ask.Price})
				ask.Volume -= order.Volume
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

func (orderbook *Orderbook) TradeBids(order *Order) {
	trades := []*Trade{}

	for count := 0; count < len(orderbook.Bids); count++ {
		bid := orderbook.Bids[count]
		if order.Price == 0 || bid.Price <= order.Price {
			if bid.Volume <= order.Volume {
				trades = append(trades, &Trade{Bid: bid, Ask: order, Volume: bid.Volume, Price: bid.Price})
				order.Volume -= bid.Volume
				bid.Volume = 0
				orderbook.Bids = append(orderbook.Bids[:count], orderbook.Bids[count+1:]...)
				count--
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
