# MusicOperations
(*musicOperations*)

### Available Operations

* [bandNumberPlay](#bandnumberplay) - play band

## bandNumberPlay

Choose your jazz style.

### Example Usage

```typescript
import { SDK } from "openapi";
import { PlayRequestBodyStyle } from "openapi/dist/sdk/models/shared";

(async() => {
  const sdk = new SDK();

  const res = await sdk.musicOperations.bandNumberPlay({
    style: PlayRequestBodyStyle.Bebop,
  });

  if (res.statusCode == 200) {
    // handle response
  }
})();
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `request`                                                        | [shared.PlayRequestBody](../../models/shared/playrequestbody.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `config`                                                         | [AxiosRequestConfig](https://axios-http.com/docs/req_config)     | :heavy_minus_sign:                                               | Available config options for making requests.                    |


### Response

**Promise<[operations.BandNumberPlayResponse](../../models/operations/bandnumberplayresponse.md)>**

