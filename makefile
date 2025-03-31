vercel_prd: 
	@echo "Deploying to Vercel production..."
	vercel --prod --cwd ./vercel-functions

run:
	@echo "Running go server..."
	go run main.go

test: 
	@echo "Running go tests..."
	go test ./... -v