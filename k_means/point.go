package k_means

type Point struct {
	pointId   int32
	clusterId int32
	name      string
	values    []float64
}

func NewPoint(pointId int32, values []float64, name string) *Point {
	return &Point{
		pointId: pointId,
		clusterId: -1,
		name: name,
		values: values,
	}
}

func (p *Point) GetId() int32 {
	return p.pointId
}

func (p *Point) GetClusterId() int32 {
	return p.clusterId
}

func (p *Point) GetName() string {
	return p.name
}

func (p *Point) GetValue(index int32) float64 {
	return p.values[index]
}

func (p *Point) GetValues() []float64 {
	return p.values
}

func (p *Point) GetElementsCount() int32 {
	return int32(len(p.values))
}

func (p *Point) SetCluster(clusterId int32) {
	p.clusterId = clusterId
}