# Band
(*band*)

## Overview

A band that plays jazz.

### Available Operations

* [bandNumberOpenapiJson](#bandnumberopenapijson) - Download ./gen/http/openapi.json

## bandNumberOpenapiJson

Download ./gen/http/openapi.json

### Example Usage

```typescript
import { SDK } from "openapi";

(async() => {
  const sdk = new SDK();

  const res = await sdk.band.bandNumberOpenapiJson();

  if (res.statusCode == 200) {
    // handle response
  }
})();
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `config`                                                     | [AxiosRequestConfig](https://axios-http.com/docs/req_config) | :heavy_minus_sign:                                           | Available config options for making requests.                |


### Response

**Promise<[operations.BandNumberOpenapiJsonResponse](../../models/operations/bandnumberopenapijsonresponse.md)>**

