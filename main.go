package main

import (
	_ "k-means/k_means"
	_"log"
)

func main() {
	//var test = []*k_means.Point{
	//	k_means.NewPoint(0, []float64{1, 1 }, ""),
	//	k_means.NewPoint(1, []float64{1.5, 2 }, ""),
	//	k_means.NewPoint(2, []float64{3, 4 }, ""),
	//	k_means.NewPoint(3, []float64{5, 7 }, ""),
	//	k_means.NewPoint(4, []float64{3.5, 5 }, ""),
	//	k_means.NewPoint(5, []float64{4.5, 5 }, ""),
	//	k_means.NewPoint(6, []float64{3.5, 4.5 }, ""),
	//}
	//
	//cls, err := k_means.Calc(test, 2, 100, k_means.EuclidDistance)
	//if err != nil {
	//	log.Fatal(err.Error())
	//} else {
	//	for _, cl := range cls {
	//		log.Println("Cluster", cl.GetId(), ":")
	//		for i := int32(0); i < cl.GetPointsCount(); i++ {
	//			p := cl.GetPoint(i)
	//			log.Println("Point", p.GetId(), ":", p.GetValues())
	//		}
	//		log.Println("Values:", cl.GetCentralValues())
	//	}
	//}

	Start()
}