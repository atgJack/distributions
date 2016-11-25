package distributions

import (
  "math"
  "math/rand"
)

//The Pareto Distribution is a continuous probability distribution
// with parameters α > 0, k > 0.
//
// See: https://en.wikipedia.org/wiki/Pareto_distribution
type Pareto struct {
  Scale  float64
  Shape  float64
}

func (dist Pareto) validate() error {
  if dist.Scale <= 0 {
    return InvalidParamsError{ "Scale must be greater than zero." }
  }
  if dist.Shape <= 0 {
    return InvalidParamsError{ "Shape must be greater than zero." }
  }
  return nil
}

func (dist Pareto) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Shape <= 1.0) {
    return math.Inf(1), nil
  }
  result := (dist.Shape * dist.Scale) / (dist.Shape - 1)
  return result, nil
}

func (dist Pareto) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Shape <= 2.0) {
    return math.Inf(1), nil
  }
  result := (dist.Shape * dist.Scale * dist.Scale) / ((dist.Shape - 1) * (dist.Shape - 1) * (dist.Shape - 2))
  return result, nil
}

func (dist Pareto) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Shape < 3.0) {
    return math.NaN(), nil
  }
  result := 2 * (1 + dist.Shape) / (dist.Shape - 3) * math.Sqrt((dist.Shape - 2) / dist.Shape)
  return result, nil
}

func (dist Pareto) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Shape < 3.0) {
    return math.NaN(), nil
  }
  result := 6 * ((dist.Shape * dist.Shape * dist.Shape) + (dist.Shape * dist.Shape) - (6 * (dist.Shape - 2))) / (dist.Shape * (dist.Shape - 3) * (dist.Shape - 4))
  return result, nil
}

func (dist Pareto) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, _ := dist.Variance()
  if math.IsInf(variance, 0) {
    return math.Inf(1), nil
  }
  result := math.Sqrt(variance)
  return result, nil
}

func (dist Pareto) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, _ := dist.Variance()
  if math.IsInf(variance, 0) {
    return math.Inf(1), nil
  }
  mean, _ := dist.Mean()
  if math.IsInf(mean, 0) {
    return math.Inf(1), nil
  }
  result := math.Sqrt(variance) / mean
  return result, nil
}

func (dist Pareto) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < dist.Scale {
    return 0.0, nil
  }
  result := dist.Shape * math.Pow(dist.Scale, dist.Shape) / math.Pow(x, dist.Shape + 1)
  return result, nil
}

func (dist Pareto) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (x < dist.Scale) {
    return 0.0, nil
  }
  result := 1 - math.Pow(dist.Scale / x, dist.Shape)
  return result, nil
}

func (dist Pareto) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  value := dist.Scale / math.Pow(rand.Float64(), 1 / dist.Shape)
  return value, nil
}
