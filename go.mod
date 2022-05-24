module github.com/emer/axon

go 1.15

//// This file https://stackoverflow.com/questions/1274057/how-can-i-make-git-forget-about-a-file-that-was-tracked-but-is-now-in-gitign
replace github.com/emer/emergent => ../emergent

require (
	github.com/c2h5oh/datasize v0.0.0-20200825124411-48ed595a09d2
	github.com/emer/emergent v1.2.5
	github.com/emer/empi v1.0.12
	github.com/emer/etable v1.0.45
	github.com/goki/gi v1.2.17
	github.com/goki/ki v1.1.5
	github.com/goki/mat32 v1.0.10
)
