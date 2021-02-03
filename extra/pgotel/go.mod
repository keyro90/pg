module github.com/keyro90/pg/extra/pgotel

go 1.15

replace github.com/keyro90/pg/v10 => ../..

require (
	github.com/keyro90/pg/v10 v10.5.0
	go.opentelemetry.io/otel v0.15.0
)
