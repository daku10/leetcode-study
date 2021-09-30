package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMyPow(t *testing.T) {

	type TestCase struct {
		argX float64
		argN int
		expect float64
	}

	testCases := []TestCase{
		{2.0000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
		{8.88023, 3, 700.28148},
		{34.00515, -3, 3e-05},
		{-2.00000, 2, 4.00000},
		{-13.62608, 3, -2529.95504},
		{1.00012, 1024, 1.13074},
		{0.44528, 0, 1.0000},
		{1.00001, 123456, 1.0},
		{-1.00001, 447125, -87.46403},
		{0.25519, -6, 3620.91299},
		{2.00000, -2147483648, 1},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := myPow(testCase.argX, testCase.argN)
			if actual != testCase.expect {
				t.Errorf("x: %v n: %v expect: %v actual: %v", testCase.argX, testCase.argN, testCase.expect, actual)
			}
		})
	}
}

func TestCreateDecimal(t *testing.T) {

	f := 1.23
	d := createDecimal(f)
	if !(d.scale == 2 && reflect.DeepEqual(d.arr, []int{3, 2, 1})) {
		t.Errorf("actual: %v", d)
	}
}

func TestProduct(t *testing.T) {

	d1 := createDecimal(0.2)
	d2 := createDecimal(0.2)

	p := product(d1, d2)
	if !(p.scale == 2 && reflect.DeepEqual(p.arr, []int{4, 0, 0})) {
		t.Errorf("actual: %v", p)
	}

}

func TestToString(t *testing.T) {
	d1 := createDecimal(2)

	s := d1.toString()
	if s != "2" {
		t.Errorf("actual: %v", s)
	}
}

func TestDestruct(t *testing.T) {
	m, l := destructPow(1023)
	if !(m == 9 && l == 511) {
		t.Errorf("m: %v l: %v", m, l)
	}
}

func TestFloor(t *testing.T) {
	d := createDecimal(1.2345)
	d.floorScale(1)
	if !(d.scale == 1 && reflect.DeepEqual(d.arr, []int{2,1})) {
		t.Errorf("d: %v", d)
	}


	d = createDecimal(0.1234567)
	d.floorScale(2)
	if !(d.scale == 2 && reflect.DeepEqual(d.arr, []int{2, 1,0})) {
		t.Errorf("d: %v", d)
	}

	d = createDecimal(1234.56789)
	d.floorScale(3)
	if !(d.scale == 3 && reflect.DeepEqual(d.arr, []int{7, 6, 5, 4, 3, 2, 1})) {
		t.Errorf("d: %v", d)
	}
}

func TestCeil(t *testing.T) {
	d := createDecimal(123456789.000)
	fmt.Println(d)
	d.floorArr(3)
	fmt.Println(d)
	if !(d.scale == 0 && reflect.DeepEqual(d.arr, []int{0, 0, 0, 0, 0, 0, 3, 2, 1})) {
		t.Errorf("d: %v", d)
	}
}
