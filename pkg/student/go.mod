module test.pkg/student

go 1.14

require (
	github.com/JeeShao/dependence v1.1.1
	github.com/pkg/errors v0.9.1
)

replace (
	github.com/JeeShao/dependence => github.com/JeeShao/dependence v1.1.2
	github.com/pkg/errors => github.com/pkg/errors v0.9.0
)
