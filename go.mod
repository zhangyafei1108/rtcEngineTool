module rtcEngineTool

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190418165655-df01cb2cc480
	golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190419153524-e8e3143a4f4a
	golang.org/x/text => github.com/golang/text v0.3.0
)

go 1.12
