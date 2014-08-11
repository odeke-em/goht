package ordset

import (
    "testing"
)

const (
    Stocks = 8727
    Software = 21828
    Logistics = 100
    Marketing = 929
    Hospitality = 2346
    SafariBookings = 2938
)

func TestUsage(t *testing.T) {
    od := New()

    contractors := map[uint]string {
        Logistics: "Emmanuel Odeke", Stocks: "GoldmanSachs",
        Software: "Piatzi Analytics", SafariBookings: "Kijongo Bay Hospitality",
        Hospitality: "Kijongo Bay Resort", Marketing: "Wild Wild West(www) Inc",
        TourGuide: "Elaete Tours", InsuranceProvider: "ITriv Insurance Company",
    }

    for k, v := range contractors {
        od.Add(uint(k), v)
    }

    for q, v := range(contractors) {
        retr, existance := od.Get(uint(q))
        if existance != true {
            t.Errorf("Item at index %v should exist", existance)
        } else if retr != v {
            t.Errorf("Expected: %v instead got: %v", v, retr)
        }
    }

    _, existed := od.Remove(Logistics)
    if existed != true {
        t.Errorf("Logistics contractor pre-existed, reported as non existant on last remove")
    }

    // Double check
    dRetr, dCheck := od.Get(Logistics)
    if dCheck != false {
        t.Errorf("Logistics member got removed yet is reported as found")
    } else if dRetr != nil {
        t.Errorf("Value for Logistics got removed yet is reported as found")
    }
}
