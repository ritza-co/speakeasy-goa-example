const SDK = require("./dist/index");

(async() => {
  const sdk = new SDK.SDK();

  const res = await sdk.drinkOperations.orderNumberTea({
    includeMilk: true,
    isGreen: false,
    numberSugars: 1584355970564842800,
  });

  if (res.statusCode == 200) {
    'A nice cup of tea'
  }
})();