import * as operations from "./models/operations";
import { SDKConfiguration } from "./sdk";
import { AxiosRequestConfig } from "axios";
/**
 * A band that plays jazz.
 */
export declare class Band {
    private sdkConfiguration;
    constructor(sdkConfig: SDKConfiguration);
    /**
     * Download ./gen/http/openapi.json
     */
    bandNumberOpenapiJson(config?: AxiosRequestConfig): Promise<operations.BandNumberOpenapiJsonResponse>;
}
