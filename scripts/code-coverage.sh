COVERAGE_THRESHOLD=80

mkdir -p ./reports
rm -rf ./reports/*

go test -coverprofile=./reports/coverage.out ./...
go tool cover -html=./reports/coverage.out -o ./reports/coverage.html
TOTAL_COVERAGE=$(go tool cover -func=reports/coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+' | awk '{print int($1)}')
echo "Threshold:              $COVERAGE_THRESHOLD %"
echo "Current Test Coverage:  $TOTAL_COVERAGE %"
if (($TOTAL_COVERAGE < $COVERAGE_THRESHOLD)); then exit 1; fi
