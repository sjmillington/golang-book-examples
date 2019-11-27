package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//to compile 

/*
to compile 
set paths if required: 

export GOPATH=~/go
mkdir ~/go/bin
export GOBIN=$GOPATH/bin

cd math (be in this dir)
go install


*/