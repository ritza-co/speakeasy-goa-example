import * as operations from "./models/operations";
import * as shared from "./models/shared";
import { SDKConfiguration } from "./sdk";
import { AxiosRequestConfig } from "axios";
export declare class DrinkOperations {
    private sdkConfiguration;
    constructor(sdkConfig: SDKConfiguration);
    /**
     * tea order
     *
     * @remarks
     * Order a cup of tea.
     */
    orderNumberTea(req: shared.TeaRequestBody, config?: AxiosRequestConfig): Promise<operations.OrderNumberTeaResponse>;
}
