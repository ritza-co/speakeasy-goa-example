# DrinkOperations
(*drinkOperations*)

### Available Operations

* [orderNumberTea](#ordernumbertea) - tea order

## orderNumberTea

Order a cup of tea.

### Example Usage

```typescript
import { SDK } from "openapi";

(async() => {
  const sdk = new SDK();

  const res = await sdk.drinkOperations.orderNumberTea({
    includeMilk: true,
    isGreen: false,
    numberSugars: 1584355970564842800,
  });

  if (res.statusCode == 200) {
    // handle response
  }
})();
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `request`                                                      | [shared.TeaRequestBody](../../models/shared/tearequestbody.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `config`                                                       | [AxiosRequestConfig](https://axios-http.com/docs/req_config)   | :heavy_minus_sign:                                             | Available config options for making requests.                  |


### Response

**Promise<[operations.OrderNumberTeaResponse](../../models/operations/ordernumbertearesponse.md)>**

