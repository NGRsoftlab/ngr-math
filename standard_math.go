// Copyright 2020 NGR Softlab
//
package ngr_math

import (
	"errors"
	"math"
)

const (
	Intervals = 100
	Precision = 0.00001
)

type FuncParams struct {
	Start    float64 `json:"start"`
	End      float64 `json:"end"`
	Func     string  `json:"func"`
	Addition float64 `json:"addition"`
}

// Functions map
var Functions = map[string]interface{}{
	"0": makeLinX,
	"1": makeX2,
	"2": makeValDivX,
	"3": makeLogX,
	"4": makeSqrtX,
}

func DoMathFunc(params FuncParams) (res []map[string]float64, err error) {
	resFunc, ok := Functions[params.Func].(func(FuncParams) ([]map[string]float64, error))
	if !ok {
		return nil, errors.New("no such function")
	}

	return resFunc(params)
}

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
////////////////////////////////////////////////////////

func makeLinX(params FuncParams) (res []map[string]float64, err error) {

	delta := (params.End - params.Start) / Intervals

	for params.Start < params.End+delta/2 {
		res = append(res, map[string]float64{"X": params.Start, "Y": params.Start * params.Addition})
		params.Start += delta
	}

	return res, nil
}

func makeX2(params FuncParams) (res []map[string]float64, err error) {

	delta := (params.End - params.Start) / Intervals

	for params.Start < params.End+delta/2 {
		res = append(res, map[string]float64{"X": params.Start, "Y": math.Pow(params.Start, 2)})
		params.Start += delta
	}

	return res, nil
}

func makeValDivX(params FuncParams) (res []map[string]float64, err error) {

	delta := (params.End - params.Start) / Intervals

	for params.Start < params.End+delta/2 {
		if math.Abs(params.Start) >= Precision {
			res = append(res, map[string]float64{"X": params.Start, "Y": params.Addition / params.Start})
		}
		params.Start += delta
	}

	return res, nil
}

func makeLogX(params FuncParams) (res []map[string]float64, err error) {

	delta := (params.End - params.Start) / Intervals

	for params.Start < params.End+delta/2 {
		if params.Start > 0 {
			res = append(res, map[string]float64{"X": params.Start, "Y": math.Log(params.Start)})
		}
		params.Start += delta
	}

	return res, nil
}

func makeSqrtX(params FuncParams) (res []map[string]float64, err error) {

	delta := (params.End - params.Start) / Intervals

	for params.Start < params.End+delta/2 {
		if params.Start >= 0 {
			res = append(res, map[string]float64{"X": params.Start, "Y": math.Sqrt(params.Start)})
		}
		params.Start += delta
	}

	return res, nil
}
