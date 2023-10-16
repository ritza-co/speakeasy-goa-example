<!-- Start SDK Example Usage -->


```typescript
import { SDK } from "openapi";

(async () => {
    const sdk = new SDK();

    const res = await sdk.band.bandNumberOpenapiJson();

    if (res.statusCode == 200) {
        // handle response
    }
})();

```
<!-- End SDK Example Usage -->