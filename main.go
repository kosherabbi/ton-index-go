package main

import (
    // ... existing imports ...
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

    // Get factory data
    factoryData, err := pool.GetFactoryData()
    if err != nil {
        return err
    }

    // Add factory data to the response
    res.FactoryData = factoryData

    return c.JSON(res)
}

func main() {
    // ... existing code ...

    // Add the new endpoint
    app.Get("/api/v3/deployWalletsAndTokens", GetDeployWalletsAndTokens)

    // ... rest of the main function ...
}