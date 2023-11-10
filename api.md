# Shared Response Types

- <a href="https://pkg.go.dev/github.com/example/example-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/example/example-go/internal/shared#Commit">Commit</a>
- <a href="https://pkg.go.dev/github.com/example/example-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/example/example-go/internal/shared#Discount">Discount</a>
- <a href="https://pkg.go.dev/github.com/example/example-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/example/example-go/internal/shared#ID">ID</a>
- <a href="https://pkg.go.dev/github.com/example/example-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/example/example-go/internal/shared#Override">Override</a>
- <a href="https://pkg.go.dev/github.com/example/example-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/example/example-go/internal/shared#SchedulePointInTime">SchedulePointInTime</a>

# ContractPricing

## Products

Response Types:

- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ProductListItem">ProductListItem</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductNewResponse">ContractPricingProductNewResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductGetResponse">ContractPricingProductGetResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductUpdateResponse">ContractPricingProductUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductListResponse">ContractPricingProductListResponse</a>

Methods:

- <code title="post /contract-pricing/products/create">client.ContractPricing.Products.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductNewParams">ContractPricingProductNewParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductNewResponse">ContractPricingProductNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/products/get">client.ContractPricing.Products.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductGetParams">ContractPricingProductGetParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductGetResponse">ContractPricingProductGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/products/update">client.ContractPricing.Products.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductUpdateParams">ContractPricingProductUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductUpdateResponse">ContractPricingProductUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/products/list">client.ContractPricing.Products.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductListParams">ContractPricingProductListParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingProductListResponse">ContractPricingProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RateCards

Params Types:

- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#RateParam">RateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#Rate">Rate</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#RateCard">RateCard</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardNewResponse">ContractPricingRateCardNewResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardGetResponse">ContractPricingRateCardGetResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardUpdateResponse">ContractPricingRateCardUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardListResponse">ContractPricingRateCardListResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardAddRateResponse">ContractPricingRateCardAddRateResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardSetRateCardProductsOrderResponse">ContractPricingRateCardSetRateCardProductsOrderResponse</a>

Methods:

- <code title="post /contract-pricing/rate-cards/create">client.ContractPricing.RateCards.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardNewParams">ContractPricingRateCardNewParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardNewResponse">ContractPricingRateCardNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/rate-cards/get">client.ContractPricing.RateCards.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardGetParams">ContractPricingRateCardGetParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardGetResponse">ContractPricingRateCardGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/rate-cards/update">client.ContractPricing.RateCards.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardUpdateParams">ContractPricingRateCardUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardUpdateResponse">ContractPricingRateCardUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/rate-cards/list">client.ContractPricing.RateCards.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardListParams">ContractPricingRateCardListParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardListResponse">ContractPricingRateCardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/rate-cards/addRate">client.ContractPricing.RateCards.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardService.AddRate">AddRate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardAddRateParams">ContractPricingRateCardAddRateParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardAddRateResponse">ContractPricingRateCardAddRateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contract-pricing/rate-cards/setRateCardProductsOrder">client.ContractPricing.RateCards.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardService.SetRateCardProductsOrder">SetRateCardProductsOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardSetRateCardProductsOrderParams">ContractPricingRateCardSetRateCardProductsOrderParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardSetRateCardProductsOrderResponse">ContractPricingRateCardSetRateCardProductsOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### MoveRateCardProducts

Response Types:

- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardMoveRateCardProductUpdateResponse">ContractPricingRateCardMoveRateCardProductUpdateResponse</a>

Methods:

- <code title="post /contract-pricing/rate-cards/moveRateCardProducts">client.ContractPricing.RateCards.MoveRateCardProducts.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardMoveRateCardProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardMoveRateCardProductUpdateParams">ContractPricingRateCardMoveRateCardProductUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractPricingRateCardMoveRateCardProductUpdateResponse">ContractPricingRateCardMoveRateCardProductUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Contracts

Response Types:

- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#Contract">Contract</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractNewResponse">ContractNewResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractGetResponse">ContractGetResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractListResponse">ContractListResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractAmendResponse">ContractAmendResponse</a>
- <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractUpdateEndDateResponse">ContractUpdateEndDateResponse</a>

Methods:

- <code title="post /contracts/create">client.Contracts.<a href="https://pkg.go.dev/github.com/example/example-go#ContractService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractNewParams">ContractNewParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractNewResponse">ContractNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contracts/get">client.Contracts.<a href="https://pkg.go.dev/github.com/example/example-go#ContractService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractGetParams">ContractGetParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractGetResponse">ContractGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contracts/list">client.Contracts.<a href="https://pkg.go.dev/github.com/example/example-go#ContractService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractListParams">ContractListParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractListResponse">ContractListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contracts/amend">client.Contracts.<a href="https://pkg.go.dev/github.com/example/example-go#ContractService.Amend">Amend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractAmendParams">ContractAmendParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractAmendResponse">ContractAmendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /contracts/setUsageFilter">client.Contracts.<a href="https://pkg.go.dev/github.com/example/example-go#ContractService.SetUsageFilter">SetUsageFilter</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractSetUsageFilterParams">ContractSetUsageFilterParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /contracts/updateEndDate">client.Contracts.<a href="https://pkg.go.dev/github.com/example/example-go#ContractService.UpdateEndDate">UpdateEndDate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractUpdateEndDateParams">ContractUpdateEndDateParams</a>) (<a href="https://pkg.go.dev/github.com/example/example-go">example</a>.<a href="https://pkg.go.dev/github.com/example/example-go#ContractUpdateEndDateResponse">ContractUpdateEndDateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Alerts

# CustomerAlerts

# Plans

# Credits

# CreditTypes

# Customers

## Plans

### PriceAdjustments

## Archive

## BillableMetrics

## Invoices

## BillingConfig

## Dashboards

# Ingest

# Usage

## Groups

# AuditLogs

# CustomFields
