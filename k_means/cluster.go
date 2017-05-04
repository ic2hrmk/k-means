package k_means

import "errors"

type Cluster struct {
	clusterId     int32
	centralValues []float64
	points        []*Point
}


func NewCluster(clusterId int32, cm *Point) *Cluster {
	var cmValues []float64
	for i := int32(0); i < cm.GetElementsCount(); i++ {
		cmValues = append(cmValues, cm.GetValue(i))
	}

	return &Cluster{
		clusterId: clusterId,
		centralValues: cmValues,
		points: []*Point{cm},
	}
}

func (cl *Cluster) GetId() int32 {
	return cl.clusterId
}

func (cl *Cluster) GetCentralValue(index int32) float64 {
	return cl.centralValues[index]
}

func (cl *Cluster) GetCentralValues() []float64 {
	return cl.centralValues
}

func (cl *Cluster) GetCentralValuesCount() int32 {
	return int32(len(cl.centralValues))
}

func (cl *Cluster) SetCentralValue(index int32, val float64) error {
	if index > int32(len(cl.centralValues)) {
		return errors.New("index out of range")
	} else {
		cl.centralValues[index] = val
	}

	return nil
}

func (cl *Cluster) GetPoint(index int32) *Point {
	return cl.points[index]
}

func (cl *Cluster) GetPointsCount() int32 {
	return int32(len(cl.points))
}

func (cl *Cluster) AddPoint(p *Point) {
	cl.points = append(cl.points, p)
}

func (cl *Cluster) RemovePoint(pointId int32) bool {
	for i, p := range cl.points {
		if p.GetId() == pointId {
			cl.points = append(cl.points[:i], cl.points[i+1:]...)
			return true
		}
	}

	return false
}