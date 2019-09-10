#########################
###      DEFS         ###
#########################

# Don't ask, for to understand it is to look 
# into the void and know the void is not only 
#looking back but also reading your emails.
SHELL=/bin/bash -e -o pipefail

AUTHORIZATIONHEADERKEY?=
APPLICATIONID?=
INTEGRATIONACCOUNTHASH?=

#########################
###      TARGETS      ###
#########################

.PHONY: test cover unit integration cleanup

test: ## Runs basic go test
	go test -v ./... --cover --coverprofile=wurlwind.out -short

cover: ## Generate coverage report
	go tool cover --html=wurlwind.out

integration: ## Perform integration tests
	go test -v ./... -run Integration