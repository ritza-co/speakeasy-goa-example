<!-- Start SDK Example Usage -->


```typescript
import { SDK } from "openapi";

(async () => {
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
<!-- End SDK Example Usage -->