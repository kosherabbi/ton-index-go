package main

import (
    // ... existing imports ...
    "github.com/kdimentionaltree/ton-index-go/index"
)

// ... existing code ...

// @summary Get Deploy Wallets and Tokens
// @description Get deploy wallets and tokens for an array of addresses (up to 40 per query)
// @id api_v3_get_deploy_wallets_and_tokens
// @tags blockchain
// @Accept       json
// @Produce      json
// @success      200 {object} index.DeployWalletsResponse
// @failure      400 {object} index.RequestError
// @param        addresses query []string true "Array of addresses to query (max 40)" collectionFormat(multi)
// @router       /api/v3/deployWalletsAndTokens [get]
// @security     APIKeyHeader
// @security     APIKeyQuery
func GetDeployWalletsAndTokens(c *fiber.Ctx) error {
    request_settings := GetRequestSettings(c, &settings)
    var deploy_req index.DeployWalletsRequest

    if err := c.QueryParser(&deploy_req); err != nil {
        return index.IndexError{Code: 422, Message: err.Error()}
    }

    if len(deploy_req.Addresses) == 0 {
        return index.IndexError{Code: 422, Message: "at least 1 address required"}
    }

    if len(deploy_req.Addresses) > 40 {
        return index.IndexError{Code: 422, Message: "maximum 40 addresses allowed"}
    }

    res, err := pool.QueryDeployWallets(deploy_req, request_settings)
    if err != nil {
        return err
    }

    return c.JSON(res)
}

// @summary Get Token Transactions
// @description Get transactions for a specific token, with optional filtering by operation type
// @id api_v3_get_token_transactions
// @tags blockchain
// @Accept       json
// @Produce      json
// @success      200 {object} index.TokenTransactionsResponse
// @failure      400 {object} index.RequestError
// @param        token_address query string true "Token address"
// @param        operation_type query string false "Operation type (buy or sell)"
// @router       /api/v3/tokenTransactions [get]
// @security     APIKeyHeader
// @security     APIKeyQuery
func GetTokenTransactions(c *fiber.Ctx) error {
    request_settings := GetRequestSettings(c, &settings)
    var tx_req index.TokenTransactionsRequest

    if err := c.QueryParser(&tx_req); err != nil {
        return index.IndexError{Code: 422, Message: err.Error()}
    }

    if tx_req.TokenAddress == "" {
        return index.IndexError{Code: 422, Message: "token_address is required"}
    }

    if tx_req.OperationType != "" && tx_req.OperationType != "buy" && tx_req.OperationType != "sell" {
        return index.IndexError{Code: 422, Message: "invalid operation_type, must be 'buy' or 'sell'"}
    }

    res, err := pool.QueryTokenTransactions(tx_req, request_settings)
    if err != nil {
        return err
    }

    return c.JSON(res)
}

func main() {
    // ... existing code ...

    // Add the new endpoints
    app.Get("/api/v3/deployWalletsAndTokens", GetDeployWalletsAndTokens)
    app.Get("/api/v3/tokenTransactions", GetTokenTransactions)

    // ... rest of the main function ...
}