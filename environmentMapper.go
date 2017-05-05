package main

import (
	"k-means/k_means"
	"strconv"
	"errors"
)

type Claim struct {
	Id        int32  `json:"id"`
	Theme     string `json:"theme"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Claims []*Claim

const DimensionsNumber = 2	// Latitude and Longitude

func ClaimsToPoints(claims Claims) (points k_means.Points, err error) {
	for _, claim := range claims {
		var lat, lng float64

		lat, err = strconv.ParseFloat(claim.Latitude, 64)
		if err != nil {
			err = errors.New("claim #" + strconv.Itoa(int(claim.Id)) + " wrong latitude")
			return
		}

		lng, err = strconv.ParseFloat(claim.Longitude, 64)
		if err != nil {
			err = errors.New("claim #" + strconv.Itoa(int(claim.Id)) + " wrong longitude")
			return
		}

		points = append(points, k_means.NewPoint(claim.Id, []float64{lat, lng}, claim.Theme))
	}

	return
}